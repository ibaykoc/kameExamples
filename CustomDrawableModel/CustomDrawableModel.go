package main

import (
	"github.com/ibaykoc/kame"
)

var window *kame.Window
var customModel kame.DrawableModel

func main() {
	var err error

	window, err = kame.TurnOn(update, draw)
	if err != nil {
		panic(err)
	}
	defer kame.TurnOff()

	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	customModel, err = kame.CreateDrawableModel0(
		[]float32{
			+0.0, +1.0, 0.0,
			-1.0, -1.0, 0.0,
			+1.0, -1.0, 0.0,
		}, []uint32{
			0, 1, 2,
		},
	)
	if err != nil {
		panic(err)
	}
	for !window.WannaClose {
		window.DoMagic()
	}
}

func update(timeSinceLastFrame float32) {
	i := window.GetInput()
	if i.GetKeyStat(kame.KeyLeftAlt) == kame.Press && i.GetKeyStat(kame.KeyF4) == kame.Press || i.GetKeyStat(kame.KeyEscape) == kame.Press {
		window.Close()
	}
}

func draw(drawer *kame.Drawer) {
	// Drawable model with no defined texture will be using default texture which is purple
	drawer.Draw(customModel)
}
