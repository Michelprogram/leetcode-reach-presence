package discord

import (
	"context"
	"leetcode-rich-presence/internal/server/handlers"
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

	auth, err := d.Ipc.Authorization(d.ClientID)

	if err != nil {
		return err
	}

	oauth, err := d.Ipc.OAuth2Exchange(auth.Data.Code, "http://localhost", d.ClientID, d.ClientSecret)

	if err != nil {
		return err
	}

	err = d.Ipc.Authenticate(oauth.AccessToken)

	if err != nil {
		return err
	}

	return nil
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
