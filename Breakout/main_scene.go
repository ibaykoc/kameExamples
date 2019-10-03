package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ibaykoc/kame"
)

type MainScene struct {
	entities      []*kame.Entity
	drawerSystems []*kame.DrawerSystem
}

var brickBlockDrawable kame.Kdrawable2d
var blueBlockDrawable kame.Kdrawable2d
var greenBlockDrawable kame.Kdrawable2d
var yellowBlockDrawable kame.Kdrawable2d
var ballDrawable kame.Kdrawable2d
var paddleDrawable kame.Kdrawable2d

func ReadLevel() [][]int {
	lvlFile, err := os.Open("level.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(lvlFile)

	level := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		col := []int{}
		for _, r := range line {
			blockType, _ := strconv.Atoi(string(r))
			col = append(col, blockType)
		}
		level = append(level, col)
	}
	fmt.Printf("%v\n", level)
	return level
}

func CreateDrawableModel() {
	var err error
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
	// Store color to drawer
	whiteCol := kwindowDrawer2DCon.StoreTintColor(kame.Kcolor{R: 1, G: 1, B: 1, A: 1})
	blueCol := kwindowDrawer2DCon.StoreTintColor(kame.Kcolor{R: 0, G: 0, B: 1, A: 1})
	greenCol := kwindowDrawer2DCon.StoreTintColor(kame.Kcolor{R: 0, G: 1, B: 0, A: 1})
	yellowCol := kwindowDrawer2DCon.StoreTintColor(kame.Kcolor{R: 1, G: 1, B: 0, A: 1})

	brickBlockTextureID, err := kwindowDrawer2DCon.StoreTexturePNG("../Texture/block_solid.png")
	if err != nil {
		panic(err)
	}
	brickBlockDrawable = kame.Kdrawable2d{
		ShaderID:    kwindowDrawer2DCon.DefaultShaderID(),
		MeshID:      quad,
		TextureID:   brickBlockTextureID,
		TintColorID: whiteCol,
	}

	blockTextureID, err := kwindowDrawer2DCon.StoreTexturePNG("../Texture/block.png")
	if err != nil {
		panic(err)
	}

	blueBlockDrawable = kame.Kdrawable2d{
		ShaderID:    kwindowDrawer2DCon.DefaultShaderID(),
		MeshID:      quad,
		TextureID:   blockTextureID,
		TintColorID: blueCol,
	}
	greenBlockDrawable = kame.Kdrawable2d{
		ShaderID:    kwindowDrawer2DCon.DefaultShaderID(),
		MeshID:      quad,
		TextureID:   blockTextureID,
		TintColorID: greenCol,
	}
	yellowBlockDrawable = kame.Kdrawable2d{
		ShaderID:    kwindowDrawer2DCon.DefaultShaderID(),
		MeshID:      quad,
		TextureID:   blockTextureID,
		TintColorID: yellowCol,
	}

	paddleTextureID, err := kwindowDrawer2DCon.StoreTexturePNG("../Texture/paddle.png")
	if err != nil {
		panic(err)
	}
	paddleDrawable = kame.Kdrawable2d{
		ShaderID:    kwindowDrawer2DCon.DefaultShaderID(),
		MeshID:      quad,
		TextureID:   paddleTextureID,
		TintColorID: whiteCol,
	}
	ballTextureID, err := kwindowDrawer2DCon.StoreTexturePNG("../Texture/gopher_circle.png")
	if err != nil {
		panic(err)
	}
	ballDrawable = kame.Kdrawable2d{
		ShaderID:    kwindowDrawer2DCon.DefaultShaderID(),
		MeshID:      quad,
		TextureID:   ballTextureID,
		TintColorID: whiteCol,
	}
}

func (ms *MainScene) CreateEntities() {
	CreateDrawableModel()
	level := ReadLevel()
	col := len(level[0])
	row := len(level)
	fmt.Printf("Row:%d, col:%d\n", row, col)
	f := kwindowDrawer2DCon.Camera().Frustum()

	var entities = []*kame.Entity{}
	w, _ := f.NearPlane.GetSize()
	// fmt.Printf("Width:%v\n", f)
	blockWidth := w / float32(col)
	blockHeight := blockWidth * 0.7
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			var trans kame.Component = &TransformComponent{
				position: mgl32.Vec3{
					(f.NearPlane.BottomLeft.X() + blockWidth/2) + (float32(c) * blockWidth),
					(f.NearPlane.TopLeft.Y() - blockHeight/2) - (float32(r) * blockHeight),
					-10},
				scale: mgl32.Vec3{blockWidth, blockHeight, 1},
			}
			o, _ := trans.(TransformComponent)

			println(o.position.X)
			var drawable kame.Component
			switch level[r][c] {
			case 0:
				continue
			case 1:
				drawable = &DrawableComponent{brickBlockDrawable}
			case 2:
				drawable = &DrawableComponent{blueBlockDrawable}
			case 3:
				drawable = &DrawableComponent{greenBlockDrawable}
			case 4:
				drawable = &DrawableComponent{yellowBlockDrawable}
			default:
				drawable = &DrawableComponent{blueBlockDrawable}
			}
			var blockComp kame.Component = &BlockComponent{level[r][c]}
			var block kame.Entity = &Block{
				components: []*kame.Component{
					&trans,
					&drawable,
					&blockComp,
				},
			}
			entities = append(entities, &block)
		}
	}
	println(len(entities))
	var ball kame.Entity = &Ball{}
	entities = append(entities, &ball)

	var trans kame.Component = &TransformComponent{
		position: mgl32.Vec3{0, f.NearPlane.BottomLeft.Y() + 0.5, 0},
		scale:    mgl32.Vec3{2.5, 0.5, 1},
	}
	var draw kame.Component = &DrawableComponent{paddleDrawable}
	var paddle kame.Entity = &Paddle{
		components: []*kame.Component{
			&trans,
			&draw,
		},
	}
	entities = append(entities, &paddle)

	ms.entities = entities
}

func (ms *MainScene) GetEntityPointers() []*kame.Entity {
	return ms.entities
}

func (ms *MainScene) CreateProcessorSystems() {

}

func (ms *MainScene) GetProcessorSystemPointers() []*kame.ProcessorSystem {
	return []*kame.ProcessorSystem{}
}

func (ms *MainScene) CreateDrawerSystems() {
	var ds kame.DrawerSystem = &DrawingSystem{}
	ms.drawerSystems = []*kame.DrawerSystem{
		&ds,
	}
}

func (ms *MainScene) GetDrawerSystemPointers() []*kame.DrawerSystem {
	return ms.drawerSystems
}

func (ms *MainScene) OnRemoveEntities(entityIDs []int) {
	for _, removedID := range entityIDs {
		for i, entity := range ms.entities {
			if removedID == (*entity).GetID() {
				ms.entities = append(ms.entities[:i], ms.entities[i+1:]...)
			}
		}
	}
}
