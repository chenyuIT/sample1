package providers

import (
	"github.com/chenyuIT/framework/contracts/event"
	"github.com/chenyuIT/framework/facades"
)

type EventServiceProvider struct {
}

func (receiver *EventServiceProvider) Register() {
	facades.Event.Register(receiver.listen())
}

func (receiver *EventServiceProvider) Boot() {

}

func (receiver *EventServiceProvider) listen() map[event.Event][]event.Listener {
	return map[event.Event][]event.Listener{}
}
