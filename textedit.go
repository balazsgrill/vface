package vface

import (
	"github.com/vugu/vugu"
	js "github.com/vugu/vugu/js"
)

type TexteditModel struct {
	Model
	Content string
}

type ITexteditModel interface {
	IModel
	GetContent() string
	SetContent(string, vugu.DOMEvent)
}

var _ ITexteditModel = &TexteditModel{}

func (m *TexteditModel) GetContent() string {
	return m.Content
}

func (m *TexteditModel) SetContent(updated string, _ vugu.DOMEvent) {
	m.Content = updated
}

type Textedit struct {
	View[ITexteditModel]
	Multiline    bool
	DefaultValue string
	DatalistID   string

	editing bool
}

func (c *Textedit) focusLost(event vugu.DOMEvent) {
	updated := event.JSEventTarget().Get("value").String()
	c.editing = !c.editing
	c.Model.SetContent(updated, event)
	c.Update(event)
}

func setFocus(element js.Value) {
	element.Call("focus")
}

func (c *Textedit) onClick(event vugu.DOMEvent) {
	if !c.Model.IsReadonly() {
		c.editing = true
	}
}

func (c *Textedit) displayContent() string {
	var value string
	if c.Model != nil {
		value = c.Model.GetContent()
	}
	if value == "" {
		value = c.DefaultValue
	}
	return value
}

func (c *Textedit) displayClass() string {
	var value string
	var defaultclass string
	if c.Model != nil {
		value = c.Model.GetContent()
	}
	if value == "" {
		defaultclass = " textedit-default"
	}
	return c.Class + defaultclass
}

func (c *Textedit) onKey(event vugu.DOMEvent, kind KeyEventKind) {
	if keyhandler, ok := c.Model.(KeyHandler); ok {
		jsevent := event.JSEvent()
		key := jsevent.Get("key").String()
		keyhandler.HandleKey(KeyEvent{
			DOMEvent: event,
			Kind:     kind,
			Key:      key,
		})
	}
}
