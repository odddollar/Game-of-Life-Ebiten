package main

import (
	"bytes"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/goregular"
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

// Load generic font with size
func loadFont(size float64) (text.Face, error) {
	s, _ := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	return &text.GoTextFace{Source: s, Size: size}, nil
}

func main() {
	// Setup game window
	ebiten.SetWindowTitle("Game of Life (Ebiten) (Press <p> to toggle UI)")
	ebiten.SetWindowSize(wWidth, wHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	// Run game
	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}
