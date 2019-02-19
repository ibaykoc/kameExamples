package main

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/ibaykoc/kame"
)

type PositionComponent struct {
	position mgl32.Vec3
}
type VelocityComponent struct {
	velocity mgl32.Vec2
}
type DrawableComponent struct {
	drawableModelID kame.DrawableModelID
}
