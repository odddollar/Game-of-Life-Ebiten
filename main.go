package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	gWidth             = 1000
	gHeight            = 750
	cSize              = 8
	probInitiallyAlive = 0.2
)

// Game struct to hold current state
type Game struct {
	currentGrid           [][]bool
	nextGrid              [][]bool
	running               bool
	gridWidth, gridHeight int
	cellSize              int
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

	// Calculate dimensions of image to draw
	imageWidth := gWidth * cSize
	imageHeight := gHeight * cSize

	// Create new game struct in running state
	g := &Game{
		currentGrid: g1,
		nextGrid:    g2,
		running:     true,
		gridWidth:   gWidth,
		gridHeight:  gHeight,
		cellSize:    cSize,
		image:       ebiten.NewImage(imageWidth, imageHeight),
		pixels:      make([]byte, imageWidth*imageHeight*4),
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
	x += g.gridWidth
	x %= g.gridWidth
	y += g.gridHeight
	y %= g.gridHeight

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

	// Draw current grid
	for gridY := range g.gridHeight {
		for gridX := range g.gridWidth {
			// Skip current cell if not alive
			if !g.currentGrid[gridY][gridX] {
				continue
			}

			// Calculate position of cell in image
			px := gridX * g.cellSize
			py := gridY * g.cellSize

			for cy := range g.cellSize {
				row := ((py+cy)*(g.gridWidth*g.cellSize) + px) * 4
				for cx := range g.cellSize {
					i := row + cx*4
					g.pixels[i] = 0xff
					g.pixels[i+1] = 0xff
					g.pixels[i+2] = 0xff
					g.pixels[i+3] = 0xff
				}
			}
		}
	}

	// Draw pixels to screen
	g.image.WritePixels(g.pixels)
	screen.DrawImage(g.image, nil)
}

// Set internal canvas size
func (g *Game) Layout(_, _ int) (int, int) {
	return g.gridWidth * g.cellSize, g.gridHeight * g.cellSize
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
