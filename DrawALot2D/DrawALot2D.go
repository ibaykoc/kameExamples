package main

import (
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ibaykoc/kame"
)

var windowCon kame.KwindowController
var kwindowDrawer2DCon kame.KwindowDrawer2DController
var gopherDrawable kame.Kdrawable2d
var gopherCircleDrawable kame.Kdrawable2d
var blockDrawable kame.Kdrawable2d
var gopherPos []mgl32.Mat4

func main() {
	var err error
	err = kame.TurnOn()
	if err != nil {
		panic(err)
	}
	defer kame.TurnOff()

	// Create window
	windowCon, err = kame.KwindowBuilder().
		SetTitle("DrawALot2D").
		SetProcessInputFunc(processInput).
		SetUpdateFunc(update).
		SetDrawFunc(draw).
		SetTargetFPS(60).
		SetSize(600, 600).
		IsResizable().
		Build()
	if err != nil {
		panic(err)
	}

	// Enable CameraMovement Control
	// For 2d drawer
	// w,a,s,d move up, left, down, right
	// mouse scrool zoom in/out
	// mouse click & drag drag the screen
	windowCon.EnableCameraMovementControl(true)

	// Create Window Drawer 2D
	kwindowDrawer2DCon, err = kame.KwindowDrawer2DBuilder().
		SetBackgroundColor(kame.Kcolor{R: 1, G: 1, B: 1, A: 1}).
		BuildTo(windowCon.ID())

	// Store Texture to drawer
	gopherCircleTextureID, err := kwindowDrawer2DCon.StoreTexturePNG("../Texture/gopher_circle.png")
	if err != nil {
		panic(err)
	}
	gopherTextureID, err := kwindowDrawer2DCon.StoreTexturePNG("../Texture/gopher.png")
	if err != nil {
		panic(err)
	}
	blockTextureID, err := kwindowDrawer2DCon.StoreTexturePNG("../Texture/block.png")
	if err != nil {
		panic(err)
	}

	// Store color to drawer
	whiteCol := kwindowDrawer2DCon.StoreTintColor(kame.Kcolor{R: 1, G: 1, B: 1, A: 1})

	// Store mesh to drawer
	quad, err := kwindowDrawer2DCon.StoreMesh(
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
	gopherCircleDrawable = kame.Kdrawable2d{
		ShaderID:    kwindowDrawer2DCon.DefaultShaderID(),
		MeshID:      quad,
		TextureID:   gopherCircleTextureID,
		TintColorID: whiteCol,
	}
	gopherDrawable = kame.Kdrawable2d{
		ShaderID:    kwindowDrawer2DCon.DefaultShaderID(),
		MeshID:      quad,
		TextureID:   gopherTextureID,
		TintColorID: whiteCol,
	}
	blockDrawable = kame.Kdrawable2d{
		ShaderID:    kwindowDrawer2DCon.DefaultShaderID(),
		MeshID:      quad,
		TextureID:   blockTextureID,
		TintColorID: whiteCol,
	}
	// Create position for drawable to draw to
	gopherPos = []mgl32.Mat4{}
	for x := float32(-20); x < 20; x += 0.2 {
		for y := float32(-20); y < 20; y += 0.2 {
			gopherPos = append(gopherPos,
				mgl32.Translate3D(x, y, 0).
					Mul4(mgl32.Scale3D(0.05, 0.05, 1)),
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
		windowCon.Close()
	}
}

func update(timeSinceLastFrame float32) {
	// Just to keep track of the performance
	fmt.Printf("Total Entity: %d,\t%2.2f FPS\n", len(gopherPos), 60/timeSinceLastFrame)
}

func draw(drawer *kame.KwindowDrawer) {
	// Append all drawable to the drawer to draw
	for i, pos := range gopherPos {
		if i%3 == 0 {
			(*drawer).AppendDrawable(gopherDrawable, pos)
		} else if i%3 == 1 {
			(*drawer).AppendDrawable(gopherCircleDrawable, pos)
		} else {
			(*drawer).AppendDrawable(blockDrawable, pos)
		}
	}
}
