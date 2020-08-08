package misskey

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"path"
)

type Misskey struct {
	Host  *url.URL
	Token string
}

func NewMisskey(host, token string) (*Misskey, error) {
	url, err := url.Parse(host)
	if nil != err {
		return nil, err
	}

	return &Misskey{Host: url, Token: token}, nil
}
func (m *Misskey) NotePost(message string) error {
	payload := map[string]interface{}{
		"i":          m.Token,
		"cw": "@gizenchan@best-friends.chat",
		"visibility": "home",
		"text":       message,
	}
	payloadBytes, err := json.Marshal(payload)
	if nil != err {
		return err
	}

	url := m.Host
	url.Path = path.Join(url.Path, "/api/notes/create")

	req, err := http.NewRequest(
		"POST",
		url.String(),
		bytes.NewBuffer(payloadBytes),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return err
	}

	return nil
}
