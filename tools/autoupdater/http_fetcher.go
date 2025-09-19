package autoupdater

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

// HTTPFetcher fetches project data from given URLs (expects JSON array of fetched projects)
type HTTPFetcher struct {
	URLs []string
	Client *http.Client
}

func NewHTTPFetcher(urls []string, timeout time.Duration) *HTTPFetcher {
	return &HTTPFetcher{URLs: urls, Client: &http.Client{Timeout: timeout}}
}

func (h *HTTPFetcher) Fetch(ctx context.Context) ([]FetchedProject, error) {
	var out []FetchedProject
	for _, u := range h.URLs {
		req, _ := http.NewRequestWithContext(ctx, "GET", u, nil)
		resp, err := h.Client.Do(req)
		if err != nil {
			// skip failed URL
			continue
		}
		var list []FetchedProject
		if err := json.NewDecoder(resp.Body).Decode(&list); err == nil {
			out = append(out, list...)
		}
		resp.Body.Close()
	}
	return out, nil
}
