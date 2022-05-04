package vface

import (
	"github.com/vugu/vugu"
	"github.com/vugu/vugu/domrender"
)

type vmanager[M IModel] struct {
	eventenv    vugu.EventEnv
	buildenv    *vugu.BuildEnv
	renderer    *domrender.JSRenderer
	viewFactory ViewFactory

	IModel M
}

type ControlConfig struct {
	MountPoint  string
	ViewFactory ViewFactory
}

func NewDefaultConfig() *ControlConfig {
	return &ControlConfig{
		MountPoint:  "#vugu_mount_point",
		ViewFactory: ViewFactoryFunc(DefaultViewFactory),
	}
}

func NewControl[T IModel](config *ControlConfig, IModel T) Control {
	renderer, err := domrender.New(config.MountPoint)
	if err != nil {
		panic(err)
	}

	buildEnv, err := vugu.NewBuildEnv(renderer.EventEnv())
	if err != nil {
		panic(err)
	}
	v := &vmanager[T]{
		eventenv:    renderer.EventEnv(),
		buildenv:    buildEnv,
		renderer:    renderer,
		viewFactory: config.ViewFactory,
		IModel:      IModel,
	}
	buildEnv.SetWireFunc(v.SetUp)
	return v
}

func (v *vmanager[T]) SetUp(component vugu.Builder) {
	if c, ok := component.(IView); ok {
		c.setControl(v)
	}
	if c, ok := component.(IDynamicView); ok {
		c.setViewFactory(v.viewFactory)
	}
}

func (v *vmanager[T]) Close() {
	v.renderer.Release()
}

func (v *vmanager[T]) Run(root vugu.Builder) error {
	for ok := true; ok; ok = v.renderer.EventWait() {

		buildResults := v.buildenv.RunBuild(root)

		err := v.renderer.Render(buildResults)
		if err != nil {
			return err
		}
	}
	return nil
}

func (v *vmanager[T]) Update(_ IModel, event vugu.DOMEvent) {
	// Currently, simply trigger a reload, if there is no event context. In that case the update is already handled
	if event == nil {
		v.eventenv.Lock()
		v.eventenv.UnlockRender()
	}
}
