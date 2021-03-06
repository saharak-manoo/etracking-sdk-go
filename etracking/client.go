package etracking

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

// APIEndpoint constants
const (
	APIEndpointBase             = "https://fast.etrackings.com"

	APIEndpointTracksFind       = "/api/v3/tracks/find"
	APIEndpointTracks           = "/api/v3/tracks/%s"
)

// Client type
type Client struct {
	etrackingsApiKey           string
	etrackingsKeySecret        string
	endpointBase               *url.URL
	httpClient                 *http.Client
}

// ClientOption type
type ClientOption func(*Client) error

// New returns a new bot client instance.
func New(etrackingsApiKey, etrackingsKeySecret string) (*Client, error) {
	if etrackingsApiKey == "" {
		return nil, errors.New("Missing etrackings api key")
	}

	if etrackingsKeySecret == "" {
		return nil, errors.New("Missing etrackings key secret")
	}

	c := &Client{
		etrackingsApiKey: etrackingsApiKey,
		etrackingsKeySecret: etrackingsKeySecret,
		httpClient: http.DefaultClient,
	}

	if c.endpointBase == nil {
		u, err := url.ParseRequestURI(APIEndpointBase)
		if err != nil {
			return nil, err
		}
		c.endpointBase = u
	}

	return c, nil
}

func (client *Client) url(base *url.URL, endpoint string) string {
	u := *base
	u.Path = path.Join(u.Path, endpoint)
	return u.String()
}

func (client *Client) do(req *http.Request) (*http.Response, error) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Etrackings-api-key", client.etrackingsApiKey)
	req.Header.Set("Etrackings-key-secret", client.etrackingsKeySecret)

	return client.httpClient.Do(req)
}

func (client *Client) post(endpoint string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, client.url(client.endpointBase, endpoint), body)

	if err != nil {
		return nil, err
	}

	return client.do(req)
}

func closeResponse(res *http.Response) error {
	defer res.Body.Close()
	_, err := io.Copy(ioutil.Discard, res.Body)
	return err
}