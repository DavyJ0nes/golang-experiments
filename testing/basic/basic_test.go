package basic

import (
	"net/http"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

// TestDownload validates the http Get function can download content
func TestDownload(t *testing.T) {
	url := "https://www.ardanlabs.com/blog/index.xml"
	statusCode := 200

	t.Log("Given the need to test downloading content.")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"", url, statusCode)
		resp, err := http.Get(url)
		if err != nil {
			t.Fatal("\t\tShould be able to make Get request.", ballotX, err)
		}
		t.Log("\t\tShould be able to make Get request.", checkMark)
		defer resp.Body.Close()

		if resp.StatusCode == statusCode {
			t.Logf("\t\tShould receive a \"%d\" status. %v", statusCode, checkMark)
		} else {
			t.Errorf("\t\tShould receive a \"%d\" status. Got: \"%v\" %v", statusCode, resp.StatusCode, ballotX)
		}
	}
}
