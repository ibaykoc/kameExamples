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

func (ds *DrawingSystem) OnCreate() {
	ds.matchEntityIDToComponentSet = make(map[int]DrawingComponentSet)
}

func (ds *DrawingSystem) CreateEntityFilters() {
	ds.filterSets = []kame.EntityFilterSet{
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

func (ds *DrawingSystem) GetEntityFilters() []kame.EntityFilterSet {
	return ds.filterSets
}

func (ds *DrawingSystem) OnEntityMatch(entity *kame.Entity, components []*kame.Component) {
	ds.matchEntityIDToComponentSet[(*entity).GetID()] = DrawingComponentSet{
		PositionComponent: (*components[0]).(*PositionComponent),
		DrawableComponent: (*components[1]).(*DrawableComponent),
	}
}

func (ds *DrawingSystem) OnRemoveEntities(entityIDs []int) {
	for _, de := range entityIDs {
		for e := range ds.matchEntityIDToComponentSet {
			if e == de {
				delete(ds.matchEntityIDToComponentSet, e)
				break
			}
		}
	}
}

func (ds *DrawingSystem) Draw(kdrawer *kame.KDrawer) {
	for _, cSet := range ds.matchEntityIDToComponentSet {
		pos := cSet.position
		dmID := cSet.drawableModelID
		kdrawer.DrawAtPosition(dmID, pos)
	}
}
