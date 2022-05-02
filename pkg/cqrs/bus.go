package cqrs

import (
	"context"
)

var Events = make(chan string, 2)

// EventBus is a local event bus that delegates handling of published events
// to all matching registered handlers, in order of registration.
type EventBus interface {
	// Publish the event on the bus.
	Publish(ctx context.Context) error

	// Subscribe the event on the bus
	Subscribe(event interface{}) error
}