package pinejs

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

	if !strings.HasPrefix(path, "/") {
		path = "/" + path
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
