package main

import (
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ibaykoc/kame"
)

var kwindowID kame.KwindowID
var kwindowDrawer3DCon kame.KwindowDrawer3DController
var gopherDrawable kame.Kdrawable3d
var gopherPos []mgl32.Mat4

func main() {
	var err error
	err = kame.TurnOn()
	if err != nil {
		panic(err)
	}
	defer kame.TurnOff()

	// Create window
	kwindowID, err = kame.KwindowBuilder().
		SetTitle("Kame").
		SetProcessInputFunc(processInput).
		SetUpdateFunc(update).
		SetDrawFunc(draw).
		SetTargetFPS(60).
		SetSize(600, 600).
		IsResizable().
		Build()

	// Enable CameraMovement Control
	// For 3d drawer
	// w,a,s,d move forward, left, backward, right
	// Mouse movement to look
	// Shift to run
	kwindowID.EnableCameraMovementControl(true)
	kwindowID.LockCursor()

	// Create Window Drawer 2D
	kwindowDrawer3DCon, err = kame.KwindowDrawer3DBuilder().
		SetBackgroundColor(kame.Kcolor{R: 1, G: 1, B: 1, A: 1}).
		BuildTo(kwindowID)

	// Store Texture to drawer
	goperCircleID, err := kwindowDrawer3DCon.StoreTexturePNG("../Texture/gopher_circle.png")
	if err != nil {
		panic(err)
	}

	// Store mesh to drawer
	quad, err := kwindowDrawer3DCon.StoreMesh(
		// Vertex Position
		[]float32{
			-0.5, +0.5, 0.0, //Left top
			+0.5, +0.5, 0.0, //Right top
			-0.5, -0.5, 0.0, //Left bottom
			+0.5, -0.5, 0.0, //Right bottom
		},
		// Vertex UV
		[]float32{
			0.0, 1.0, //Left top
			1.0, 1.0, //Right top
			0.0, 0.0, //Left bottom
			1.0, 0.0, //Right bottom
		},
		[]float32{
			0.0, 0.0, -1.0, //Left top
			0.0, 0.0, -1.0, //Right top
			0.0, 0.0, -1.0, //Left bottom
			0.0, 0.0, -1.0, //Right bottom
		},
		// Element
		[]uint32{
			0, 2, 1, // First triangle
			1, 2, 3, // Second triangle
		},
	)
	if err != nil {
		panic(err)
	}

	// Create drawable to draw
	gopherDrawable = kame.Kdrawable3d{
		ShaderID:  kwindowDrawer3DCon.GetDefaultShaderID(),
		MeshID:    quad,
		TextureID: goperCircleID,
	}

	// Create position for drawable to draw to
	gopherPos = []mgl32.Mat4{}
	for x := float32(-200); x < 200; x += 1.5 {
		for y := float32(-200); y < 200; y += 1.5 {
			gopherPos = append(gopherPos,
				mgl32.Translate3D(x, y, 0).
					Mul4(mgl32.Scale3D(1, 1, 1)),
			)
		}
	}

	// Run the loop
	for !kame.ShouldClose() {
		kame.DoMagic()
	}
}

func processInput(windowInput kame.KwindowInput) {
	// Close when user just release escape key
	if windowInput.GetKeyStat(kame.KeyEscape) == kame.JustRelease {
		kwindowID.Close()
	}
}

func update(timeSinceLastFrame float32) {
	// Just to keep track of the performance
	fmt.Printf("Total Entity: %d,\t%2.2f FPS\n", len(gopherPos), 60/timeSinceLastFrame)
}

func draw(drawer *kame.KwindowDrawer) {
	// Append all drawable to the drawer to draw
	for _, pos := range gopherPos {
		(*drawer).AppendDrawable(gopherDrawable, pos)
	}
}
