package main

import (
	"fmt"
	"strings"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ibaykoc/kame"
)

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

//---------- The real deal is down here ----------

var window *kame.Window
var models []kame.DrawableModel
var t float32

func update(timeSinceLastFrame float64) {
	t += 0.01 * float32(timeSinceLastFrame)
	i := window.GetInput()
	if i.GetKeyStat(kame.KeyLeftAlt) == kame.Press && i.GetKeyStat(kame.KeyF4) == kame.Press || i.GetKeyStat(kame.KeyEscape) == kame.Press {
		window.Close()
	}
}

func draw(drawer *kame.Drawer) {
	for i, model := range models {
		drawer.DrawAt(model, mgl32.Translate3D(-3+float32(i*3), 0, 0).Mul4(mgl32.HomogRotate3D(t, mgl32.Vec3{0, 1, 0})))
	}
}

func onDropFile(filePath string) {
	if !strings.HasSuffix(filePath, ".obj") {
		return
	}
	objModel, err := kame.LoadOBJ(filePath, "../Texture/gopher_smooth_ball.png")
	if err != nil {
		panic(err)
	}
	models = append(models, objModel)
	fmt.Printf("Total loaded model: %d\n", len(models))
}
