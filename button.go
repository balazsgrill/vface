package vface

import "github.com/vugu/vugu"

type ButtonModel struct {
	Model
	Label  string
	Action func(vugu.DOMEvent)
}

type IButtonModel interface {
	IModel
	GetLabel() string
	Pressed(vugu.DOMEvent)
}

func (m *ButtonModel) GetLabel() string {
	return m.Label
}

func (m *ButtonModel) Pressed(event vugu.DOMEvent) {
	if m.Action != nil {
		m.Action(event)
	}
}

type Button struct {
	View[IButtonModel]
}

func (c *Button) onClick(event vugu.DOMEvent) {
	c.Model.Pressed(event)
}
