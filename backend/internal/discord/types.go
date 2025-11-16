package discord

type Handshake struct {
	V        int    `json:"v"`
	ClientID string `json:"client_id"`
}

type Args struct {
	ClientId string   `json:"client_id"`
	Scopes   []string `json:"scopes"`
}

type Authorize struct {
	Nonce     string `json:"nonce"`
	Arguments Args   `json:"args"`
	Command   string `json:"cmd"`
}

type Data struct {
	Code string `json:"code"`
}

type ResponseAuthorize struct {
	Cmd   string `json:"cmd"`
	Data  Data   `json:"data"`
	Evt   any    `json:"evt"`
	Nonce string `json:"nonce"`
}

type Oauth2Response struct {
	TokenType    string `json:"token_type"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

type Authenticate struct {
	Nonce     string `json:"nonce"`
	Arguments struct {
		AccessToken string `json:"access_token"`
	} `json:"args"`
	Command string `json:"cmd"`
}

type SetActivityMessage struct {
	Cmd   string          `json:"cmd"`
	Args  SetActivityArgs `json:"args"`
	Nonce string          `json:"nonce"`
}

type SetActivityArgs struct {
	PID      int       `json:"pid"`
	Activity *Activity `json:"activity,omitempty"`
}

type Activity struct {
	Name       string      `json:"name,omitempty"`
	State      string      `json:"state,omitempty"`
	StateUrl   string      `json:"state_url,omitempty"`
	Details    string      `json:"details,omitempty"`
	Timestamps *Timestamps `json:"timestamps,omitempty"`
	Assets     *Assets     `json:"assets,omitempty"`
	Instance   bool        `json:"instance,omitempty"`
}

type Timestamps struct {
	Start int64 `json:"start,omitempty"`
	End   int64 `json:"end,omitempty"`
}

type Assets struct {
	LargeImage string `json:"large_image,omitempty"`
	LargeText  string `json:"large_text,omitempty"`
	SmallImage string `json:"small_image,omitempty"`
	SmallText  string `json:"small_text,omitempty"`
}
