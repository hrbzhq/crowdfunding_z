package autoupdater

import (
	"context"
	"fmt"
	"strconv"
)

// ScoringAnalyzer scores projects based on simple rules and returns recommendations with scores
type ScoringAnalyzer struct{}

func (s *ScoringAnalyzer) Analyze(ctx context.Context, projects []FetchedProject) (AnalysisResult, error) {
	recs := []string{}
	meta := map[string]string{}
	for _, p := range projects {
		score := 0
		if p.Progress >= 100 {
			score += 100
		} else if p.Progress >= 50 {
			score += 60
		} else if p.Progress >= 20 {
			score += 30
		} else {
			score += 10
		}
		if p.Goal > 5000 {
			score += 5
		}
		recs = append(recs, p.Title+": score="+fmtScore(score))
	}
	meta["count"] = strconv.Itoa(len(projects))
	return AnalysisResult{Recommendations: recs, Metadata: meta}, nil
}

func fmtScore(s int) string {
	return fmt.Sprintf("%d", s)
}
