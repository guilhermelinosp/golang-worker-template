package main

import (
	"context"
	"testing"
	"time"
)

func TestDoWork(t *testing.T) {
	now := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	if err := doWork(now); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestWorkerLoop(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	// Run worker loop; it should exit when context is cancelled.
	workerLoop(ctx)
}
