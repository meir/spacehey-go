package spacehey_test

import (
	"os"
	"testing"

	"github.com/meir/spacehey-go"
)

var email = os.Getenv("SPACEHEY_EMAIL")
var password = os.Getenv("SPACEHEY_PASSWORD")

func spaceheyClient() *spacehey.Client {
	return spacehey.NewClient(email, password)
}

func Test_Authenticate(t *testing.T) {
	client := spaceheyClient()

	err := client.Authenticate()
	if err != nil {
		t.Fatal(err)
	}
}
