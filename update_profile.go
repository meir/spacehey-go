package spacehey

import (
	"fmt"
	"net/http"
)

type ProfileSections struct {
	AboutMe    string
	Meet       string
	General    string
	Music      string
	Movies     string
	Television string
	Books      string
	Heroes     string
}

const spacehey_edit_uri = "https://spacehey.com/edit"

func (c *Client) UpdateProfile(sections ProfileSections) error {
	payload := payloadMap{
		"category": map[string]string{
			"about_me":   sections.AboutMe,
			"meet":       sections.Meet,
			"general":    sections.General,
			"music":      sections.Music,
			"movies":     sections.Movies,
			"television": sections.Television,
			"books":      sections.Books,
			"heroes":     sections.Heroes,
		},
	}

	request, err := http.NewRequest(
		http.MethodPost,
		spacehey_edit_uri,
		buildForm(payload, true),
	)
	if err != nil {
		return err
	}

	response, err := c.client.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("status code %d", response.StatusCode)
	}

	return nil
}
