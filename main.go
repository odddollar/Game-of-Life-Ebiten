package main

import (
	"fmt"
	"image/color"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	gridWidth          = 1000
	gridHeight         = 750
	cellSize           = 8
	probInitiallyAlive = 0.2
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

	// Create new game struct in running state
	g := &Game{
		currentGrid: g1,
		nextGrid:    g2,
		running:     true,
		width:       gridWidth,
		height:      gridHeight,
	}
	g.initialiseRandomAlivePositions()

	return g
}

// Initialise random alive positions
func (g *Game) initialiseRandomAlivePositions() {
	for y := range g.height {
		for x := range g.width {
			g.currentGrid[y][x] = rand.Float64() < probInitiallyAlive
		}
	}
}

// Get alive state at position with wrapping
func (g *Game) isAlive(x, y int) bool {
	x += g.width
	x %= g.width
	y += g.height
	y %= g.height

	return g.currentGrid[y][x]
}

// Get number of neighbours around position, with wrapping
func (g *Game) numNeighbours(x, y int) int {
	neighbours := 0

	// Iterate through all spaces around (x, y) co-ordinates
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			// Skip current position
			if i == 0 && j == 0 {
				continue
			}

			// Check alive state of neighbour
			if g.isAlive(x+i, y+j) {
				neighbours++
			}
		}
	}

	return neighbours
}

// Step to next state
func (g *Game) step() {
	// Iterate through each grid position
	for y := range g.height {
		for x := range g.width {
			// If cell has 2 neighbours then leave as is
			// Otherwise make cell alive if it has 3 neighbours
			switch g.numNeighbours(x, y) {
			case 2:
				g.nextGrid[y][x] = g.currentGrid[y][x]
			case 3:
				g.nextGrid[y][x] = true
			default:
				g.nextGrid[y][x] = false
			}
		}
	}

	// Make next grid current one
	// Double buffering
	g.currentGrid, g.nextGrid = g.nextGrid, g.currentGrid
}

// Update current game frame
func (g *Game) Update() error {
	// Update window title with TPS/FPS
	ebiten.SetWindowTitle(fmt.Sprintf(
		"Game of Life (Ebitengine) (FPS: %.2f, TPS: %.2f)",
		ebiten.ActualFPS(),
		ebiten.ActualTPS(),
	))

	// Toggle pause input
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.running = !g.running
	}

	// Randomise grid input
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.initialiseRandomAlivePositions()
	}

	// Clear grid input
	if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		for y := range g.height {
			for x := range g.width {
				g.currentGrid[y][x] = false
			}
		}
	}

	// Update grid
	if g.running {
		g.step()
	}

	return nil
}

// Draw current state to screen
func (g *Game) Draw(screen *ebiten.Image) {
	// Draw current grid
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
	return gridWidth * cellSize, gridHeight * cellSize
}

func main() {
	// Setup game window
	ebiten.SetWindowSize(1000, 750)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	// Run game
	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}
