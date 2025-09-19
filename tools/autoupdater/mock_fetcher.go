package autoupdater

import (
	"context"
)

// MockFetcher returns deterministic sample data for tests and local runs
type MockFetcher struct{}

func (m *MockFetcher) Fetch(ctx context.Context) ([]FetchedProject, error) {
	// Return two sample projects
	return []FetchedProject{
		{Title: "Project Alpha", Description: "AI health device", Goal: 10000, Raised: 2500, Progress: 25, URL: "https://example.com/p/alpha"},
		{Title: "Project Beta", Description: "Open-source hardware", Goal: 5000, Raised: 5000, Progress: 100, URL: "https://example.com/p/beta"},
	}, nil
}
