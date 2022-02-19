package vface

import (
	"github.com/vugu/vugu"
	js "github.com/vugu/vugu/js"
)

type TexteditModel struct {
	Model
	Content string
}

type Textedit struct {
	View[*TexteditModel]
	Multiline bool

	editing bool
}

func (c *Textedit) focusLost(event vugu.DOMEvent) {
	updated := event.JSEventTarget().Get("value").String()
	c.editing = !c.editing
	c.Model.Content = updated
	c.Update(event)
}

func setFocus(element js.Value) {
	element.Call("focus")
}

func (c *Textedit) onClick(event vugu.DOMEvent) {
	if !c.Model.Readonly {
		c.editing = true
	}
}
