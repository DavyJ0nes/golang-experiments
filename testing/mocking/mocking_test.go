package mocking

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

// mockServer returns pointer to server to handle request call
func mockServer() *httptest.Server {
	var feed = `<?xml version="1.0" encoding="UTF-8"?>
<rss>
<channel>
	<title>Going Go Programming</title>
	<description>Golang : https://github.com/goinggo</description>
	<link>http://www.goinggo.net/</link>
	<item>
		<pubDate></pubDate>
		<title></title>
		<description></description>
		<link></link>
	</item>
</channel>
</rss>`

	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintf(w, feed)
	}

	return httptest.NewServer(http.HandlerFunc(f))
}

// TestDownload validates the http Get function can download content
func TestDownload(t *testing.T) {
	srv := mockServer()
	defer srv.Close()

	t.Log("Given the need to test downloading content.")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"", srv.URL, http.StatusOK)
		resp, err := http.Get(srv.URL)
		if err != nil {
			t.Fatal("\t\tShould be able to make Get request.", ballotX, err)
		}
		t.Log("\t\tShould be able to make Get request.", checkMark)
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			t.Logf("\t\tShould receive a \"%d\" status. %v", http.StatusOK, checkMark)
		} else {
			t.Errorf("\t\tShould receive a \"%d\" status. Got: \"%v\" %v", http.StatusOK, resp.StatusCode, ballotX)
		}
	}
}
