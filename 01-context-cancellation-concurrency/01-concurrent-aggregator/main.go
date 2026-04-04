package main

import (
	"log/slog"
	"os"
	"time"
)

func main() {
	userAggregator, err := New(
		WithTimeout(time.Second*1),
		WithLogger(slog.Default()),
	)
	if err != nil {
		os.Exit(420)
	}

	userAggregator.Aggregate(123)
}
