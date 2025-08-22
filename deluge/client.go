package deluge

import (
	"bytes"
	"delugerpc/jsonrpc"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

type Client struct {
	httpClient *http.Client
	targetUrl  string
	password   string
}

func NewClient(delugeHost string, delugeWebuiPort int, delugePassword string) (*Client, error) {
	scheme := "http://"
	if strings.HasPrefix(delugeHost, "http") {
		scheme = ""
	}

	targetUrl, err := url.Parse(fmt.Sprintf("%s%s:%d/json", scheme, delugeHost, delugeWebuiPort))

	if err != nil {
		return nil, err
	}

	cookieJar, _ := cookiejar.New(nil)
	httpClient := &http.Client{
		Jar: cookieJar,
	}

	return &Client{
		httpClient: httpClient,
		targetUrl:  targetUrl.String(),
		password:   delugePassword,
	}, nil
}

// SendRequest sends a JSON-RPC v1 request to the Deluge WebUI. If the request fails due to an authentication error,
// the client will attempt to re-authenticate and resend the request.
func SendRequest[TResp any](c *Client, method string, params ...any) (*jsonrpc.V1Response[TResp], error) {
	response, err := sendRequest[TResp](c, method, params...)
	if err != nil && err.(*jsonrpc.V1Error).IsNotAuthenticated() {
		// Auth and try again
		_, err = c.Login()
		if err != nil {
			return nil, err
		}
		response, err = sendRequest[TResp](c, method, params...)
	}

	// Either failed twice or Auth failed
	return response, err
}

func sendRequest[TResp any](c *Client, method string, params ...any) (*jsonrpc.V1Response[TResp], error) {
	// Prepare rpc request body
	requestBody := jsonrpc.NewV1Request(method, params...)
	requestBytes, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(requestBytes)

	// Send request
	response, err := c.httpClient.Post(c.targetUrl, "application/json", buffer)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	// Extract response
	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	responseBody := &jsonrpc.V1Response[TResp]{}
	err = json.Unmarshal(responseBytes, responseBody)
	if err != nil {
		return nil, err
	}

	if responseBody.Error != nil {
		return responseBody, responseBody.Error
	}

	return responseBody, nil
}
