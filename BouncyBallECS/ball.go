package main

import (
	"math/rand"
	"time"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/ibaykoc/kame"
)

type Ball struct {
	components []*kame.Component
	id         int
}

func (b *Ball) ReceiveID(id int) {
	b.id = id
}

func (b Ball) GetID() int {
	return b.id
}

func (b *Ball) CreateComponents() {
	rand.Seed(time.Now().UnixNano())
	w, h := gameWindow.GetSize()
	w /= 50
	h /= 50
	x := (-float32(w) / 2) + (rand.Float32() * float32(w) * 2)
	y := (-float32(h) / 2) + (rand.Float32() * float32(h) * 2)
	xVel := ((rand.Float32() * 2) - 1) / 20
	yVel := ((rand.Float32() * 2) - 1) / 20
	var position kame.Component = &PositionComponent{mgl32.Vec3{x, y, float32(b.id) * 0.0001}}
	var velocity kame.Component = &VelocityComponent{mgl32.Vec2{xVel, yVel}}
	var drawable kame.Component = &DrawableComponent{quadModelID}
	b.components = []*kame.Component{
		&position,
		&velocity,
		&drawable,
	}

}

func (b *Ball) GetComponentPointers() []*kame.Component {
	return b.components
}
