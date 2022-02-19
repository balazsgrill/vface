package main

import (
	"github.com/balazsgrill/vface"
	"github.com/vugu/vugu"
)

var _ vugu.DOMEvent // import fixer

type RootModel struct {
	vface.Model

	vface.TexteditModel
	vface.ButtonModel
	vface.SelectorModel
}

// Root is a Vugu component and implements the vugu.Builder interface.
type Root struct {
	vface.View[*RootModel]
}
