package pinejs

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func (c *Client) sanitisePath(path string) (string, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	// Strip host information so we request /foo/bar rather than
	// http://example.com/foo/bar. This affects the GET header and failing to do
	// this can cause the server to erroneously 404.
	//
	// We generate a URL combining the endpoint and the provided path so we pick
	// up all of the path to request, e.g. endpoint
	// 'http://example.com/foo/bar/baz', path '/gwargh/fwargh/mwargh' should
	// produce a sanitised path of '/foo/bar/baz/gwargh/fwargh/mwargh'.
	if u, err := url.Parse(c.Endpoint + path); err != nil {
		return "", err
	} else {
		return u.Path, nil
	}
}

	query.Add("apikey", a.APIKey)
	//query.Add("$expand", "device")

	// path += "?" + query.Encode()

	req, _ := http.NewRequest(method, a.Endpoint, nil)
	req.URL.Opaque = path + "?apikey=" + a.APIKey + "&$expand=device"

	req.Header.Add("User-Agent", "PineJS/v1 GoBindings/"+VERSION)

	log.Printf("Requesting %v %q\n", method, path)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Printf("Request failed: %v\n", err)
		return err
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
}
