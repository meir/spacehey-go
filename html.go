package spacehey

import (
	"io"

	"golang.org/x/net/html"
)

type htmlField struct {
	tag        string
	attributes map[string]string
}

func getHtmlField(src io.Reader, field htmlField) string {
	dom := html.NewTokenizer(src)

	for {
		tokenType := dom.Next()
		if tokenType == html.ErrorToken {
			return ""
		}

		token := dom.Token()
		if tokenType == html.StartTagToken {
			if token.Data == field.tag {
				if field.attributes != nil {
					for _, attr := range token.Attr {
						if field.attributes[attr.Key] == attr.Val {
							return dom.Next().String()
						}
					}
				}
			}
		}
	}
}
