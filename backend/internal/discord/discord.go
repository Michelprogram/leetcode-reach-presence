package discord

import (
	"context"
	"leetcode-rich-presence/internal/server/handlers"
	"time"
)

type Discord struct {
	ClientID     string
	ClientSecret string
	Ipc          *IPC
}

func NewDiscord(clientID, clientSecret string) (*Discord, error) {

	ipc, err := NewIPC()

	if err != nil {
		return nil, err
	}

	return &Discord{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Ipc:          ipc,
	}, nil
}

func (d Discord) InitConnectionIpc() error {

	err := d.Ipc.Handshake(d.ClientID)
	if err != nil {
		return err
	}

	if t, err := LoadTokens(); err == nil {
		// If still valid, use it
		if time.Now().Before(t.ExpiresAt.Add(-30 * time.Second)) {
			return d.Ipc.Authenticate(t.AccessToken)
		}
		// Try refresh
		if t.RefreshToken != "" {
			if oauth, err := d.Ipc.RefreshToken(t.RefreshToken, d.ClientID, d.ClientSecret); err == nil && oauth.AccessToken != "" {
				newT := Tokens{
					AccessToken:  oauth.AccessToken,
					RefreshToken: oauth.RefreshToken,
					ExpiresAt:    time.Now().Add(time.Duration(oauth.ExpiresIn) * time.Second),
				}
				_ = SaveTokens(newT)
				return d.Ipc.Authenticate(oauth.AccessToken)
			}
		}
	}

	auth, err := d.Ipc.Authorization(d.ClientID)

	if err != nil {
		return err
	}

	oauth, err := d.Ipc.OAuth2Exchange(auth.Data.Code, "http://localhost", d.ClientID, d.ClientSecret)

	if err != nil {
		return err
	}

	_ = SaveTokens(Tokens{
		AccessToken:  oauth.AccessToken,
		RefreshToken: oauth.RefreshToken,
		ExpiresAt:    time.Now().Add(time.Duration(oauth.ExpiresIn) * time.Second),
	})

	return d.Ipc.Authenticate(oauth.AccessToken)
}

func (d Discord) ListenWithContext(ctx context.Context, queue <-chan handlers.Message) error {

	err := d.InitConnectionIpc()
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case v, ok := <-queue:
			if !ok {
				return nil
			}

			err := d.Ipc.SetActivity(v.Title, v.Url)

			if err != nil {
				return err
			}
		}
	}
}
