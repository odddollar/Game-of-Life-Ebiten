package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	gWidth             = 100
	gHeight            = 75
	probInitiallyAlive = 0.2
	nSteppingSpeed     = 20
	xSteppingSpeed     = 1
)

func main() {
	// Setup game window
	ebiten.SetWindowSize(1000, 750)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	// Run game
	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}
