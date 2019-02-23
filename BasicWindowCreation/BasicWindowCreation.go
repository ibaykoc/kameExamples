package main

import (
	"strconv"

	"github.com/ibaykoc/kame"
)

func main() {
	var err error
	err = kame.TurnOn()
	if err != nil {
		panic(err)
	}
	defer kame.TurnOff()

	CreateMultipleWindows()

	for !kame.ShouldClose() {
		kame.DoMagic()
	}
}

func CreateMultipleWindows() {

	monW, monH := kame.GetMonitorSize()
	wRow := 2
	wCol := 2
	wW := monW / wCol
	wH := monH / wRow
	for r := 0; r < wRow; r++ {
		for c := 0; c < wCol; c++ {
			_, err := kame.KwindowBuilder().
				SetTitle("Kwindow"+strconv.Itoa(r*wCol+c)).
				SetTargetFPS(60).
				SetPosition(c*wW, r*wH).
				SetSize(wW, wH).
				Build()
			if err != nil {
				panic(err)
			}
		}
	}
}
