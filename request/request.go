package request

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/spf13/afero"
	"gopkg.in/yaml.v3"
)

var fs = afero.NewOsFs()

type Request struct {
	Method  string            `yaml:"method"`
	Host    string            `yaml:"host"`
	Path    string            `yaml:"path"`
	Headers map[string]string `yaml:"headers"`
	Body    []byte            `yaml:"body"`
}

func Open(name string) (*Request, error) {
	f, err := fs.Open(name)
	if err != nil {
		return nil, err
	}

	r := new(Request)
	err = yaml.NewDecoder(f).Decode(&r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (r Request) Execute(client *http.Client) (*http.Response, error) {
	url, err := url.Parse(r.Host + r.Path)
	if err != nil {
		return nil, fmt.Errorf("failed to parse host, %w", err)
	}

	req, err := http.NewRequest(r.Method, url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create new request: %w", err)
	}

	for k, v := range r.Headers {
		req.Header.Add(k, v)
	}

	return client.Do(req)
}
