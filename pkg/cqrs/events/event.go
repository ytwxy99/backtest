package events

import "context"

type Event interface {
	DoEvent(ctx context.Context) (string, error)
}
