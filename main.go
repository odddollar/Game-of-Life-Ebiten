package main

import (
	"image/color"
	"math/rand/v2"

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
func NewGame(gridWidth, gridHeight, probInitiallyAlive int) *Game {
	// Create two 2D arrays of equal size
	g1 := make([][]bool, gridHeight)
	g2 := make([][]bool, gridHeight)

	for i := range gridHeight {
		g1[i] = make([]bool, gridWidth)
		g2[i] = make([]bool, gridWidth)
	}

	// Initial random positions of first grid
	for y := range gridHeight {
		for x := range gridWidth {
			g1[y][x] = rand.IntN(100) > probInitiallyAlive
		}
	}

	// Create new game struct in running state
	return &Game{
		currentGrid: g1,
		nextGrid:    g2,
		running:     true,
		width:       gridWidth,
		height:      gridHeight,
	}
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
	ebiten.SetWindowTitle("Game of Life (Ebitengine)")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	// Run game
	if err := ebiten.RunGame(NewGame(100, 75, 80)); err != nil {
		panic(err)
	}
}
