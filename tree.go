package vface

type ITree interface {
	IModel

	Item() IModel
	IsExpanded() bool
	SetExpanded(event IEventContext, expanded bool)
	Children() []ITree
}
