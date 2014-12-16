package pinejs

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type mockServerCommand int

const mockPort = 1234

var (
	pathRegexp = regexp.MustCompile(`(\w+)/(\w+)(?:\(\d+\))?`)
	server     *http.Server
	sourceData string
)

func sanitisePath(path string) string {
	// Make the path somewhat canonical.
	ret := strings.ToLower(path)
	ret = strings.TrimSpace(ret)
	return strings.TrimLeft(ret, "/")
}

func mockServer(data string) {
	sourceData = data

	if server != nil {
		return
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		path := sanitisePath(req.URL.Path)
		matches := pathRegexp.FindStringSubmatch(path)

		if len(matches) == 0 {
			http.NotFound(w, req)
			return
		}

		// if len(matches) != 0, len(matches) == 3.
		model := matches[1]
		//resource := matches[2]
		//id := matches[3]

		if model != "ewa" {
			http.NotFound(w, req)
			return
		}

		fmt.Fprintf(w, sourceData)
	})

	go func() {
		url := fmt.Sprintf(":%d", mockPort)

		server = &http.Server{Addr: url, Handler: mux}

		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
}

func mockServerFromPath(path string) error {
	if file, err := os.Open(path); err != nil {
		return err
	} else {
		defer file.Close()
		if bs, err := ioutil.ReadAll(file); err != nil {
			return err
		} else {
			mockServer(string(bs))
		}

		return nil
	}
}

func mockServerClientFromPath(path string) (*Client, error) {
	if err := mockServerFromPath(path); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("http://localhost:%d/ewa", mockPort)
	return NewClient(url, ""), nil
}
