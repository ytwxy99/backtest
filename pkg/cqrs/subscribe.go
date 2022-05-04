package cqrs

import (
	"context"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/ytwxy99/backtest/pkg/database"
	"github.com/ytwxy99/backtest/pkg/utils"
)

type SubscribeBus struct{}

// Subscribe implements the method of the bus.EventBus interface.
func (subscribeBus *SubscribeBus) Subscribe(ctx context.Context) error {
	logrus.Info("back test subscribe bus system is starting !!")

	// set pprof service
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	go func() {
		for {
			time.Sleep(time.Duration(1) * time.Second)
			publishes, err := database.GetAllPublishes(ctx)
			if err != nil {
				logrus.Errorf("get all publishes failed: ", err)
			}

			for _, publish := range publishes {
				if publish.Status == utils.NewPublish {
					BusEvents <- publish.Event
					publish.Status = utils.Published
					publish.UpdatePublish(ctx)
				}
			}
		}
	}()

	// handle all kinds of events
	for {
		select {
		case event := <-BusEvents:
			go func() {
				asynchronous := &Asynchronous{
					Event: event,
				}
				asynchronous.Dispatch(ctx)
			}()

			select {
			case response := <-DispatchResponse:
				logrus.Infof("handle the specified '%s' event success, the response is '%s'.", event, response)
			}
		}
	}
	return nil
}
