package vface

type IList interface {
	IModel
	GetElements() []IModel
}

type List struct {
	Model
	Elements []IModel
}

func (l *List) GetElements() []IModel {
	return l.Elements
}

type Composite struct {
	View[IList]
}
