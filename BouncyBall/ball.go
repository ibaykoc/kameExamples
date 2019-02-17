package main

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/ibaykoc/kame"
)

type Ball struct {
	components []kame.Component
	id         int
}

func (b *Ball) ReceiveID(id int) {
	b.id = id
}

func (b Ball) GetID() int {
	return b.id
}

func (b *Ball) CreateComponents() {
	dm, err := kame.CreateDrawableModelT(kame.Quad, "../Texture/gopher_circle.png")
	if err != nil {
		panic(err)
	}
	b.components = []kame.Component{
		&PositionComponent{},
		&VelocityComponent{mgl32.Vec2{0.1, 0.1}},
		&DrawableComponent{dm},
	}

}

func (b *Ball) GetComponents() *[]kame.Component {
	return &b.components
}
