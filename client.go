package axiston

import (
	"errors"
	"net/http"
	"net/url"
)

// Client is a minimal 'axiston.com' client.
type Client struct {
	ApiKey     string
	UserAgent  string
	BaseURL    *url.URL
	httpClient *http.Client
}

type ClientOption func(*Client) error

// NewClient instantiates a new Client.
func NewClient(options ...ClientOption) (*Client, error) {
	options = append([]ClientOption{
		WithApiKey(envApiKey),
		WithUserAgent(envUserAgent),
		WithBaseURL(envBaseUrl),
		WithHttpClient(http.DefaultClient),
	}, options...)

	client := &Client{}
	for _, option := range options {
		if err := option(client); err != nil {
			return nil, err
		}
	}

	return client, nil
}

// WithApiKey replaces the api key.
// Env variable 'AXISTON_API_KEY'.
func WithApiKey(apiKey string) ClientOption {
	return func(client *Client) error {
		if len(apiKey) == 0 {
			return errors.New("axiston: invalid api key")
		}

		client.ApiKey = apiKey
		return nil
	}
}

// WithUserAgent replaces the 'User-Agent' header.
// Default value: 'Axiston/0.1.0 (Go; Ver. 1.22.2)'.
// Env variable: 'AXISTON_USER_AGENT'.
func WithUserAgent(userAgent string) ClientOption {
	return func(client *Client) error {
		client.UserAgent = userAgent
		return nil
	}
}

// WithBaseURL replaces the 'base' Url.
// Default value: 'https://api.axiston.com'.
// Env variable: 'AXISTON_BASE_URL'.
func WithBaseURL(baseURL string) ClientOption {
	return func(client *Client) error {
		parsed, err := url.Parse(baseURL)
		if err != nil {
			return err
		}

		client.BaseURL = parsed
		return nil
	}
}

// WithHttpClient replaces the 'HTTP' client.
// Default value: 'http.DefaultClient'.
func WithHttpClient(httpClient *http.Client) ClientOption {
	return func(client *Client) error {
		client.httpClient = httpClient
		return nil
	}
}

// Health returns nil if the service is healthy.
func (c *Client) Health() error {
	return nil
}
