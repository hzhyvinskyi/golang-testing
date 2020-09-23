package cases

import (
	"math/rand"
	"time"
)

type _case struct {
	observerList []Observer
	title        string
	status       string
}

func NewCase(title string, status string) *_case {
	return &_case{
		title: title,
		status: status,
	}
}

func (c *_case) register(observer Observer) {
	c.observerList = append(c.observerList, observer)
}

func (c *_case) deregister(observerToDelete Observer) {
	observerListLength := len(c.observerList)
	for i, observer := range c.observerList {
		if observerToDelete.getID() == observer.getID() {
			c.observerList[observerListLength-1], c.observerList[i] = c.observerList[i], c.observerList[observerListLength-1]
		}
	}
}

func (c *_case) notifyAll() {
	rand.Seed(time.Now().UnixNano())
	p := &Payload{
		ID: uint64(rand.Intn(100000000)),
		Command: "Update Case Owner",
		Title: "Case Updated",
		Timestamp: time.Now(),
	}
	for _, observer := range c.observerList {
		observer.update(p)
	}
}
