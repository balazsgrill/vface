package vface

import "github.com/vugu/vugu"

type DynamicView struct {
	View[IModel]
	viewFactory ViewFactory
	inner       vugu.Builder
}

func (v *DynamicView) setViewFactory(viewFactory ViewFactory) {
	v.viewFactory = viewFactory
}

func (v *DynamicView) getView() vugu.Builder {
	if v.inner == nil {
		v.inner = v.viewFactory.CreateView(v.Model)
	}
	return v.inner
}
