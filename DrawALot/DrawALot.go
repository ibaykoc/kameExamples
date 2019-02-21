package main

import (
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ibaykoc/kame"
)

var kwindowID kame.KwindowID
var kwindowDrawer2DID kame.KwindowDrawer2DID
var gopherDrawable kame.Kdrawable2d
var gopherPos []mgl32.Mat4

func main() {
	var err error
	err = kame.Initialize()
	if err != nil {
		panic(err)
	}
	kwindowID, err = kame.KwindowBuilder().
		SetTitle("Kame").
		SetProcessInputFunc(processInput).
		SetUpdateFunc(update).
		SetDrawFunc(draw).
		SetTargetFPS(60).
		SetSize(600, 600).
		Build()

	kwindowDrawer2DID, err := kame.KwindowDrawer2DBuilder().
		SetBackgroundColor(mgl32.Vec4{1, 1, 1, 1}).
		BuildTo(kwindowID)

	goperCircleID, err := kwindowDrawer2DID.StoreTexturePNG("../Texture/gopher.png")
	if err != nil {
		panic(err)
	}
	whiteCol := kwindowDrawer2DID.StoreTintColor(kame.Kcolor{R: 1, G: 1, B: 1, A: 1})
	quad, err := kwindowDrawer2DID.StoreMesh(
		[]float32{
			-0.5, +0.5, 0.0, //ltop
			+0.5, +0.5, 0.0, //rtop
			-0.5, -0.5, 0.0, //lbot
			+0.5, -0.5, 0.0, //rbot
		},
		[]float32{
			0.0, 1.0,
			1.0, 1.0,
			0.0, 0.0,
			1.0, 0.0,
		},
		[]uint32{
			0, 2, 1,
			1, 2, 3,
		},
	)
	gopherDrawable = kame.Kdrawable2d{
		Shader:    kwindowDrawer2DID.DefaultShaderID(),
		Mesh:      quad,
		Texture:   goperCircleID,
		TintColor: whiteCol,
	}

	if err != nil {
		panic(err)
	}

	gopherPos = []mgl32.Mat4{}
	for x := float32(-20); x < 20; x += 0.175 {
		for y := float32(-20); y < 20; y += 0.175 {
			gopherPos = append(gopherPos,
				mgl32.Translate3D(x, y, 0).
					Mul4(mgl32.Scale3D(0.05, 0.05, 1)),
			)
		}
	}
	for !kame.ShouldClose() {
		kame.DoMagic()
	}
}

func processInput(i map[kame.Key]kame.KeyAction) {
	if i[kame.KeyEscape] == kame.JustRelease {
		kwindowID.Close()
	}
}

var t = float32(0)

func update(timeSinceLastFrame float32) {

	// t += 0.1 * timeSinceLastFrame
	// ts := float32(math.Sin(float64(t)))
	// tc := float32(math.Cos(float64(t)))
	fmt.Printf("Total Entity: %d,\t%2.2f FPS\n", len(gopherPos), 60/timeSinceLastFrame)
}

func draw(drawer *kame.KwindowDrawer) {
	for _, pos := range gopherPos {
		(*drawer).AppendDrawable(gopherDrawable, pos)
	}
}
