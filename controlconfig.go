package vface

import "github.com/vugu/vugu"

type ControlConfig struct {
	MountPoint  string
	ViewFactory ViewFactory
	SetUp       []SetUpFunc
}

type SetUpFunc func(component vugu.Builder)

func NewDefaultConfig() *ControlConfig {
	return &ControlConfig{
		MountPoint:  "#vugu_mount_point",
		ViewFactory: ViewFactoryFunc(DefaultViewFactory),
	}
}

func (c *ControlConfig) WithMountPoint(mountpoint string) *ControlConfig {
	c.MountPoint = mountpoint
	return c
}

type cascadedViewFactory struct {
	base     ViewFactory
	addition ViewFactory
}

func (vf *cascadedViewFactory) CreateView(m IModel) vugu.Builder {
	var result vugu.Builder
	result = vf.base.CreateView(m)
	if result == nil {
		result = vf.addition.CreateView(m)
	}
	return result
}

func (c *ControlConfig) CascadeViewFactory(viewfactory ViewFactory) *ControlConfig {
	c.ViewFactory = &cascadedViewFactory{
		base:     c.ViewFactory,
		addition: viewfactory,
	}
	return c
}

func (c *ControlConfig) WithSetUp(setup SetUpFunc) *ControlConfig {
	c.SetUp = append(c.SetUp, setup)
	return c
}
