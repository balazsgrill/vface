package vface

import (
	"fmt"

	"github.com/vugu/vugu"
)

type Control interface {
	/* Indicates that the given model is changed by a control */
	Update(IModel, vugu.DOMEvent)
	Close()
	Run(root vugu.Builder) error
}

type IEventContext interface {
	TriggerChanged(IModel)
}

type vuguEvent struct {
	vugu.DOMEvent
}

func WrapEvent(event vugu.DOMEvent) IEventContext {
	return &vuguEvent{
		DOMEvent: event,
	}
}

func (c *vuguEvent) TriggerChanged(IModel) {
	c.EventEnv().Lock()
	c.EventEnv().UnlockRender()
}

type IModel interface {
	IsReadonly() bool
}

type IClassProvider interface {
	GetClass() string
}

type Identifiable interface {
	Identifier() string
}

/* Base struct for the Model */
type Model struct {
	Readonly bool
	ID       string
}

func (m *Model) IsReadonly() bool {
	return m.Readonly
}

func (m *Model) Identifier() string {
	return m.ID
}

type IView interface {
	vugu.Builder
	setControl(Control)
	getModel() IModel
	getClass() string
}

type IDynamicView interface {
	IView
	setViewFactory(ViewFactory)
}

func (v *View[_]) setControl(c Control) {
	v.control = c
}

func (v *View[T]) getModel() IModel {
	return v.Model
}

/* Base struct for every Component */
type View[M IModel] struct {
	/* Reference to the Controller */
	control Control
	Model   M

	/* CSS Class */
	Class string
}

func (v *View[T]) getClass() string {
	if cp, ok := any(v.Model).(IClassProvider); ok {
		return v.Class + " " + cp.GetClass()
	}
	return v.Class
}

func (v *View[T]) GetKey(iterator int, model IModel) string {
	if model != nil {
		if ide, ok := model.(Identifiable); ok {
			i := ide.Identifier()
			if i != "" {
				return i
			}
		}
	}
	return fmt.Sprint(iterator)
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
