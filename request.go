package pinejs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/fatih/structs"
)

type envelope struct {
	Data []interface{} `json:"d"`
}

func (a *Client) request(method, path string, query *url.Values, body *url.Values, v interface{}) (err error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = a.Endpoint + path

	if query == nil {
		query = &url.Values{}
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

	var e envelope
	json.Unmarshal(resBody, &e)

	fmt.Println(e.Data)
	normalise(e.Data, structs.New(resourceFromSlice(v)))

	// json.Unmarshal(json.Marshal(data), v)
	return
}
