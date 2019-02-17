package main

import (
	"fmt"

	"github.com/ibaykoc/kame"
)

var gameWindow *kame.GameWindow

func main() {
	var err error
	gameWindow, err = kame.GameOn2D([]kame.Scene{&MainScene{}})
	if err != nil {
		panic(err)
	}
	w, h := gameWindow.GetSize()
	fmt.Printf("W: %d, H: %d\n", w, h)

	for !gameWindow.WannaClose {
		gameWindow.DoMagic()
	}
}
