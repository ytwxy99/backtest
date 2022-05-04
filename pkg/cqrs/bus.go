package cqrs

import (
	"context"
)

var BusEvents = make(chan string)
var DispatchResponse = make(chan string)

// EventBus is a local event bus that delegates handling of published events
// to all matching registered handlers, in order of registration.
type EventBus interface {
	// Publish the event on the bus.
	Publish(ctx context.Context) error

	// Subscribe the event on the bus
	Subscribe(event interface{}) error
}

// AsynchronousDispatch asynchronous dispatch handler
type AsynchronousDispatch interface {
	Dispatch(ctx context.Context) error
}
