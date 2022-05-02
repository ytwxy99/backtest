package cqrs

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type SubscribeBus struct{}

// Subscribe implements the method of the bus.EventBus interface.
func (subscribeBus *SubscribeBus) Subscribe() error {
	logrus.Info("back test subscribe bus system is starting !!")

	// set pprof service
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	// handle all kinds of events
	i := 1
	for i < 2 {
		select {
		case event := <-Events:
			logrus.Info("Fetch a new event: ", event)
		}
	}
	return nil
}
