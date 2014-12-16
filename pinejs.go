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
}

func (a *Client) List(resSlice interface{}) error {
	path := resourceName(resourceFromSlice(resSlice))

	return a.request("GET", path, nil, nil, resSlice)
}

func (a *Client) Create(res interface{}) error {
	// Should POST
	return errors.New("Not implemented")
}

func (a *Client) Update(res interface{}) error {
	// Should PUT
	return errors.New("Not implemented")
}

func (a *Client) Patch(res interface{}) error {
	// Should PATCH
	return errors.New("Not implemented")
}
func (a *Client) Delete(res interface{}) error {
	// Should DELETE
	return errors.New("Not implemented")
}
