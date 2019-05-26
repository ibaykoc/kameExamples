package main

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/ibaykoc/kame"
)

type TransformComponent struct {
	position mgl32.Vec3
	scale    mgl32.Vec3
}

type DrawableComponent struct {
	drawable kame.Kdrawable2d
}

type BlockComponent struct {
	blockType int
}
