package vface

import (
	"github.com/vugu/vugu"
)

type SelectorModel struct {
	Model
	Options   []string
	Labels    []string
	Selection string
}

type Selector struct {
	View[*SelectorModel]
}

func (c *Selector) handleChange(event vugu.DOMEvent) {
	newVal := event.PropString("target", "value")
	c.Model.Selection = newVal
	c.Update(event)
}
