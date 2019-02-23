package main

import (
	"strings"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ibaykoc/kame"
)

var winCon kame.KwindowController
var drawables map[kame.Kdrawable3d]mgl32.Mat4
var drawer3DCon kame.KwindowDrawer3DController
var gopherSmoothBallTexID kame.KtextureID

func main() {
	var err error
	err = kame.TurnOn()
	if err != nil {
		panic(err)
	}
	defer kame.TurnOff()

	winCon, err = kame.KwindowBuilder().
		SetTitle("Drag & Drop OBJ").
		SetTargetFPS(60).
		SetSize(600, 600).
		SetProcessInputFunc(processInput).
		SetDrawFunc(draw).
		SetOnDropFileFunc(onDropFile).
		IsResizable().
		Build()

	drawer3DCon, err = kame.KwindowDrawer3DBuilder().
		SetBackgroundColor(kame.Kcolor{R: 0.25, G: 0.25, B: 0.25, A: 1}).
		BuildTo(winCon.ID())
	if err != nil {
		panic(err)
	}
	drawables = make(map[kame.Kdrawable3d]mgl32.Mat4)
	gopherSmoothBallTexID, err = drawer3DCon.StoreTexturePNG("../Texture/gopher_smooth_ball.png")

	for !kame.ShouldClose() {
		kame.DoMagic()
	}
}

func processInput(windowInput kame.KwindowInput) {
	if windowInput.GetKeyStat(kame.KeyC) == kame.JustRelease {
		winCon.EnableCameraMovementControl(true)
		winCon.LockCursor(true)
	} else if windowInput.GetKeyStat(kame.KeyX) == kame.JustRelease {
		winCon.EnableCameraMovementControl(false)
		winCon.LockCursor(false)
	}
}

func draw(drawer *kame.KwindowDrawer) {
	for dw, tr := range drawables {
		(*drawer).AppendDrawable(dw, tr)
	}
}

func onDropFile(mouseX, mouseY float32, filePath string) {
	if !strings.HasSuffix(filePath, ".obj") {
		return
	}
	meshID, err := drawer3DCon.StoreMeshFromOBJ(filePath)
	if err != nil {
		panic(err)
	}
	mouseWorldPos := drawer3DCon.Camera().ScreenToWorldPos(mgl32.Vec3{mouseX, mouseY, 10})
	drawables[kame.Kdrawable3d{
		ShaderID:  drawer3DCon.DefaultShaderID(),
		MeshID:    meshID,
		TextureID: gopherSmoothBallTexID,
	}] = mgl32.Translate3D(mouseWorldPos.Elem())
}
