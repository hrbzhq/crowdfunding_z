package autoupdater

import (
	"context"
	"testing"
	"time"
)

func TestSchedulerRunsCycle(t *testing.T) {
	fetcher := &MockFetcher{}
	analyzer := &MockAnalyzer{}
	updater := &MockUpdater{}

	s := NewScheduler(fetcher, analyzer, updater, 100*time.Millisecond)
	ctx, cancel := context.WithTimeout(context.Background(), 350*time.Millisecond)
	defer cancel()
	s.Start(ctx)
	// wait for scheduler to run at least once
	<-ctx.Done()
	// no assertions needed — ensure no panics and the mock updater prints recommendations
}

// capturingUpdater records the last AnalysisResult it received
type capturingUpdater struct{
	got AnalysisResult
}
func (c *capturingUpdater) Apply(ctx context.Context, result AnalysisResult) error {
	c.got = result
	return nil
}

func TestSchedulerInvokesUpdater(t *testing.T) {
	fetcher := &MockFetcher{}
	analyzer := &MockAnalyzer{}
	updater := &capturingUpdater{}

	s := NewScheduler(fetcher, analyzer, updater, 50*time.Millisecond)
	ctx, cancel := context.WithTimeout(context.Background(), 180*time.Millisecond)
	defer cancel()
	s.Start(ctx)
	<-ctx.Done()
	if len(updater.got.Recommendations) == 0 {
		t.Fatalf("expected updater to receive recommendations, got none")
	}
}
