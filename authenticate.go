package spacehey

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const spacehey_auth_uri = "https://auth.spacehey.com/"
const spacehey_cookie = "SPACEHEY_SESSID"

func (c *Client) Authenticate() error {
	payload := map[string][]string{
		"client_id": {"web"},
		"return":    {""},
		"email":     {c.credentials.email},
		"password":  {c.credentials.password},
		"action":    {"login"},
	}

	request, err := http.NewRequest(
		http.MethodPost,
		spacehey_auth_uri,
		strings.NewReader(url.Values(payload).Encode()),
	)

	response, err := c.client.Do(request)
	if err != nil {
		return err
	}

	for _, cookie := range response.Cookies() {
		if cookie.Name == spacehey_cookie {
			return nil
		}
	}

	if response.Request.URL.Path == "/home" {
		return nil
	}

	return fmt.Errorf("authentication failed")
}
