package main

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/ibaykoc/kame"
)

type TransformComponent interface {
	GetPosition() mgl32.Vec3
	SetPosition(mgl32.Vec3)
	GetScale() mgl32.Vec3
	SetScale(mgl32.Vec3)
}

type BasicTransformComponent struct {
	position mgl32.Vec3
	scale    mgl32.Vec3
}

func (btc *BasicTransformComponent) GetPosition() mgl32.Vec3 {
	return btc.position
}

func (btc *BasicTransformComponent) SetPosition(newPosition mgl32.Vec3) {
	btc.position = newPosition
}

func (btc *BasicTransformComponent) GetScale() mgl32.Vec3 {
	return btc.scale
}

func (btc *BasicTransformComponent) SetScale(newScale mgl32.Vec3) {
	btc.position = newScale
}

type UserControlledTransformComponent struct {
	position mgl32.Vec3
	scale    mgl32.Vec3
}

func (btc *UserControlledTransformComponent) GetPosition() mgl32.Vec3 {
	return btc.position
}

func (btc *UserControlledTransformComponent) SetPosition(newPosition mgl32.Vec3) {
	btc.position = newPosition
}

func (btc *UserControlledTransformComponent) GetScale() mgl32.Vec3 {
	return btc.scale
}

func (btc *UserControlledTransformComponent) SetScale(newScale mgl32.Vec3) {
	btc.position = newScale
}

type DrawableComponent struct {
	drawable kame.Kdrawable2d
}

type BlockComponent struct {
	blockType int
}

type InputProcessorComponent struct{}
