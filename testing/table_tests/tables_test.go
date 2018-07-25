package table_tests

import (
	"net/http"
	"testing"
	"time"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestDownload(t *testing.T) {
	var urls = []struct {
		url        string
		statusCode int
	}{
		{
			"https://www.ardanlabs.com/blog/index.xml",
			http.StatusOK,
		},
		{
			"http://rss.cnn.com/rss/cnn_topstbadurl.rss",
			http.StatusNotFound,
		},
	}

	client := http.Client{
		Timeout: time.Second * 2,
	}

	for _, url := range urls {
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"", url.url, url.statusCode)
		{
			resp, err := client.Get(url.url)
			if err != nil {
				t.Fatalf("\t\tShould be able to GET URL. %v\n\t\t\t%v", ballotX, err)
			}
			t.Log("\t\tShould be able to GET the URL", checkMark)
			defer resp.Body.Close()

			if resp.StatusCode == url.statusCode {
				t.Logf("\t\tShould have a \"%d\" status. %v", url.statusCode, checkMark)
			} else {
				t.Errorf("\t\tShould have a \"%d\" status. %v\n\t\t\t%v", url.statusCode, ballotX, resp.StatusCode)
			}
		}
	}
}
