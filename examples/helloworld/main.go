package main

import (
	"fmt"

	"github.com/balazsgrill/vface"
	"github.com/vugu/vugu"
)

func main() {
	defer fmt.Printf("Exiting main()\n")

	root := &Root{}
	model := &RootModel{
		TexteditModel: vface.TexteditModel{
			Content: "Text, click to edit me!",
		},
		ButtonModel: vface.ButtonModel{
			Label: "Click me!",
		},
		SelectorModel: vface.SelectorModel{
			Options: []string{
				"op1",
				"op2",
				"op3",
			},
			Labels: []string{
				"No, Select me!",
				"Select me!",
				"Don't select me!",
			},
			Selection: "op2",
		},
	}
	model.ButtonModel.Action = func(vugu.DOMEvent) {
		fmt.Printf("Current seleciton is %s\n", model.SelectorModel.Selection)
		fmt.Printf("Text field content is %s\n", model.TexteditModel.Content)
	}
	root.Model = model
	control := vface.NewControl("#vugu_mount_point", model)
	defer control.Close()

	control.Run(root)
}
