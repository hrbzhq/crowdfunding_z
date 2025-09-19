package autoupdater

import (
	"context"
	"time"
)

// Fetcher fetches external crowdfunding project data
type Fetcher interface {
	Fetch(ctx context.Context) ([]FetchedProject, error)
}

// Analyzer analyzes fetched data and produces suggestions
type Analyzer interface {
	Analyze(ctx context.Context, projects []FetchedProject) (AnalysisResult, error)
}

// Updater applies updates to local system based on analysis
type Updater interface {
	Apply(ctx context.Context, result AnalysisResult) error
}

// FetchedProject is a normalized representation of an observed project
type FetchedProject struct {
	Title       string
	Description string
	Goal        float64
	Raised      float64
	Progress    float64
	URL         string
}

// AnalysisResult contains recommendations produced by Analyzer
type AnalysisResult struct {
	Recommendations []string
	Metadata        map[string]string
}

// Scheduler periodically runs fetch -> analyze -> update
type Scheduler struct {
	Fetcher  Fetcher
	Analyzer Analyzer
	Updater  Updater
	Interval time.Duration
	stop     chan struct{}
}

// NewScheduler creates a scheduler with interval
func NewScheduler(fetcher Fetcher, analyzer Analyzer, updater Updater, interval time.Duration) *Scheduler {
	return &Scheduler{Fetcher: fetcher, Analyzer: analyzer, Updater: updater, Interval: interval, stop: make(chan struct{})}
}

// Start runs the scheduler in background
func (s *Scheduler) Start(ctx context.Context) {
	ticker := time.NewTicker(s.Interval)
	go func() {
		for {
			select {
			case <-ticker.C:
				// run a single cycle
				projects, err := s.Fetcher.Fetch(ctx)
				if err != nil {
					continue
				}
				res, err := s.Analyzer.Analyze(ctx, projects)
				if err != nil {
					continue
				}
				_ = s.Updater.Apply(ctx, res)
			case <-s.stop:
				ticker.Stop()
				return
			case <-ctx.Done():
				ticker.Stop()
				return
			}
		}
	}()
}

// Stop stops the scheduler
func (s *Scheduler) Stop() {
	close(s.stop)
}
