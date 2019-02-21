package main

import (
	"github.com/go-gl/mathgl/mgl32"

	"github.com/ibaykoc/kame"
)

type DrawingSystemComponentSet struct {
	*TransformComponent
	*DrawableComponent
}

type DrawingSystem struct {
	entityIDtoComponentSet map[int]DrawingSystemComponentSet
	entityFilter           []kame.EntityFilterSet
}

func (ds *DrawingSystem) OnCreate() {
	ds.entityIDtoComponentSet = make(map[int]DrawingSystemComponentSet)
}

func (ds *DrawingSystem) CreateEntityFilters() {
	ds.entityFilter = []kame.EntityFilterSet{
		kame.EntityFilterSet{
			ComponentType: &TransformComponent{},
			Need:          true,
		},
		kame.EntityFilterSet{
			ComponentType: &DrawableComponent{},
			Need:          true,
		},
	}
}

func (ds *DrawingSystem) GetEntityFilters() []kame.EntityFilterSet {
	return ds.entityFilter
}

func (ds *DrawingSystem) OnEntityMatch(entity *kame.Entity, components []*kame.Component) {
	// fmt.Println("Entity Match")
	ds.entityIDtoComponentSet[(*entity).GetID()] = DrawingSystemComponentSet{
		TransformComponent: (*components[0]).(*TransformComponent),
		DrawableComponent:  (*components[1]).(*DrawableComponent),
	}

}

func (ds *DrawingSystem) Draw(drawer *kame.KDrawer) {
	// fmt.Println("Draw")
	for _, cset := range ds.entityIDtoComponentSet {
		drawer.DrawAt(cset.drawableID,
			mgl32.Translate3D(cset.position.Elem()).
				Mul4(mgl32.Scale3D(cset.TransformComponent.scale.Elem())))
	}
}

func (ds *DrawingSystem) OnRemoveEntities(entityIDs []int) {
	for _, removedEID := range entityIDs {
		for eID := range ds.entityIDtoComponentSet {
			if eID == removedEID {
				delete(ds.entityIDtoComponentSet, eID)
			}
		}
	}
}
