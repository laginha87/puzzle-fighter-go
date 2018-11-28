// +build js,wasm

package html

import "syscall/js"

type HTMLDocument struct {
	doc js.Value
}

func NewDocument(el js.Value) HTMLDocument {
	return HTMLDocument{
		doc: el,
	}
}

var Document = NewDocument(js.Global().Get("document"))

func (d HTMLDocument) getElementById(id string) js.Value {
	return d.doc.Call("getElementById", id)
}

func (d HTMLDocument) OnKeyDown(callback func(code string)) {
	c := func(ev []js.Value) {
		event := ev[0]
		callback(event.Get("code").String())
	}
	d.doc.Set("onkeydown", js.NewCallback(c))
}

func (d HTMLDocument) Log(o interface{}) {
	d.doc.Get("console").Call("log", o)
}
