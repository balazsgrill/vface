package vface

import "github.com/vugu/vugu"

func DefaultViewFactory(m IModel) vugu.Builder {
	if m2, ok := m.(ITexteditModel); ok {
		v := &Textedit{
			DefaultValue: "Empty text",
		}
		v.Model = m2
		return v
	}
	if m2, ok := m.(ITable); ok {
		v := &TableView{}
		v.Model = m2
		return v
	}
	if m2, ok := m.(IButtonModel); ok {
		v := &Button{}
		v.Model = m2
		return v
	}
	if m2, ok := m.(IList); ok {
		v := &Composite{}
		v.Model = m2
		return v
	}
	return nil
}
