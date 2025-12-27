package main

import (
	"image/color"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	gridWidth          = 100
	gridHeight         = 75
	cellSize           = 8
	probInitiallyAlive = 0.8
)

// Game struct to hold current state
type Game struct {
	currentGrid   [][]bool
	nextGrid      [][]bool
	running       bool
	width, height int
}

// Create new game object
func NewGame() *Game {
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
			g1[y][x] = rand.Float64() > probInitiallyAlive
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

// Draw current grid
func (g *Game) Draw(screen *ebiten.Image) {
	for y := range g.height {
		for x := range g.width {
			// Skip current cell if not alive
			if !g.currentGrid[y][x] {
				continue
			}

			// Draw cell
			vector.FillRect(
				screen,
				float32(x*cellSize),
				float32(y*cellSize),
				cellSize,
				cellSize,
				color.White,
				false,
			)
		}
	}
}

// Set internal canvas size
func (g *Game) Layout(_, _ int) (int, int) {
	return 800, 600
}

func main() {
	// Setup game window
	ebiten.SetWindowTitle("Game of Life (Ebitengine)")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	// Run game
	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}
