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

var brickBlockDrawable kame.DrawableModelID
var blueBlockDrawable kame.DrawableModelID
var greenBlockDrawable kame.DrawableModelID
var yellowBlockDrawable kame.DrawableModelID
var ballDrawable kame.DrawableModelID
var paddleDrawable kame.DrawableModelID

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

	brickBlockDrawable, err = kame.CreateSprite()
	brickBlockDrawable.LoadTexture("../Texture/block_solid.png")
	brickBlockDrawable.SetTintColor(mgl32.Vec3{1, 1, 1})

	blueBlockDrawable, err = kame.CreateSprite()
	blueBlockDrawable.LoadTexture("../Texture/block.png")
	blueBlockDrawable.SetTintColor(mgl32.Vec3{0.2, 0.6, 1.0})

	greenBlockDrawable, err = kame.CreateSprite()
	greenBlockDrawable.LoadTexture("../Texture/block.png")
	greenBlockDrawable.SetTintColor(mgl32.Vec3{0.0, 0.7, 0.0})

	yellowBlockDrawable, err = kame.CreateSprite()
	yellowBlockDrawable.LoadTexture("../Texture/block.png")
	yellowBlockDrawable.SetTintColor(mgl32.Vec3{0.8, 0.8, 0.4})

	paddleDrawable, err = kame.CreateSprite()
	paddleDrawable.LoadTexture("../Texture/paddle.png")

	ballDrawable, err = kame.CreateSprite()
	ballDrawable.LoadTexture("../Texture/gopher_circle.png")
	if err != nil {
		panic(err)
	}
}

func (ms *MainScene) CreateEntities() {
	CreateDrawableModel()
	level := ReadLevel()
	col := len(level[0])
	row := len(level)

	f := gameWindow.GetCameraFrustum()

	var entities = []*kame.Entity{}
	w, _ := f.NearPlane.GetSize()
	// fmt.Printf("Width:%v\n", f)
	blockWidth := w / float32(col)
	blockHeight := blockWidth * 0.7
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			var trans kame.Component = &TransformComponent{
				position: mgl32.Vec3{
					(f.NearPlane.Min.X() + blockWidth/2) + (float32(c) * blockWidth),
					(f.NearPlane.Max.Y() - blockHeight/2) - (float32(r) * blockHeight),
					-10},
				scale: mgl32.Vec3{blockWidth, blockHeight, 1},
			}
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

	var ball kame.Entity = &Ball{}
	entities = append(entities, &ball)

	var trans kame.Component = &TransformComponent{
		position: mgl32.Vec3{0, f.NearPlane.Min.Y() + 0.5, 0},
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
