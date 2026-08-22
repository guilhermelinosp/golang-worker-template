// Package main provides a minimal worker entry point for golang-worker-template.
// Replace the work function with your actual background job logic
// (e.g., Kafka consumer, SQS polling, scheduled tasks, etc.).
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle OS signals for graceful shutdown.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	// Start worker loop.
	done := make(chan struct{})
	go func() {
		defer close(done)
		workerLoop(ctx)
	}()

	fmt.Println("worker started")

	select {
	case <-sig:
		fmt.Println("shutdown signal received, stopping worker...")
		cancel()
	case <-done:
		fmt.Println("worker finished")
	}

	// Wait for worker to clean up.
	select {
	case <-done:
	case <-time.After(10 * time.Second):
		fmt.Fprintln(os.Stderr, "worker shutdown timed out")
	}
}

func workerLoop(ctx context.Context) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker loop exiting")
			return
		case t := <-ticker.C:
			if err := doWork(t); err != nil {
				fmt.Fprintf(os.Stderr, "work error: %v\n", err)
			}
		}
	}
}

func doWork(t time.Time) error {
	fmt.Printf("tick at %s\n", t.Format(time.RFC3339))
	return nil
}
