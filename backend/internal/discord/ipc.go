package discord

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type Opcode = uint

const (
	OPCODE_ONE  Opcode = 1
	OPCODE_ZERO Opcode = 0
)

type IPC struct {
	Connector net.Conn
}

func NewIPC() (*IPC, error) {

	tmpDirectory := os.Getenv("TMPDIR")
	if tmpDirectory == "" {
		tmpDirectory = "/tmp/"
	}

	sockPath := fmt.Sprintf("%sdiscord-ipc-0", tmpDirectory)
	conn, err := net.Dial("unix", sockPath)
	if err != nil {
		return nil, err
	}

	return &IPC{
		Connector: conn,
	}, nil
}

func (ipc IPC) sendMessage(data any, opcode Opcode) ([]byte, error) {

	var buf bytes.Buffer

	payload, err := json.Marshal(data)

	if err != nil {
		return []byte{}, err
	}

	binary.Write(&buf, binary.LittleEndian, int32(opcode))
	binary.Write(&buf, binary.LittleEndian, int32(len(payload)))
	buf.Write(payload)

	ipc.Connector.Write(buf.Bytes())

	resp := make([]byte, 2048)
	n, err := ipc.Connector.Read(resp)

	if err != nil {
		return []byte{}, err
	}

	return resp[8:n], nil

}

func (ipc IPC) Handshake(clientID string) error {
	res, err := ipc.sendMessage(Handshake{V: 1, ClientID: clientID}, OPCODE_ZERO)

	if err != nil {
		return err
	}

	log.Printf("Handshake Response: %s\n", res)

	return nil

}

func (ipc IPC) Authorization(clientID string) (ResponseAuthorize, error) {
	auth := Authorize{
		Nonce:   RandStringRunes(15),
		Command: "AUTHORIZE",
		Arguments: Args{
			ClientId: clientID,
			Scopes:   []string{"rpc", "identify", "rpc.activities.write"},
		},
	}

	res, err := ipc.sendMessage(auth, OPCODE_ONE)

	if err != nil {
		return ResponseAuthorize{}, err
	}

	var response ResponseAuthorize
	err = json.Unmarshal(res, &response)
	if err != nil {
		return ResponseAuthorize{}, err
	}

	return response, err
}

func (ipc IPC) OAuth2Exchange(code, redirectUri, clientID, secretID string) (Oauth2Response, error) {
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", redirectUri)

	req, err := http.NewRequest("POST", "https://discord.com/api/v10/oauth2/token", strings.NewReader(data.Encode()))
	if err != nil {
		return Oauth2Response{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(clientID, secretID)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return Oauth2Response{}, err
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return Oauth2Response{}, err
	}

	var oauth2Response Oauth2Response
	err = json.Unmarshal(bodyBytes, &oauth2Response)
	if err != nil {
		return Oauth2Response{}, err
	}

	return oauth2Response, err

}

func (ipc IPC) Authenticate(accessToken string) error {
	authenticate := Authenticate{
		Nonce:   RandStringRunes(15),
		Command: "AUTHENTICATE",
		Arguments: struct {
			AccessToken string `json:"access_token"`
		}{
			AccessToken: accessToken,
		},
	}

	_, err := ipc.sendMessage(authenticate, OPCODE_ONE)

	if err != nil {
		return err
	}

	return nil

}

func (ipc IPC) RefreshToken(refreshToken, clientID, secretID string) (Oauth2Response, error) {
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refreshToken)

	req, err := http.NewRequest("POST", "https://discord.com/api/v10/oauth2/token", strings.NewReader(data.Encode()))
	if err != nil {
		return Oauth2Response{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(clientID, secretID)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return Oauth2Response{}, err
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return Oauth2Response{}, err
	}

	var oauth2Response Oauth2Response
	if err := json.Unmarshal(bodyBytes, &oauth2Response); err != nil {
		return Oauth2Response{}, err
	}
	return oauth2Response, nil
}

func (ipc IPC) SetActivity(title, url string) error {
	now := time.Now()
	activity := SetActivityMessage{
		Cmd: "SET_ACTIVITY",
		Args: SetActivityArgs{
			PID: os.Getpid(),
			Activity: &Activity{
				State:    title,
				StateUrl: url,
				Details:  "Problem Solving",
				Timestamps: &Timestamps{
					Start: now.UnixNano() / int64(time.Millisecond),
				},
				Instance: true,
				Name:     "LeetCode",
			},
		},
		Nonce: RandStringRunes(15),
	}

	res, err := ipc.sendMessage(activity, OPCODE_ONE)

	if err != nil {
		return err
	}

	log.Printf("Activity Response: %s\n", res)

	return nil
}
