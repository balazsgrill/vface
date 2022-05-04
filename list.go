package vface

type SliceList[M IModel] struct {
	Model
	slice []M
}

var _ IList[ITexteditModel] = &SliceList[ITexteditModel]{}

func (s *SliceList[M]) Size() uint {
	z := len(s.slice)
	if z < 0 {
		return 0
	}
	return uint(z)
}

func (s *SliceList[M]) Get(i uint) M {
	var result M
	if i < s.Size() {
		result = s.slice[i]
	}
	return result
}

func (s *SliceList[M]) Update(i uint, v M) {
	if i < s.Size() {
		s.slice[i] = v
	}
}

func (s *SliceList[M]) Remove(i uint) {
	if i < s.Size() {
		s.slice = append(s.slice[:i], s.slice[i+1:]...)
	}
}

func (s *SliceList[M]) Insert(i uint, v M) {
	if i >= s.Size() {
		s.slice = append(s.slice, v)
	} else {
		var nul M
		s.slice = append(s.slice, nul)
		copy(s.slice[i+1:], s.slice[i:])
		s.slice[i] = v
	}
}
