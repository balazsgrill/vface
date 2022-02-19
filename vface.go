package vface

import "github.com/vugu/vugu"

type Control interface {
	/* Indicates that the given model is changed by a control */
	Update(IModel, vugu.DOMEvent)
	Close()
	Run(root vugu.Builder) error
}

type IModel interface {
	Identifer() string
}

/* Base struct for the Model */
type Model struct {
	ID       string
	Readonly bool
}

func (m *Model) Identifer() string {
	return m.ID
}

type IView interface {
	setControl(Control)
}

func (v *View[_]) setControl(c Control) {
	v.Control = c
}

/* Base struct for every Component */
type View[M IModel] struct {
	/* Reference to the Controller */
	Control
	Model M

	/* CSS Class */
	Class string
}

func (v *View[T]) Update(event vugu.DOMEvent) {
	if v.Control != nil {
		v.Control.Update(v.Model, event)
	}
}
