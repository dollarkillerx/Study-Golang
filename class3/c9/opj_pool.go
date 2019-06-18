/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-18
* Time: 下午6:41
* */
package c9

import (
	"errors"
	"time"
)

type TargetObj struct {
}

type ObjPoll struct {
	poll chan *TargetObj
}

func NewObjPoll(num int) *ObjPoll {
	poll := &ObjPoll{
		poll: make(chan *TargetObj, num),
	}
	for i := 0; i < num; i++ {
		poll.poll <- &TargetObj{}
	}
	return poll
}

func (o *ObjPoll) Get(timeout time.Duration) (*TargetObj, error) {
	select {
	case data := <-o.poll:
		return data, nil
	case <-time.After(timeout):
		return nil, errors.New("time out")

	}
}

func (o *ObjPoll) Pul(obj *TargetObj) error {
	select {
	case o.poll <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}
