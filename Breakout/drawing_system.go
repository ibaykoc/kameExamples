package main

import (
	"github.com/go-gl/mathgl/mgl32"

	"github.com/ibaykoc/kame"
)

type DrawingSystemComponentSet struct {
	TransformComponent
	*DrawableComponent
}

type DrawingSystem struct {
	entityIDtoComponentSet map[int]DrawingSystemComponentSet
	entityFilter           []kame.EntityFilterSet
}

func (ds *DrawingSystem) OnCreate() {
	ds.entityIDtoComponentSet = make(map[int]DrawingSystemComponentSet)
}

func (ds *DrawingSystem) OnEntitiesAdded(entities []*kame.Entity) {
	for _, e := range entities {
		hasTransform := false
		var t TransformComponent
		hasDrawable := false
		var d *DrawableComponent
		for _, c := range (*e).GetComponentPointers() {
			if hasTransform && hasDrawable {
				break
			}
			dComp, dOk := (*c).(*DrawableComponent)
			if dOk {
				d = dComp
				hasDrawable = true
				continue
			}
			tComp, tOk := (*c).(TransformComponent)
			if tOk {
				t = tComp
				hasTransform = true
			}
		}
		if hasTransform && hasDrawable {
			ds.entityIDtoComponentSet[(*e).GetID()] = DrawingSystemComponentSet{
				TransformComponent: t,
				DrawableComponent:  d,
			}
		}
	}
}

func (ds *DrawingSystem) Draw(kdrawer *kame.KwindowDrawer) {
	for _, cset := range ds.entityIDtoComponentSet {
		(*kdrawer).AppendDrawable(cset.drawable,
			mgl32.Translate3D(cset.TransformComponent.GetPosition().Elem()).
				Mul4(mgl32.Scale3D(cset.TransformComponent.GetScale().Elem())))
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
