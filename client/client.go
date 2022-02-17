package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/hysem/top-word-service/topword"
)

type Client interface {
	FindTopWords(ctx context.Context, req *topword.FindTopWordsRequest) ([]*topword.WordInfo, error)
}

var _ Client = (*client)(nil)

type client struct {
	host       string
	httpClient *http.Client
}

func NewClient(host string) *client {
	return &client{
		httpClient: http.DefaultClient,
	}
}

func (c *client) FindTopWords(ctx context.Context, req *topword.FindTopWordsRequest) ([]*topword.WordInfo, error) {
	var form url.Values
	form.Add("text", req.Text)

	res, err := c.httpClient.PostForm(fmt.Sprintf("%s/top-words", c.host), form)
	if err != nil {
		return nil, fmt.Errorf("failed to find top words: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid response: %d", res.StatusCode)
	}

	defer res.Body.Close()

	var topWords []*topword.WordInfo
	if err := json.NewDecoder(res.Body).Decode(&topWords); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}
	return topWords, nil
}
