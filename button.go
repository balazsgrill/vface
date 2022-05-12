package vface

import "github.com/vugu/vugu"

type ButtonModel struct {
	Model
	Label  string
	Action func(IEventContext)
}

type IButtonModel interface {
	IModel
	GetLabel() string
	Pressed(IEventContext)
}

func (m *ButtonModel) GetLabel() string {
	return m.Label
}

func (m *ButtonModel) Pressed(event IEventContext) {
	if m.Action != nil {
		m.Action(event)
	}
}

type Button struct {
	View[IButtonModel]
}

func (c *Button) onClick(event vugu.DOMEvent) {
	c.Model.Pressed(WrapEvent(event))
}
