package vface

import "github.com/vugu/vugu"

type KeyEventKind int

const (
	PRESS KeyEventKind = 0
	DOWN  KeyEventKind = 1
	UP    KeyEventKind = 2
)

type KeyEvent struct {
	vugu.DOMEvent
	Kind KeyEventKind
	Key  string
}

type KeyHandler interface {
	HandleKey(KeyEvent)
}

type KeyFunc func(KeyEvent)

func (f KeyFunc) HandleKey(event KeyEvent) {
	f(event)
}
