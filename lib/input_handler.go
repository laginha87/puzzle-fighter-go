package lib

import (
	"github.com/laginha87/puzzle-fighter-go/html"
)

type InputHandler struct {
	actions   []*func()
	callbacks map[string][]*func()
}

func NewInputHandler() *InputHandler {
	ih := &InputHandler{
		actions:   make([]*func(), 0),
		callbacks: make(map[string][]*func()),
	}
	html.Document.OnKeyDown(ih.onKeyDown)
	return ih
}

func (ih *InputHandler) On(code string, f func()) {
	callbacks, present := ih.callbacks[code]
	if !present {
		callbacks = make([]*func(), 0)
	}
	ih.callbacks[code] = append(callbacks, &f)
}

func (ih *InputHandler) onKeyDown(code string) {
	callbacks, present := ih.callbacks[code]
	if present {
		ih.actions = append(ih.actions, callbacks...)
	}
}

func (ih *InputHandler) CheckInputs() {
	if len(ih.actions) > 0 {
		now := ih.actions
		ih.actions = make([]*func(), 0)
		for _, callback := range now {
			(*callback)()
		}
	}
}
