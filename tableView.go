package vface

import "github.com/vugu/vugu"

func (c *TableView) Build(vgin *vugu.BuildIn) (vgout *vugu.BuildOut) {

	vgout = &vugu.BuildOut{}

	var vgiterkey interface{}
	_ = vgiterkey
	table := &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "table", Attr: []vugu.VGAttribute(nil)}
	table.AddAttrInterface("class", c.Class)
	vgout.Out = append(vgout.Out, table) // root for output
	{
		for vgiterkeyt, m := range c.Model.Rows() {
			var vgiterkey interface{} = vgiterkeyt
			_ = vgiterkey
			m := m
			_ = m
			{
				vgcompKey := vugu.MakeCompKey(0x7B0BFB805C3E66F^vgin.CurrentPositionHash(), c.GetKey(vgiterkeyt, m))
				// ask BuildEnv for prior instance of this specific component
				vgcomp, _ := vgin.BuildEnv.CachedComponent(vgcompKey).(*RowView)
				if vgcomp == nil {
					// create new one if needed
					vgcomp = new(RowView)
					vgin.BuildEnv.WireComponent(vgcomp)
				}
				vgin.BuildEnv.UseComponent(vgcompKey, vgcomp) // ensure we can use this in the cache next time around
				vgcomp.Model = m
				vgout.Components = append(vgout.Components, vgcomp)
				vgn := &vugu.VGNode{Component: vgcomp}
				table.AppendChild(vgn)
			}
		}
	}
	return vgout
}
