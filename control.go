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
	setup       []SetUpFunc
	modelSetup  []ModelSetupFunc

	IModel M
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
		setup:       config.SetUp,
		modelSetup:  config.ModelSetup,
		IModel:      IModel,
	}
	buildEnv.SetWireFunc(v.SetUp)
	return v
}

func (v *vmanager[T]) SetUp(component vugu.Builder) {
	if c, ok := component.(IView); ok {
		c.setControl(v)
		m := c.getModel()
		for _, setup := range v.modelSetup {
			setup(m)
		}
	}
	if c, ok := component.(IDynamicView); ok {
		c.setViewFactory(v.viewFactory)
	}
	for _, setup := range v.setup {
		setup(component)
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
