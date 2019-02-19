package main

import (
	"github.com/ibaykoc/kame"
)

var quadModelID kame.DrawableModelID

type MainScene struct {
	entities         []*kame.Entity
	processorSystems []*kame.ProcessorSystem
	drawerSystems    []*kame.DrawerSystem
}

func (ms *MainScene) CreateEntities() {
	var err error
	quadModelID, err = kame.CreateBuiltInDrawableModelT(kame.Quad, "../Texture/gopher_circle.png")
	if err != nil {
		panic(err)
	}

	entities := make([]*kame.Entity, 1000)
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
