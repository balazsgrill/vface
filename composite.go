package vface

type IMapModel interface {
	IModel
	GetByKey(key string) IModel
	SetByKey(key string, value IModel)
}

type IList[M IModel] interface {
	IModel
	Size() uint
	Get(uint) M
	Update(uint, M)
	Remove(uint)
	Insert(uint, M)
}
