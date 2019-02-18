package main

import (
	"github.com/ibaykoc/kame"
)

var gameWindow *kame.GameWindow

func main() {
	var err error
	gameWindow, err = kame.GameOn2D([]kame.Scene{&MainScene{}})
	if err != nil {
		panic(err)
	}

	gameWindow.Start()
	for !gameWindow.WannaClose {
		gameWindow.DoMagic()
	}
}
