package spacehey

import (
	"fmt"
	"net/http"
)

const spacehey_auth_uri = "https://auth.spacehey.com/"

func (c *Client) Authenticate() error {
	payload := payloadMap{
		"client_id": "web",
		"return":    "",
		"email":     c.credentials.email,
		"password":  c.credentials.password,
		"action":    "login",
	}

	request, err := http.NewRequest(
		http.MethodPost,
		spacehey_auth_uri,
		buildForm(payload, false),
	)

	response, err := c.client.Do(request)
	if err != nil {
		return err
	}

	// Check if the email input exists on the page, if it does, we know that the login failed.
	if getHtmlField(response.Body, htmlField{
		tag: "input",
		attributes: map[string]string{
			"name": "email",
		},
	}) != "" {
		return fmt.Errorf("failed to log in")
	}

	return nil
}
