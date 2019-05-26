package main

import (
	"github.com/ibaykoc/kame"
)

var gameWindow *kame.KGameWindow
var kwindowDrawer2DCon kame.KwindowDrawer2DController

func main() {
	var err error
	err = kame.TurnOn()
	if err != nil {
		panic(err)
	}
	defer kame.TurnOff()

	gameWindow, err = kame.KgameWindowBuilder().
		SetTitle("BouncyBallECS").
		SetSize(600, 600).
		IsResizable().
		BuildWith([]kame.Scene{
			&MainScene{},
		},
		)

	if err != nil {
		panic(err)
	}

	kwindowDrawer2DCon, err = kame.KwindowDrawer2DBuilder().
		SetBackgroundColor(kame.Kcolor{R: 1, G: 1, B: 1, A: 1}).
		BuildTo(gameWindow.ID())
	if err != nil {
		panic(err)
	}

	gameWindow.LockCursor(false)
	gameWindow.EnableCameraMovementControl(true)

	gameWindow.Start()
	for !kame.ShouldClose() {
		kame.DoMagic()
	}
}
