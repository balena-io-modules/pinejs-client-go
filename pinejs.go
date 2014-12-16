// pinejs implements an interface to pine.js
// See https://bitbucket.org/rulemotion/pinejs
package pinejs

import (
	"errors"
	"log"
	"os"
)

// The current implementation version.
func Version() string {
	return "0.0.1"
}

const (
	logFlags      = log.Ldate | log.Ltime
	logFlagsDebug = logFlags | log.Lshortfile
)

var (
	logAlert = log.New(os.Stderr, "", logFlags)
	logDebug *log.Logger // Defaults to /dev/null in init()
	nullFile *os.File
)

func init() {
	var err error
	if nullFile, err = os.Open(os.DevNull); err != nil {
		ShowDebugOutput()
		logDebug.Printf("Can't open null output at %s, defaulting to debug logs on.",
			os.DevNull)
	} else {
		// Default to no debug output.
		logDebug = log.New(nullFile, "", 0)
	}
}

// ShowDebugOutput enables the debug logger, outputting to stderr.
func ShowDebugOutput() {
	if nullFile != nil {
		nullFile.Close()
		nullFile = nil
	}

	logDebug = log.New(os.Stderr, "", logFlagsDebug)
}

// Client represents an HTTP client to a pine.js server.
type Client struct {
	APIKey, Endpoint string
	reader           io.Reader
}

// NewClient returns a client initialised to the provided endpoint and API key.
func NewClient(endpoint, apiKey string) *Client {
	return &Client{apiKey, endpoint, nil}
}

	return a.request("GET", path, nil, nil, &[]interface{}{res})
func (c *Client) Get(v interface{}, filters ...ODataFilter) error {
	if err := assertPointerToStruct(v); err != nil {
		return err
	}

	if resourceId(v) == 0 {
		return errors.New("id not set")
	}

	if data, err := c.request("GET", getSinglePath(v), nil, Filters(filters)); err != nil {
		return err
	} else {
		return decode(v, data, first)
	}
}

func (c *Client) List(v interface{}, filters ...ODataFilter) error {
	if err := assertPointerToSliceStructs(v); err != nil {
		return err
	}

	if data, err := c.request("GET", resourceName(v), nil, Filters(filters)); err != nil {
		return err
	} else {
		return decode(v, data, theD)
	}
}

func (c *Client) Create(v interface{}, filters ...ODataFilter) error {
	if err := assertPointerToStruct(v); err != nil {
		return err
	}

	if resourceId(v) > 0 {
		return errors.New("attempting to create with id set")
	}

	if data, err := c.request("POST", resourceName(v), v, Filters(filters)); err != nil {
		return err
	} else {
		return decode(v, data, self)
	}
}

func (c *Client) Update(v interface{}) error {
	if err := assertPointerToStruct(v); err != nil {
		return err
	}

	if resourceId(v) == 0 {
		return errors.New("id not set")
	}

	if _, err := c.request("PUT", getSinglePath(v), v, nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) Patch(v interface{}) error {
	if err := assertPointerToStruct(v); err != nil {
		return err
	}

	if resourceId(v) == 0 {
		return errors.New("id not set")
	}

	if _, err := c.request("PATCH", getSinglePath(v), v, nil); err != nil {
		return err
	}

	return nil
}
func (a *Client) Delete(res interface{}) error {
	// Should DELETE
	return errors.New("Not implemented")

func (c *Client) Delete(v interface{}) error {
	if err := assertPointerToStruct(v); err != nil {
		return err
	}

	if resourceId(v) == 0 {
		return errors.New("id not set")
	}

	if _, err := c.request("DELETE", getSinglePath(v), v, nil); err != nil {
		return err
	}

	return nil
}
