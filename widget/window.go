package widget

import (
	"image"

	"github.com/blizzy78/ebitenui/input"
	"github.com/hajimehoshi/ebiten"
)

type Window struct {
	Modal bool

	contents *Container
}

type WindowOpt func(w *Window)

const WindowOpts = windowOpts(true)

type windowOpts bool

func NewWindow(opts ...WindowOpt) *Window {
	w := &Window{}

	for _, o := range opts {
		o(w)
	}

	return w
}

func (o windowOpts) Contents(c *Container) WindowOpt {
	return func(w *Window) {
		w.contents = c
	}
}

func (o windowOpts) Modal() WindowOpt {
	return func(w *Window) {
		w.Modal = true
	}
}

func (w *Window) SetLocation(rect image.Rectangle) {
	w.contents.SetLocation(rect)
}

func (w *Window) RequestRelayout() {
	w.contents.RequestRelayout()
}

func (w *Window) SetupInputLayer(def input.DeferredSetupInputLayerFunc) {
	if w.Modal {
		w.contents.GetWidget().ElevateToNewInputLayer(&input.Layer{
			DebugLabel: "modal window",
			EventTypes: input.LayerEventTypeAll,
			BlockLower: true,
			FullScreen: true,
		})
	}
}

func (w *Window) Render(screen *ebiten.Image, def DeferredRenderFunc) {
	w.contents.Render(screen, def)
}
