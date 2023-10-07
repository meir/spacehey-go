package spacehey

import (
	"fmt"
	"io"
	"net/url"
	"strings"
)

type mapOrString interface{}

type payloadMap map[string]mapOrString

func buildForm(payload payloadMap, submit bool) io.Reader {
	valuesMap := map[string][]string{}

	for key, value := range payload {
		switch v := value.(type) {
		case string:
			valuesMap[key] = []string{v}
		case map[string]string:
			for subkey, subvalue := range v {
				payloadKey := fmt.Sprintf("%s[%s]", key, subkey)
				valuesMap[payloadKey] = []string{subvalue}
			}
		}
	}

	if submit {
		valuesMap["submit"] = []string{""}
	}

	return strings.NewReader(url.Values(valuesMap).Encode())
}
