/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-19
* Time: 上午9:55
* */
package class6

import (
	"context"
	"sync"
)

type EventReceiver interface {
	OnEvent(evt Event)
}

type Collector interface {
	Init(evtReceiver EventReceiver) error
	Start(agtCtx context.Context) error
	Stop() error
	Destroy() error
}

type Agent struct {
	collector sync.Map
	evtBuf chan Event
	cancel context.CancelFunc
	ctx context.Context
	stare int
}

