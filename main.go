package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	gWidth             = 100
	gHeight            = 75
	probInitiallyAlive = 0.2
)

// Game struct to hold current state
type Game struct {
	currentGrid           [][]bool
	nextGrid              [][]bool
	running               bool
	gridWidth, gridHeight int
	image                 *ebiten.Image
	pixels                []byte
}

// Create new game object
func NewGame() *Game {
	// Create two 2D arrays of equal size
	g1 := make([][]bool, gHeight)
	g2 := make([][]bool, gHeight)

	for i := range gHeight {
		g1[i] = make([]bool, gWidth)
		g2[i] = make([]bool, gWidth)
	}

	// Create new game struct in running state
	g := &Game{
		currentGrid: g1,
		nextGrid:    g2,
		running:     true,
		gridWidth:   gWidth,
		gridHeight:  gHeight,
		image:       ebiten.NewImage(gWidth, gHeight),
		pixels:      make([]byte, gWidth*gHeight*4),
	}
	g.initialiseRandomAlivePositions()

	return g
}

// Initialise random alive positions
func (g *Game) initialiseRandomAlivePositions() {
	for y := range g.gridHeight {
		for x := range g.gridWidth {
			g.currentGrid[y][x] = rand.Float64() < probInitiallyAlive
		}
	}
}

// Get alive state at position with wrapping
func (g *Game) isAlive(x, y int) bool {
	x = (x + g.gridWidth) % g.gridWidth
	y = (y + g.gridHeight) % g.gridHeight

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
	for y := range g.gridHeight {
		for x := range g.gridWidth {
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
		for y := range g.gridHeight {
			for x := range g.gridWidth {
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
	// Clear pixel buffer
	clear(g.pixels)

	// Draw current grid to buffer
	for gridY := range g.gridHeight {
		for gridX := range g.gridWidth {
			// Skip current cell if not alive
			if !g.currentGrid[gridY][gridX] {
				continue
			}

			i := (gridY*g.gridWidth + gridX) * 4
			g.pixels[i+0] = 255
			g.pixels[i+1] = 255
			g.pixels[i+2] = 255
			g.pixels[i+3] = 255
		}
	}

	// Write pixel array to image
	g.image.WritePixels(g.pixels)

	// Calculate scaling factors
	sw, sh := screen.Bounds().Dx(), screen.Bounds().Dy()
	scaleX := float64(sw) / float64(g.gridWidth)
	scaleY := float64(sh) / float64(g.gridHeight)

	// Create scaling matrix
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scaleX, scaleY)

	screen.DrawImage(g.image, nil)
}

// Set internal canvas size
func (g *Game) Layout(_, _ int) (int, int) {
	return g.gridWidth, g.gridHeight
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
