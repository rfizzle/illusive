package illusive

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// DefaultTimeout defines the default timeout for HTTP clients.
var DefaultTimeout = 10 * time.Second

// A Client manages communication with the Illusive API.
type Client struct {
	c        http.Client
	baseUrl  string
	makeAuth func(req *http.Request, params *url.Values) error
	t        *http.Transport
	token    string
}

// Dynamic options for client setup
type OptionClient func(*Client) error

// Dynamic parameter options for requests
type RequestParameter func() (string, string)

// NewClient creates a new Illusive client with Basic Authentication.
//
// baseUrl is the domain with the correct protocol set. ex (https://illusive.example.com)
func NewClient(basePath, token string, options ...OptionClient) (*Client, error) {
	t := &http.Transport{}
	client := &Client{
		c:       http.Client{Transport: t, Timeout: DefaultTimeout},
		baseUrl: basePath,
		t:       t,
		token:   token,
		makeAuth: func(req *http.Request, params *url.Values) error {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
			return nil
		},
	}

	// Option parameter values:
	for _, op := range options {
		err := op(client)
		if err != nil {
			return nil, err
		}
	}

	return client, nil
}

// Build request with correct headers, parameters, and body
func (c *Client) buildRequest(method, restPath string, data *url.Values) (*http.Request, error) {
	var (
		req *http.Request
		err error
	)
	if data != nil {
		methodUpper := strings.ToUpper(method)
		switch methodUpper {
		case "GET":
			req, err = http.NewRequest(methodUpper, c.makeURL(fmt.Sprintf("%s?%s", restPath, data.Encode())), nil)
			if err != nil {
				return nil, err
			}
			req.Header.Add("Content-Type", "application/json")
		default:
			// TODO: POST and other requests
		}
	} else {
		req, err = http.NewRequest(method, c.makeURL(restPath), nil)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to create illusive request: %v", err)
	}
	req.Header.Add("Accept", "application/json")
	if err := c.makeAuth(req, data); err != nil {
		return nil, err
	}
	return req, nil
}

// doRequest handles the request along with error handling and extracting the reponse body
func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	// Conduct request
	res, err := c.c.Do(req)

	// Handle errors
	if err != nil {
		return nil, err
	}

	// Handle non 200s
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("response status: %s", res.Status)
	}

	// Defer body close and read body
	defer res.Body.Close()
	responseData, err := ioutil.ReadAll(res.Body)

	// Handle errors
	if err != nil {
		return nil, fmt.Errorf("error: %v; response code: %d; response: %s", err, res.StatusCode, string(responseData))
	}

	return responseData, err
}

// makeURL creates an URL from the client base URL and a given REST path. What
// this function actually does is to concatenate the base URL and the REST path
// by handling trailing slashes.
func (c *Client) makeURL(restPath string) string {
	return strings.TrimSuffix(c.baseUrl, "/") + "/" + strings.TrimSuffix(strings.TrimPrefix(restPath, "/"), "/")
}
