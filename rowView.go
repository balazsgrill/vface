package vface

import "github.com/vugu/vugu"

/*
<tr>
for m in c.Model.Cells():
	if m.header: <th> else: <td>
		<DynamicView>
	</th> or </td> respectively
</tr>
*/
func (c *RowView) Build(vgin *vugu.BuildIn) (vgout *vugu.BuildOut) {

	vgout = &vugu.BuildOut{}

	tr := &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "tr", Attr: []vugu.VGAttribute(nil)}
	tr.AddAttrInterface("class", c.Class)
	vgout.Out = append(vgout.Out, tr) // root for output
	for vgiterkeyt, m := range c.Model.Cells() {
		var td *vugu.VGNode
		if c.isHeader(m) {
			td = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "th", Attr: []vugu.VGAttribute(nil)}
		} else {
			td = &vugu.VGNode{Type: vugu.VGNodeType(3), Namespace: "", Data: "td", Attr: []vugu.VGAttribute(nil)}
		}
		tr.AppendChild(td)
		vgcompKey := vugu.MakeCompKey(0x40F300827B5F8D20^vgin.CurrentPositionHash(), c.GetKey(vgiterkeyt, m))
		// ask BuildEnv for prior instance of this specific component
		vgcomp, _ := vgin.BuildEnv.CachedComponent(vgcompKey).(*DynamicView)
		if vgcomp == nil {
			// create new one if needed
			vgcomp = new(DynamicView)
			vgin.BuildEnv.WireComponent(vgcomp)
		}
		vgin.BuildEnv.UseComponent(vgcompKey, vgcomp) // ensure we can use this in the cache next time around
		vgcomp.Model = m
		vgout.Components = append(vgout.Components, vgcomp)
		td.AppendChild(&vugu.VGNode{Component: vgcomp})
	}
	return vgout
}
