package spacehey

import (
	"strings"

	"golang.org/x/net/html"
)

type htmlField struct {
	tag        string
	attributes map[string]string
}

func getHtmlField(src string, field htmlField) (string, error) {
	dom := html.NewTokenizer(strings.NewReader(src))

	for {
		tokenType := dom.Next()
		if tokenType == html.ErrorToken {
			return "", dom.Err()
		}

		token := dom.Token()
		if tokenType == html.StartTagToken {
			if token.Data == field.tag {
				if field.attributes != nil {
					for _, attr := range token.Attr {
						if field.attributes[attr.Key] == attr.Val {
							return dom.Next().String(), nil
						}
					}
				}
			}
		}
	}
}
