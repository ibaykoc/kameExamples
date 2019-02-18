package main

import (
	"github.com/ibaykoc/kame"
)

type MainScene struct {
	entities         []kame.Entity
	processorSystems []kame.ProcessorSystem
	drawerSystems    []kame.DrawerSystem
}

func (ms *MainScene) CreateEntities() {
	entities := make([]kame.Entity, 100)
	for i := 0; i < len(entities); i++ {
		entities[i] = &Ball{}
	}
	ms.entities = entities
}
func (ms *MainScene) GetEntities() *[]kame.Entity {
	return &ms.entities
}

func (ms *MainScene) CreateProcessorSystems() {
	ms.processorSystems = []kame.ProcessorSystem{
		&MovingSystem{},
	}
}

func (ms *MainScene) CreateDrawerSystems() {
	ms.drawerSystems = []kame.DrawerSystem{
		&DrawingSystem{},
	}
}
func (ms *MainScene) GetProcessorSystems() *[]kame.ProcessorSystem {
	return &ms.processorSystems
}
func (ms *MainScene) GetDrawerSystems() *[]kame.DrawerSystem {
	return &ms.drawerSystems
}
