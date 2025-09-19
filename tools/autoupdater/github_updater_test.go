package autoupdater

import (
    "context"
    "net/http"
    "net/http/httptest"
    "testing"
    "time"
)

func TestGitHubUpdaterSuccess(t *testing.T) {
    // mock GitHub API returning 201
    srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method != "POST" {
            t.Fatalf("expected POST, got %s", r.Method)
        }
        w.WriteHeader(201)
        w.Write([]byte(`{"number":1}`))
    }))
    defer srv.Close()

    g := &GitHubUpdater{Token: "t", Repo: "owner/repo", Client: &http.Client{Timeout: 2 * time.Second}}
    g.BaseURL = srv.URL

    res := AnalysisResult{Recommendations: []string{"rec1"}}
    ctx := context.Background()
    if err := g.Apply(ctx, res); err != nil {
        t.Fatalf("expected success, got err: %v", err)
    }
}

func TestGitHubUpdaterFailure(t *testing.T) {
    // mock GitHub API returning 500
    srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(500)
    }))
    defer srv.Close()

    g := &GitHubUpdater{Token: "t", Repo: "owner/repo", Client: &http.Client{Timeout: 2 * time.Second}}
    g.BaseURL = srv.URL
    res := AnalysisResult{Recommendations: []string{"rec1"}}
    ctx := context.Background()
    if err := g.Apply(ctx, res); err == nil {
        t.Fatalf("expected error when GitHub returns 500")
    }
}
