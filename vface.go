package vface

import "github.com/vugu/vugu"

type Control interface {
	/* Indicates that the given model is changed by a control */
	Update(IModel, vugu.DOMEvent)
	Close()
	Run(root vugu.Builder) error
}

type IModel interface {
	IsReadonly() bool
}

/* Base struct for the Model */
type Model struct {
	Readonly bool
}

func (m *Model) IsReadonly() bool {
	return m.Readonly
}

type IView interface {
	vugu.Builder
	setControl(Control)
}

type IDynamicView interface {
	IView
	setViewFactory(ViewFactory)
}

func (v *View[_]) setControl(c Control) {
	v.control = c
}

/* Base struct for every Component */
type View[M IModel] struct {
	/* Reference to the Controller */
	control Control
	Model   M

	/* CSS Class */
	Class string
}

func (v *View[T]) Update(event vugu.DOMEvent) {
	if v.control != nil {
		v.control.Update(v.Model, event)
	}
}

type ViewFactory interface {
	CreateView(IModel) vugu.Builder
}

type ViewFactoryFunc func(IModel) vugu.Builder

func (f ViewFactoryFunc) CreateView(v IModel) vugu.Builder {
	return f(v)
}
