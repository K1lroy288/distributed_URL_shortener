package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"redirect-service/model"
)

type ShortenerClient interface {
	Resolve(ctx context.Context, shortCode string) (*model.ShortenerResponse, error)
}

type shortenerClient struct {
	baseURL string
	client  *http.Client
}

func NewShortenerClient(baseURL string) *shortenerClient {
	return &shortenerClient{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

func (s *shortenerClient) Resolve(ctx context.Context, shortCode string) (*model.ShortenerResponse, error) {
	url := fmt.Sprintf("%s/short/%s", s.baseURL, shortCode)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("link not found")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("shortener service error: %d", resp.StatusCode)
	}

	var res *model.ShortenerResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	return res, nil
}
