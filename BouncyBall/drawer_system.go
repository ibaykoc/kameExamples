package main

import (
	"github.com/ibaykoc/kame"
)

type DrawingComponentSet struct {
	*PositionComponent
	*DrawableComponent
}
type DrawingSystem struct {
	filterSets                  []kame.EntityFilterSet
	matchEntityIDToComponentSet map[int]DrawingComponentSet
}

func (ms *DrawingSystem) OnCreate() {
	ms.matchEntityIDToComponentSet = make(map[int]DrawingComponentSet)
}

func (ms *DrawingSystem) CreateEntityFilters() {
	ms.filterSets = []kame.EntityFilterSet{
		kame.EntityFilterSet{
			ComponentType: &PositionComponent{},
			Need:          true,
		},
		kame.EntityFilterSet{
			ComponentType: &DrawableComponent{},
			Need:          true,
		},
	}
}

func (ms *DrawingSystem) GetEntityFilters() []kame.EntityFilterSet {
	return ms.filterSets
}

func (ms *DrawingSystem) OnEntityMatch(entity *kame.Entity, components []*kame.Component) {
	ms.matchEntityIDToComponentSet[(*entity).GetID()] = DrawingComponentSet{
		PositionComponent: (*components[0]).(*PositionComponent),
		DrawableComponent: (*components[1]).(*DrawableComponent),
	}
}

func (ms *DrawingSystem) Draw(kdrawer *kame.KDrawer) {
	for _, e := range ms.matchEntityIDToComponentSet {
		pos := e.position
		dm := e.drawable
		kdrawer.DrawAtPosition(dm, pos)
	}
}
