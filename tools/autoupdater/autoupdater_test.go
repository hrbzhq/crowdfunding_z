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
