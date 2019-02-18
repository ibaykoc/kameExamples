package main

import (
	"fmt"

	"github.com/go-gl/mathgl/mgl32"

	"github.com/ibaykoc/kame"
)

type MovingComponentSet struct {
	*PositionComponent
	*VelocityComponent
}
type MovingSystem struct {
	filterSets                  []kame.EntityFilterSet
	matchEntityIDToComponentSet map[int]MovingComponentSet
}

func (ms *MovingSystem) OnCreate() {
	ms.matchEntityIDToComponentSet = make(map[int]MovingComponentSet)
}

func (ms *MovingSystem) CreateEntityFilters() {
	ms.filterSets = []kame.EntityFilterSet{
		kame.EntityFilterSet{
			ComponentType: &PositionComponent{},
			Need:          true,
		},
		kame.EntityFilterSet{
			ComponentType: &VelocityComponent{},
			Need:          true,
		},
	}
}

func (ms *MovingSystem) GetEntityFilters() []kame.EntityFilterSet {
	return ms.filterSets
}

func (ms *MovingSystem) OnEntityMatch(entity *kame.Entity, components []*kame.Component) {
	ms.matchEntityIDToComponentSet[(*entity).GetID()] = MovingComponentSet{
		PositionComponent: (*components[0]).(*PositionComponent),
		VelocityComponent: (*components[1]).(*VelocityComponent),
	}
}

func (ms *MovingSystem) Process(timeSinceLastFrame float32) {
	fmt.Printf("%9.2f FPS\n", 60/timeSinceLastFrame)
	for _, e := range ms.matchEntityIDToComponentSet {
		vX, vY := e.velocity.Elem()
		e.position = e.position.Add(mgl32.Vec3{vX * timeSinceLastFrame, vY * timeSinceLastFrame, 0})
		ww, wh := gameWindow.GetSize()
		ww = ww / 50
		wh = wh / 50
		wr := float32(ww) / 2
		wl := -float32(ww) / 2
		wt := float32(wh) / 2
		wb := -float32(wh) / 2
		if e.position.X() > wr {
			e.velocity = mgl32.Vec2{-vX, vY}
			e.position = mgl32.Vec3{wr, e.position.Y(), 0}
		} else if e.position.X() < wl {
			e.velocity = mgl32.Vec2{-vX, vY}
			e.position = mgl32.Vec3{wl, e.position.Y(), 0}
		}
		if e.position.Y() > wt {
			e.velocity = mgl32.Vec2{vX, -vY}
			e.position = mgl32.Vec3{e.position.X(), wt, 0}
		} else if e.position.Y() < wb {
			e.velocity = mgl32.Vec2{vX, -vY}
			e.position = mgl32.Vec3{e.position.X(), wb, 0}
		}
	}
}
