// backend/internal/discord/token_store.go
package discord

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type Tokens struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}

func tokensPath() (string, error) {
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(cfgDir, "leetcode-rich-presence")
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return "", err
	}
	return filepath.Join(dir, "discord_tokens.json"), nil
}

func LoadTokens() (Tokens, error) {
	p, err := tokensPath()
	if err != nil {
		return Tokens{}, err
	}
	b, err := os.ReadFile(p)
	if err != nil {
		return Tokens{}, err
	}
	var t Tokens
	if err := json.Unmarshal(b, &t); err != nil {
		return Tokens{}, err
	}
	return t, nil
}

func SaveTokens(t Tokens) error {
	p, err := tokensPath()
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(t, "", "\t")
	if err != nil {
		return err
	}
	return os.WriteFile(p, b, 0o600)
}