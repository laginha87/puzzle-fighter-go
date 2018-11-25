// +build js,wasm

package html

import "syscall/js"

type Document struct {
	doc js.Value
}

func NewDocument(el js.Value) Document {
	return Document{
		doc: el,
	}
}

var document = NewDocument(js.Global().Get("document"))

func (d Document) getElementById(id string) js.Value {
	return d.doc.Call("getElementById", id)
}
