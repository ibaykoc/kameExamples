package main

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ibaykoc/kame"
)

var window *kame.Window
var quadModel kame.DrawableModel
var triangleModel kame.DrawableModel

func main() {
	var err error
	window, err = kame.TurnOn(update, draw)
	if err != nil {
		panic(err)
	}
	defer kame.TurnOff()

	window.EnableCameraFPSControl()
	window.LockCursor()

	quadModel, err = kame.CreateDrawableModelT(kame.Quad, "../Texture/gopher.png")
	if err != nil {
		panic(err)
	}
	triangleModel, err = kame.CreateDrawableModelT(kame.Triangle, "../Texture/gopher.png")
	if err != nil {
		panic(err)
	}

	for !window.WannaClose {
		window.DoMagic()
	}
}

func update(timeSinceLastFrame float32) {
	i := window.GetInput()
	if i.GetKeyStat(kame.KeyEscape) == kame.Press {
		window.Close()
	}
}

var t float32

func draw(drawer *kame.Drawer) {
	t += 0.01
	drawer.DrawAtRotation(quadModel, mgl32.Vec3{0, 0, t})
	drawer.DrawAtPosition(triangleModel, mgl32.Vec3{3, float32(math.Sin(float64(t))), 0})
}
