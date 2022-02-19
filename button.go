package vface

import "github.com/vugu/vugu"

type ButtonModel struct {
	Model
	Label  string
	Action func(vugu.DOMEvent)
}

type Button struct {
	View[*ButtonModel]
}

func (c *Button) onClick(event vugu.DOMEvent) {
	if c.Model.Action != nil {
		c.Model.Action(event)
	}
}
