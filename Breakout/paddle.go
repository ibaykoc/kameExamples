package main

import (
	"github.com/ibaykoc/kame"
)

type Paddle struct {
	id         int
	components []*kame.Component
}

func (p *Paddle) ReceiveID(id int) {
	p.id = id
}

func (p *Paddle) GetID() int {
	return p.id
}

func (p *Paddle) CreateComponents() {
}

func (p *Paddle) GetComponentPointers() []*kame.Component {
	return p.components
}
