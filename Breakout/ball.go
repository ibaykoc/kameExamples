package main

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/ibaykoc/kame"
)

type Ball struct {
	id         int
	components []*kame.Component
}

func (b *Ball) ReceiveID(id int) {
	b.id = id
}

func (b *Ball) GetID() int {
	return b.id
}

func (b *Ball) CreateComponents() {
	var trans kame.Component = &BasicTransformComponent{
		position: mgl32.Vec3{0, 0, -10},
		scale:    mgl32.Vec3{1, 1, 1},
	}
	var draw kame.Component = &DrawableComponent{ballDrawable}
	b.components = []*kame.Component{
		&trans,
		&draw,
	}
}

func (b *Ball) GetComponentPointers() []*kame.Component {
	return b.components
}
