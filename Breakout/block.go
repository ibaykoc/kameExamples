package main

import (
	"github.com/ibaykoc/kame"
)

type Block struct {
	components []*kame.Component
	id         int
}

func (b *Block) ReceiveID(id int) {
	b.id = id
}

func (b Block) GetID() int {
	return b.id
}

func (b *Block) CreateComponents() {

}

func (b *Block) GetComponentPointers() []*kame.Component {
	return b.components
}
