package vface

import "github.com/vugu/vugu"

func DefaultViewFactory(v IModel) vugu.Builder {
	if m2, ok := v.(ITexteditModel); ok {
		v := &Textedit{
			DefaultValue: "Empty text",
		}
		v.Model = m2
		return v
	}
	if m2, ok := v.(ITable); ok {
		v := &TableView{}
		v.Model = m2
		return v
	}
	return nil
}
