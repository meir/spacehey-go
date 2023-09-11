package spacehey

import (
	"fmt"
	"io"
	"net/url"
	"strings"
)

type payloadMap map[string]map[string]string

func buildForm(payload map[string]map[string]string, submit bool) io.Reader {
	valuesMap := map[string][]string{}

	for key, value := range payload {
		for subkey, v := range value {
			payloadKey := fmt.Sprintf("%s[%s]", key, subkey)
			valuesMap[payloadKey] = []string{v}
		}
	}

	if submit {
		valuesMap["submit"] = []string{""}
	}

	return strings.NewReader(url.Values(valuesMap).Encode())
}
