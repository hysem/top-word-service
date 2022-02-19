package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/hysem/top-word-service/topword"
)

// Client interface for communicating the with top-word-service
type Client interface {
	// FindTopWords finds the top words in the given text using a min heap
	// If two words have the same count then the word will be selected based on the alphabetic order
	FindTopWords(ctx context.Context, req *topword.FindTopWordsRequest) ([]*topword.WordInfo, error)
}

// Make sure that the client implements all methods in Client interface
var _ Client = (*client)(nil)

// client implements the Client interface
type client struct {
	host       string
	httpClient *http.Client
}

// NewClient returns an instance client implementation
func NewClient(host string) *client {
	return &client{
		host:       host,
		httpClient: http.DefaultClient,
	}
}

// FindTopWords finds the top words in the given text using a min heap
// If two words have the same count then the word will be selected based on the alphabetic order
func (c *client) FindTopWords(ctx context.Context, req *topword.FindTopWordsRequest) ([]*topword.WordInfo, error) {
	// Step 1: Prepare the request data
	form := url.Values{}
	form.Add("text", req.Text)

	// Step 2: Send the request
	res, err := c.httpClient.PostForm(fmt.Sprintf("%s/top-words", c.host), form)
	if err != nil {
		return nil, fmt.Errorf("failed to find top words: %v", err)
	}

	// Step 3: If response status code is not 200 then report error
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid response: %d", res.StatusCode)
	}

	// Step 3: Make sure to close the response body
	defer res.Body.Close()

	// Step 4: Unmarshal the response
	var topWords []*topword.WordInfo
	if err := json.NewDecoder(res.Body).Decode(&topWords); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}
	return topWords, nil
}
