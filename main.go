package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	gWidth             = 100
	gHeight            = 75
	rWidth             = 1280
	rHeight            = 960
	wWidth             = 1000
	wHeight            = 750
	probInitiallyAlive = 0.2
	nSteppingSpeed     = 20
	xSteppingSpeed     = 1
)

func main() {
	// Setup game window
	ebiten.SetWindowTitle("Game of Life (Ebiten) (Press <p> to toggle UI)")
	ebiten.SetWindowSize(wWidth, wHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowIcon([]image.Image{loadWindowIcon()})

	// Run game
	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}
