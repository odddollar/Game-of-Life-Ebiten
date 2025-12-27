package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// White pixel used by alive cells
var whitePixel *ebiten.Image

func initWhitePixel() {
	whitePixel = ebiten.NewImage(1, 1)
	whitePixel.Fill(color.White)
}

// Game struct to hold current state
type Game struct {
	currentGrid   [][]bool
	nextGrid      [][]bool
	running       bool
	width, height int
}

// Create new game object
func NewGame(gridWidth, gridHeight int) *Game {
	return &Game{}
}

// Update current game frame
func (g *Game) Update() error {
	return nil
}

// Draw current frame
func (g *Game) Draw(screen *ebiten.Image) {}

// Set internal canvas size
func (g *Game) Layout(w, h int) (int, int) {
	return 800, 600
}

func main() {
	// Setup game window
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Game of Life (Ebitengine)")

	// Run game
	if err := ebiten.RunGame(NewGame(100, 75)); err != nil {
		panic(err)
	}
}
