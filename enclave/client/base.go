package client

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/yangnei/enclave-go/enclave/api"
)

type BaseClient interface {
	Get(path string, params map[string]string, headers map[string]string) (*http.Response, error)
	Post(path string, body string, params map[string]string, headers map[string]string) (*http.Response, error)
	Delete(path string, body string, params map[string]string, headers map[string]string) (*http.Response, error)
	Put(path string, body string, params map[string]string, headers map[string]string) (*http.Response, error)
	HandleError(res *http.Response) error
}

// baseClient handles basic requests and authentication that are common across spotClient, perpsClient, and Cross.
type baseClient struct {
	BaseURL   string
	KeyID     string
	APISecret string
	Client    *http.Client
	UserAgent string
}

// NewBaseClient initializes a new baseClient with the provided API credentials and base URL.
func NewBaseClient(apiKey, apiSecret, baseURL string) BaseClient {
	return &baseClient{
		BaseURL:   strings.TrimRight(baseURL, "/"),
		KeyID:     apiKey,
		APISecret: apiSecret,
		Client:    &http.Client{Timeout: 10 * time.Second},
		UserAgent: "enclave-go",
	}
}

// doRequest sends an HTTP request with authentication and returns the HTTP response.
func (c *baseClient) doRequest(method, path, body string, params map[string]string, headers map[string]string) (*http.Response, error) {
	method = strings.ToUpper(method)
	if method != http.MethodGet && method != http.MethodPost && method != http.MethodDelete && method != http.MethodPut {
		return nil, fmt.Errorf("unsupported HTTP method %s", method)
	}

	// Build URL with query parameters
	u, err := url.Parse(c.BaseURL + path)
	if err != nil {
		return nil, err
	}

	if len(params) > 0 {
		q := u.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		u.RawQuery = q.Encode()
	}

	// Prepare the request body
	var bodyReader io.Reader
	if body != "" {
		bodyReader = strings.NewReader(body)
	}

	// Create the HTTP request
	req, err := http.NewRequest(method, u.String(), bodyReader)
	if err != nil {
		return nil, err
	}

	// Set default headers
	req.Header.Set("User-Agent", c.UserAgent)

	if headers == nil {
		headers = make(map[string]string)
	}

	// Set Content-Type header for POST and PUT requests with a body
	if (method == http.MethodPost || method == http.MethodPut) && body != "" {
		if _, exists := headers["Content-Type"]; !exists {
			headers["Content-Type"] = "application/json"
		}
	}

	// Add custom headers
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// Add authentication headers
	err = c.addAuthHeaders(req, body)
	if err != nil {
		return nil, err
	}

	// Send the HTTP request
	return c.Client.Do(req)
}

// addAuthHeaders calculates the authentication signature and adds the required headers to the request.
func (c *baseClient) addAuthHeaders(req *http.Request, body string) error {
	timestamp := fmt.Sprintf("%d", time.Now().UnixMilli())

	// Extract the clean path and query
	u, err := url.Parse(req.URL.String())
	if err != nil {
		return err
	}
	cleanPath := u.Path
	if u.RawQuery != "" {
		cleanPath += "?" + u.RawQuery
	}

	// Construct the message to sign
	message := timestamp + req.Method + cleanPath + body

	// Compute the HMAC-SHA256 signature
	mac := hmac.New(sha256.New, []byte(c.APISecret))
	mac.Write([]byte(message))
	signature := hex.EncodeToString(mac.Sum(nil))

	// Set authentication headers
	req.Header.Set("ENCLAVE-KEY-ID", c.KeyID)
	req.Header.Set("ENCLAVE-TIMESTAMP", timestamp)
	req.Header.Set("ENCLAVE-SIGN", signature)

	return nil
}

// Get sends a GET request to the specified path with optional parameters and headers.
func (c *baseClient) Get(path string, params map[string]string, headers map[string]string) (*http.Response, error) {
	return c.doRequest(http.MethodGet, path, "", params, headers)
}

// Post sends a POST request to the specified path with the provided body, parameters, and headers.
func (c *baseClient) Post(path string, body string, params map[string]string, headers map[string]string) (*http.Response, error) {
	return c.doRequest(http.MethodPost, path, body, params, headers)
}

// Delete sends a DELETE request to the specified path with the provided body, parameters, and headers.
func (c *baseClient) Delete(path string, body string, params map[string]string, headers map[string]string) (*http.Response, error) {
	return c.doRequest(http.MethodDelete, path, body, params, headers)
}

// Put sends a PUT request to the specified path with the provided body, parameters, and headers.
func (c *baseClient) Put(path string, body string, params map[string]string, headers map[string]string) (*http.Response, error) {
	return c.doRequest(http.MethodPut, path, body, params, headers)
}

func (c *baseClient) HandleError(res *http.Response) error {
	var apiErr api.Error
	err := json.NewDecoder(res.Body).Decode(&apiErr)
	if err != nil {
		return fmt.Errorf("failed to decode error response: %w", err)
	}
	return fmt.Errorf("API error: %s (code: %s)", apiErr.Error, apiErr.ErrorCode)
}
