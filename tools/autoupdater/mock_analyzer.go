package autoupdater

import (
	"context"
	"strconv"
)

// MockAnalyzer produces simple recommendations based on project progress
type MockAnalyzer struct{}

func (m *MockAnalyzer) Analyze(ctx context.Context, projects []FetchedProject) (AnalysisResult, error) {
	recs := []string{}
	meta := map[string]string{}
	for _, p := range projects {
		if p.Progress < 50 {
			recs = append(recs, "Consider reworking the campaign description for: "+p.Title)
		} else if p.Progress >= 100 {
			recs = append(recs, "Promote the success story for: "+p.Title)
		}
	}
	meta["count"] = strconv.Itoa(len(projects))
	return AnalysisResult{Recommendations: recs, Metadata: meta}, nil
}
