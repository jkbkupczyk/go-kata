package main

import (
	"log/slog"
	"time"
)

type UserAggregator struct {
	timeout time.Duration
	logger  *slog.Logger
}

func WithTimeout(timeout time.Duration) func(*UserAggregator) error {
	return func(ua *UserAggregator) error {
		ua.timeout = timeout
		return nil
	}
}

func WithLogger(logger *slog.Logger) func(*UserAggregator) error {
	return func(ua *UserAggregator) error {
		ua.logger = logger
		return nil
	}
}

func New(options ...func(*UserAggregator) error) (*UserAggregator, error) {
	aggregator := new(UserAggregator)
	for _, option := range options {
		err := option(aggregator)
		if err != nil {
			return nil, err
		}
	}
	return aggregator, nil
}

func (ua UserAggregator) Aggregate(id int) {

}
