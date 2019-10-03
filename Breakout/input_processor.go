package main

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/ibaykoc/kame"
)

type InputProcessorComponentSet struct {
	TransformComponent
}

type InputProcessor struct {
	entityIDtoComponentSet map[int]InputProcessorComponentSet
	entityFilter           []kame.EntityFilterSet
}

func (ip *InputProcessor) ProcessInput(windowInput kame.KwindowInput) {
	var hInput float32 = 0
	const speed = 0.1
	if windowInput.GetKeyStat(kame.KeyLeft) == kame.Press {
		hInput--
	} else if windowInput.GetKeyStat(kame.KeyRight) == kame.Press {
		hInput++
	}
	for _, cset := range ip.entityIDtoComponentSet {
		cset.SetPosition(cset.GetPosition().Add(mgl32.Vec3{hInput * speed, 0, 0}))
	}
	if windowInput.GetKeyStat(kame.KeyEscape) == kame.JustRelease {
		gameWindow.Close()
	}
}

func (ip *InputProcessor) OnCreate() {
	ip.entityIDtoComponentSet = make(map[int]InputProcessorComponentSet)
}

func (ip *InputProcessor) OnEntitiesAdded(entities []*kame.Entity) {
	for _, e := range entities {
		for _, c := range (*e).GetComponentPointers() {
			_, tOk := (*c).(*UserControlledTransformComponent)
			if tOk {
				ip.entityIDtoComponentSet[(*e).GetID()] = InputProcessorComponentSet{
					TransformComponent: (*c).(TransformComponent),
				}
				break
			}
		}
	}
}

func (ip *InputProcessor) OnRemoveEntities(entityIDs []int) {
	for _, removedEID := range entityIDs {
		for eID := range ip.entityIDtoComponentSet {
			if eID == removedEID {
				delete(ip.entityIDtoComponentSet, eID)
			}
		}
	}
}
