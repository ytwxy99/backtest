package cqrs

import (
	"context"
	"github.com/sirupsen/logrus"
)

type Asynchronous struct {
	Event string
}

func (asynchronous *Asynchronous) Dispatch (ctx context.Context) {
	logrus.Info("Fetch a new event and dispatch it : ", asynchronous.Event)
	//todo(wangxiaoyu), implement detail
	DispatchResponse<- asynchronous.Event
}