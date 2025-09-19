package autoupdater

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"fmt"
)

// HTTPFetcher fetches project data from given URLs (expects JSON array of fetched projects)
type HTTPFetcher struct {
	URLs   []string
	Client *http.Client
	UA     string
	Delay  time.Duration
	Retries int
}

func NewHTTPFetcher(urls []string, timeout time.Duration) *HTTPFetcher {
	return &HTTPFetcher{URLs: urls, Client: &http.Client{Timeout: timeout}, UA: "crowdfunding-autoupdater/1.0", Delay: 1 * time.Second, Retries: 2}
}

func (h *HTTPFetcher) Fetch(ctx context.Context) ([]FetchedProject, error) {
	var out []FetchedProject
	for _, u := range h.URLs {
		select {
		case <-ctx.Done():
			return out, ctx.Err()
		default:
		}

		// simple per-URL delay to be polite
		time.Sleep(h.Delay)

		var lastErr error
		for attempt := 0; attempt <= h.Retries; attempt++ {
			req, _ := http.NewRequestWithContext(ctx, "GET", u, nil)
			req.Header.Set("User-Agent", h.UA)
			resp, err := h.Client.Do(req)
			if err != nil {
				lastErr = err
				// exponential backoff
				backoff := time.Duration(100*(1<<attempt)) * time.Millisecond
				time.Sleep(backoff)
				continue
			}
			var list []FetchedProject
			if err := json.NewDecoder(resp.Body).Decode(&list); err == nil {
				out = append(out, list...)
			} else {
				lastErr = fmt.Errorf("decode failed: %w", err)
			}
			resp.Body.Close()
			break
		}
		// log lastErr to stdout (updater/analyzer can decide), but don't fail entire fetch
		if lastErr != nil {
			fmt.Println("http_fetcher: warning, last error for", u, ":", lastErr)
		}
	}
	return out, nil
}
