package pinejs

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// clientversion is the binding version
const clientversion = "0.0.1"

// Client is the PineJS client.
type Client struct {
	url        string
	apiKey     string
	httpClient *http.Client
}

// Init initializes the PineJS client with the appropriate API key
// as well as the endpoint of the API
func (a *Client) Init(url, key string, httpClient *http.Client) {
	a.url = url
	a.key = key
	if httpClient {
		a.httpClient = httpClient
	} else {
		a.httpClient = http.DefaultClient
	}
}

var debug bool

// SetDebug enables additional tracing globally.
// The method is designed for used during testing.
func SetDebug(value bool) {
	debug = value
}

func (a *Client) Get(resource string, id string) {
	path := resource + "(" + id + ")"
	err = a.request("GET", path, nil, nil)
}

func (a *Client) New(resource string, params interface{})
func (a *Client) Put(resource string, id string, params interface{})
func (a *Client) Patch(resource string, id string, params interface{})
func (a *Client) List(resource string, filters interface{})
func (a *Client) Del(resource string, id string)

// Call is the Backend.Call implementation for invoking Stripe APIs.
func (a *Client) request(method, path, query *url.Values, body *url.Values, v interface{}) error {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = s.url + path

	if body != nil && len(*body) > 0 {
		path += "?" + body.Encode()
	}

	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		log.Printf("Cannot create Stripe request: %v\n", err)
		return err
	}

	req.Header.Add("User-Agent", "PineJS/v1 GoBindings/"+clientversion)

	log.Printf("Requesting %v %q\n", method, path)
	start := time.Now()

	res, err := s.httpClient.Do(req)

	if debug {
		log.Printf("Completed in %v\n", time.Since(start))
	}

	if err != nil {
		log.Printf("Request to Stripe failed: %v\n", err)
		return err
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Cannot parse Stripe response: %v\n", err)
		return err
	}

	if res.StatusCode >= 400 {
		// for some odd reason, the Erro structure doesn't unmarshal
		// initially I thought it was because it's a struct inside of a struct
		// but even after trying that, it still didn't work
		// so unmarshalling to a map for now and parsing the results manually
		// but should investigate later
		var errMap map[string]interface{}
		json.Unmarshal(resBody, &errMap)

		if e, found := errMap["error"]; !found {
			err := errors.New(string(resBody))
			log.Printf("Unparsable error returned from Stripe: %v\n", err)
			return err
		} else {
			root := e.(map[string]interface{})
			err := &Error{
				Type:           ErrorType(root["type"].(string)),
				Msg:            root["message"].(string),
				HTTPStatusCode: res.StatusCode,
			}

			if code, found := root["code"]; found {
				err.Code = ErrorCode(code.(string))
			}

			if param, found := root["param"]; found {
				err.Param = param.(string)
			}

			log.Printf("Error encountered from Stripe: %v\n", err)
			return err
		}
	}

	if debug {
		log.Printf("Stripe Response: %q\n", resBody)
	}

	if v != nil {
		return json.Unmarshal(resBody, v)
	}

	return nil
}
