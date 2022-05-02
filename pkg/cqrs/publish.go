package cqrs

import "context"

type PublishBus struct {
	Contract string
	Policy   string
}

// Publish implements the method of the bus.EventBus interface.
func (PublishBus *PublishBus) Publish(ctx context.Context, event string) error {
	Events <- event
	//todo(wangxiaoyu), define all kinds of enents
	return nil
}
