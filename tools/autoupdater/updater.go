package autoupdater

import (
	"context"
	"fmt"
)

// MockUpdater applies analysis results locally (e.g., logs or writes suggestions)
type MockUpdater struct{}

func (u *MockUpdater) Apply(ctx context.Context, result AnalysisResult) error {
	// For now, just print recommendations — in real system this might create PRs, issues, or adjust parameters
	fmt.Println("Autoupdater recommendations:")
	for _, r := range result.Recommendations {
		fmt.Println("-", r)
	}
	return nil
}
