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

func (ms *MovingSystem) OnRemoveEntities(entitieIDs []int) {
	for _, de := range entitieIDs {
		for eID := range ms.matchEntityIDToComponentSet {
			if eID == de {
				delete(ms.matchEntityIDToComponentSet, eID)
				break
			}
		}
	}
}

func (ms *MovingSystem) Process(timeSinceLastFrame float32) {
	fmt.Printf("%9.2f FPS\n", 60/timeSinceLastFrame)
	for _, cSet := range ms.matchEntityIDToComponentSet {
		vX, vY := cSet.velocity.Elem()
		cSet.position = cSet.position.Add(mgl32.Vec3{vX * timeSinceLastFrame, vY * timeSinceLastFrame, 0})
		f := gameWindow.GetCameraFrustum()
		wl := f.NearPlane.Min.X()
		wr := f.NearPlane.Max.X()
		wt := f.NearPlane.Max.Y()
		wb := f.NearPlane.Min.Y()
		x, y, z := cSet.position.Elem()
		if x > wr {
			cSet.position = mgl32.Vec3{wr, y, z}
			cSet.velocity = mgl32.Vec2{-vX, vY}
		} else if x < wl {
			cSet.position = mgl32.Vec3{wl, y, z}
			cSet.velocity = mgl32.Vec2{-vX, vY}
		}
		if y > wt {
			cSet.velocity = mgl32.Vec2{vX, -vY}
			cSet.position = mgl32.Vec3{x, wt, z}
		} else if y < wb {
			cSet.velocity = mgl32.Vec2{vX, -vY}
			cSet.position = mgl32.Vec3{x, wb, z}
		}
	}
}
