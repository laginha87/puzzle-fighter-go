package lib

type EventHandler struct {
	eventListeners [][]func(interface{})
}

func NewEventHandler(numOfEvents uint8) *EventHandler {
	eventListeners := make([][]func(interface{}), numOfEvents)
	return &EventHandler{
		eventListeners: eventListeners,
	}
}

func (e *EventHandler) On(event uint8, callback func(interface{})) {
	e.eventListeners[event] = append(e.eventListeners[event], callback)
}

func (e *EventHandler) Trigger(event uint8, payload interface{}) {
	for _, callback := range e.eventListeners[event] {
		callback(payload)
	}
}
