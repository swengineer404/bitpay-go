package bitpay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type apiClient struct {
	client *http.Client
	url    string
	token  string
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"error"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func newAPIClient(url string, token string) *apiClient {
	return &apiClient{
		client: &http.Client{
			Timeout: time.Second * 30,
		},
		url:   url,
		token: token,
	}
}

func (c *apiClient) send(path, method string, params, out any) error {
	b, _ := json.Marshal(params)
	m := map[string]any{}
	json.Unmarshal(b, &m)
	m["token"] = c.token
	b, _ = json.Marshal(m)
	
	req, err := http.NewRequest(method, c.url+path, bytes.NewReader(b))
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", "bitpay-go")
	req.Header.Set("Content-Type", "application/json")

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		var resErr Error
		if err := json.NewDecoder(res.Body).Decode(&resErr); err != nil {
			return err
		}

		return &resErr
	}

	if err := json.NewDecoder(res.Body).Decode(out); err != nil {
		return err
	}

	return nil
}
