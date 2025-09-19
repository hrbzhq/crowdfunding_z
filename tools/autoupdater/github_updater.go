package autoupdater

import (
	"context"
	"fmt"
	"net/http"
	"bytes"
	"encoding/json"
	"time"
	"strings"
)

// GitHubUpdater creates issues for recommendations when token is provided
type GitHubUpdater struct {
	Token string
	Repo  string // in form owner/repo
	Client *http.Client
	BaseURL string // e.g. https://api.github.com ; can be overridden in tests
}

func NewGitHubUpdater(token, repo string) *GitHubUpdater {
	return &GitHubUpdater{Token: token, Repo: repo, Client: &http.Client{}, BaseURL: "https://api.github.com"}
}

func (g *GitHubUpdater) Apply(ctx context.Context, result AnalysisResult) error {
	if g.Token == "" || g.Repo == "" {
		return fmt.Errorf("github updater not configured")
	}
	var anySuccess bool
	for _, rec := range result.Recommendations {
		// Create an issue with a few retries and exponential backoff
		body := map[string]string{"title": rec, "body": "Autoupdater suggestion: " + rec}
		b, _ := json.Marshal(body)
	url := fmt.Sprintf("%s/repos/%s/issues", strings.TrimRight(g.BaseURL, "/"), g.Repo)
		var lastErr error
		for attempt := 0; attempt < 3; attempt++ {
			req, _ := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(b))
			req.Header.Set("Authorization", "token "+g.Token)
			req.Header.Set("Accept", "application/vnd.github+json")
			req.Header.Set("Content-Type", "application/json")
			resp, err := g.Client.Do(req)
			if err != nil {
				lastErr = err
			} else {
				// treat 2xx as success
				if resp.StatusCode >= 200 && resp.StatusCode < 300 {
					anySuccess = true
					resp.Body.Close()
					break
				}
				lastErr = fmt.Errorf("unexpected status: %d", resp.StatusCode)
				resp.Body.Close()
			}
			// backoff
			time.Sleep(time.Duration(100*(1<<attempt)) * time.Millisecond)
		}
		if lastErr != nil {
			fmt.Println("github_updater: warning, last error for recommendation:", rec, lastErr)
		}
	}
	if !anySuccess {
		return fmt.Errorf("no successful GitHub issue creations")
	}
	return nil
}
