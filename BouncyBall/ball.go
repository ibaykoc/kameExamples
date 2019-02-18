package main

import (
	"math/rand"
	"time"

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
	rand.Seed(time.Now().UnixNano())
	w, h := gameWindow.GetSize()
	w /= 50
	h /= 50
	x := (-float32(w) / 2) + (rand.Float32() * float32(w) * 2)
	y := (-float32(h) / 2) + (rand.Float32() * float32(h) * 2)
	// fmt.Printf("%2.2f, %2.2f", x, y)
	b.components = []kame.Component{
		&PositionComponent{mgl32.Vec3{x, y, 0}},
		&VelocityComponent{mgl32.Vec2{rand.Float32(), rand.Float32()}},
		&DrawableComponent{dm},
	}

}

func (b *Ball) GetComponents() *[]kame.Component {
	return &b.components
}
