package engine

import "sort"

type UI struct {
	components []UIComponent

	disposed bool
}

func NewUI(components ...UIComponent) *UI {
	return &UI{
		components: components,
	}
}

func (ui *UI) Draw() []Image {

	sort.Slice(ui.components, func(i, j int) bool {
		return ui.components[i].Depth() < ui.components[j].Depth()
	})

	imgs := make([]Image, len(ui.components))
	for i, comp := range ui.components {
		imgs[i] = comp.Draw()
	}

	return imgs
}

func (ui *UI) Dispose() {
	ui.disposed = true
}

type UIComponent interface {
	Draw() Image
	Depth() int
}

type BaseUIComponent struct {
	depth int
}

func (buic *BaseUIComponent) Depth() int {
	return buic.depth
}
