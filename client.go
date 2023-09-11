package spacehey

import (
	"net/http"
	"net/http/cookiejar"
)

type credentials struct {
	email    string
	password string
}

type Client struct {
	credentials credentials

	client *http.Client
}

func NewClient(email, password string) *Client {
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}

	return &Client{
		credentials: credentials{
			email:    email,
			password: password,
		},

		client: &http.Client{
			Jar: jar,
		},
	}
}
