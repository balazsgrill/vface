package vface

import "github.com/vugu/vugu"

type TreeView struct {
	View[ITree]
}

func (c *TreeView) toogle(event vugu.DOMEvent) {
	c.Model.SetExpanded(WrapEvent(event), !c.Model.IsExpanded())
}