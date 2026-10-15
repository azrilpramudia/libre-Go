package openlibrary

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL:    baseURL,
		httpClient: &http.Client{Timeout: 5 * time.Second},
	}
}

type SearchResult struct {
	Docs []struct {
		Title      string   `json:"title"`
		AuthorName []string `json:"author_name"`
		ISBN       []string `json:"isbn"`
		CoverID    int      `json:"cover_i"`
		Key        string   `json:"key"`
	} `json:"docs"`
}

func (c *Client) SearchByTitle(ctx context.Context, title string) (*SearchResult, error) {
	q := url.Values{}
	q.Set("title", title)
	q.Set("limit", "5")

	reqURL := fmt.Sprintf("%s/search.json?%s", c.baseURL, q.Encode())

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("openlibrary returned status %d", resp.StatusCode)
	}

	var result SearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func CoverURL(coverID int) string {
	if coverID == 0 {
		return ""
	}
	return fmt.Sprintf("https://covers.openlibrary.org/b/id/%d-L.jpg", coverID)
}
