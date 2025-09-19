package autoupdater

import (
    "context"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
    "time"
)

func TestHTTPFetcherFetchesAndDecodes(t *testing.T) {
    // prepare a test server that returns a JSON array of FetchedProject
    srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // verify User-Agent header contains our UA
        ua := r.Header.Get("User-Agent")
        if !strings.Contains(ua, "crowdfunding-autoupdater") {
            t.Fatalf("unexpected User-Agent: %s", ua)
        }
        list := []FetchedProject{{Title: "T", URL: "http://example/1"}}
        _ = json.NewEncoder(w).Encode(list)
    }))
    defer srv.Close()

    hf := NewHTTPFetcher([]string{srv.URL}, 2*time.Second)
    // reduce delays/retries for test speed
    hf.Delay = 1 * time.Millisecond
    hf.Retries = 1

    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    got, err := hf.Fetch(ctx)
    if err != nil {
        t.Fatalf("fetch error: %v", err)
    }
    if len(got) != 1 || got[0].URL != "http://example/1" {
        t.Fatalf("unexpected fetched result: %#v", got)
    }
}
