package main

import (
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ibaykoc/kame"
)

var window *kame.Window
var models []kame.DrawableModel

func main() {
	var err error

	window, err = kame.TurnOn(update, draw)
	if err != nil {
		panic(err)
	}
	defer kame.TurnOff()

	window.SetOnDropFileFunc(onDropFile)

	for !window.WannaClose {
		window.DoMagic()
	}
}

func update(timeSinceLastFrame float64) {
	i := window.GetInput()
	if i.GetKeyStat(kame.KeyLeftAlt) == kame.Press && i.GetKeyStat(kame.KeyF4) == kame.Press || i.GetKeyStat(kame.KeyEscape) == kame.Press {
		window.Close()
	}
}

var t float32

func draw(drawer *kame.Drawer) {
	t += 0.01
	for i, model := range models {
		drawer.DrawAt(model, mgl32.Translate3D(-3+float32(i*3), 0, 0).Mul4(mgl32.HomogRotate3D(t, mgl32.Vec3{0, 1, 0})))
	}
}

func onDropFile(filePath string) {
	fmt.Println(filePath)
	objModel, err := kame.LoadOBJ(filePath, "../Texture/gopher_smooth_ball.png")
	if err != nil {
		panic(err)
	}
	models = append(models, objModel)
	fmt.Printf("Total loaded model: %d\n", len(models))
}
