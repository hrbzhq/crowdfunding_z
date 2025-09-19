package autoupdater

import (
	"context"
	"fmt"
	"net/http"
	"bytes"
	"encoding/json"
)

// GitHubUpdater creates issues for recommendations when token is provided
type GitHubUpdater struct {
	Token string
	Repo  string // in form owner/repo
	Client *http.Client
}

func NewGitHubUpdater(token, repo string) *GitHubUpdater {
	return &GitHubUpdater{Token: token, Repo: repo, Client: &http.Client{}}
}

func (g *GitHubUpdater) Apply(ctx context.Context, result AnalysisResult) error {
	if g.Token == "" || g.Repo == "" {
		return fmt.Errorf("github updater not configured")
	}
	for _, rec := range result.Recommendations {
		// Create an issue
		body := map[string]string{"title": rec, "body": "Autoupdater suggestion: " + rec}
		b, _ := json.Marshal(body)
		url := fmt.Sprintf("https://api.github.com/repos/%s/issues", g.Repo)
		req, _ := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(b))
		req.Header.Set("Authorization", "token "+g.Token)
		req.Header.Set("Accept", "application/vnd.github+json")
		req.Header.Set("Content-Type", "application/json")
		resp, err := g.Client.Do(req)
		if err != nil {
			continue
		}
		resp.Body.Close()
	}
	return nil
}
