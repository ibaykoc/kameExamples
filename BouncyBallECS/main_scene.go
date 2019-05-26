package main

import (
	"github.com/ibaykoc/kame"
)

var gopherCircleDrawable kame.Kdrawable2d

type MainScene struct {
	entities         []*kame.Entity
	processorSystems []*kame.ProcessorSystem
	drawerSystems    []*kame.DrawerSystem
}

func (ms *MainScene) CreateEntities() {
	var err error
	// Store Texture to drawer
	gopherCircleTextureID, err := kwindowDrawer2DCon.StoreTextureJPG("../Texture/honest.jpg")
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

	gopherCircleDrawable = kame.Kdrawable2d{
		ShaderID:    kwindowDrawer2DCon.DefaultShaderID(),
		MeshID:      quad,
		TextureID:   gopherCircleTextureID,
		TintColorID: whiteCol,
	}

	entities := make([]*kame.Entity, 10000)
	for i := 0; i < len(entities); i++ {
		var b kame.Entity
		b = &Ball{}
		entities[i] = &b
	}

	ms.entities = entities
}
func (ms *MainScene) GetEntityPointers() []*kame.Entity {
	return ms.entities
}

func (ms *MainScene) CreateProcessorSystems() {
	var pSys kame.ProcessorSystem = &MovingSystem{}
	ms.processorSystems = []*kame.ProcessorSystem{
		&pSys,
	}
}

func (ms *MainScene) GetProcessorSystemPointers() []*kame.ProcessorSystem {
	return ms.processorSystems
}

func (ms *MainScene) CreateDrawerSystems() {
	var dSys kame.DrawerSystem = &DrawingSystem{}
	ms.drawerSystems = []*kame.DrawerSystem{
		&dSys,
	}
}

func (ms *MainScene) GetDrawerSystemPointers() []*kame.DrawerSystem {
	return ms.drawerSystems
}
func (ms *MainScene) OnRemoveEntities(entityIDs []int) {
	for _, de := range entityIDs {
		for i, e := range ms.entities {
			if (*e).GetID() == de {
				ms.entities = append(ms.entities[:i], ms.entities[i+1:]...)
				break
			}
		}
	}
}
