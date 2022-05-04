package vface

import (
	"fmt"

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
	Datalist() []string
}

var _ ITexteditModel = &TexteditModel{}

func (m *TexteditModel) GetContent() string {
	return m.Content
}

func (m *TexteditModel) SetContent(updated string, _ vugu.DOMEvent) {
	m.Content = updated
}

func (m *TexteditModel) Datalist() []string {
	return nil
}

type Textedit struct {
	View[ITexteditModel]
	Multiline    bool
	DefaultValue string

	editing bool
}

func (c *Textedit) focusLost(event vugu.DOMEvent) {
	updated := event.JSEventTarget().Get("value").String()
	c.editing = !c.editing
	c.Model.SetContent(updated, event)
	c.Update(event)
	fmt.Println("Focus lost!")
}

func setFocus(element js.Value) {
	element.Call("focus")
}

func (c *Textedit) onClick(event vugu.DOMEvent) {
	if !c.Model.IsReadonly() {
		c.editing = true
		fmt.Println("Editing!")
	} else {
		fmt.Println("Readonly!")
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

func (c *Textedit) datalistID() string {
	if len(c.Model.Datalist()) > 0 {
		return c.Model.Identifier() + "_datalist"
	}
	return ""
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
