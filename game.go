package main

import (
	"fmt"
	"image"
	"image/color"
	"math/rand/v2"

	"github.com/ebitenui/ebitenui"
	e_image "github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Game struct to hold current state
type Game struct {
	// Holds grid and rendering data
	currentGrid           [][]bool
	nextGrid              [][]bool
	gridWidth, gridHeight int
	image                 *ebiten.Image
	pixels                []byte

	// Simulation paused or not
	running bool

	// Speed of simulation
	steppingSpeed     int
	minSteppingSpeed  int
	maxSteppingSpeed  int
	currentFrameCount int

	// Ui data
	ui       *ebitenui.UI
	statsTxt *widget.Text

	// Internal rendering resolution used by Ebiten
	renderingWidth, renderingHeight int
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
		currentGrid:       g1,
		nextGrid:          g2,
		gridWidth:         gWidth,
		gridHeight:        gHeight,
		image:             ebiten.NewImage(gWidth, gHeight),
		pixels:            make([]byte, gWidth*gHeight*4),
		running:           true,
		steppingSpeed:     10,
		minSteppingSpeed:  nSteppingSpeed,
		maxSteppingSpeed:  xSteppingSpeed,
		currentFrameCount: 1,
		renderingWidth:    rWidth,
		renderingHeight:   rHeight,
	}
	g.initialiseRandomAlivePositions()

	// Initialise ui
	g.initUI()

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

// Step to next simulation state
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

// Initialise overlay ui elements
func (g *Game) initUI() {
	ui := &ebitenui.UI{}

	// Create root container
	ui.Container = widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			e_image.NewNineSliceColor(color.NRGBA{0, 0, 0, 0}),
		),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)

	// Assign ui object to game's ui field
	g.ui = ui

	// Create stats window widgets
	statsWin, statsText := createStatsWindow()
	g.statsTxt = statsText

	// Set widow position and size
	statsWin.SetLocation(image.Rect(10, 10, 195, 107))

	// Add floating windows to ui
	ui.AddWindow(statsWin)
}

// Update current game frame
func (g *Game) Update() error {
	// Update ui stats window text
	g.statsTxt.Label = fmt.Sprintf(
		"FPS: %.2f\nTPS: %.2f\nStepping Speed: %d",
		ebiten.ActualFPS(),
		ebiten.ActualTPS(),
		g.steppingSpeed,
	)

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

	// Toggle cells for drawing
	x, y := ebiten.CursorPosition()
	if x >= 0 && x < g.gridWidth && y >= 0 && y < g.gridHeight {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			g.currentGrid[y][x] = true
		}
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
			g.currentGrid[y][x] = false
		}
	}

	// Change stepping speed
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) && g.steppingSpeed > g.maxSteppingSpeed {
		g.steppingSpeed--
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) && g.steppingSpeed < g.minSteppingSpeed {
		g.steppingSpeed++
	}

	// Update grid
	if g.running {
		if g.currentFrameCount%g.steppingSpeed == 0 {
			g.step()
			g.currentFrameCount = 1
		} else {
			g.currentFrameCount++
		}
	}

	g.ui.Update()

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

			// Set pixels in buffer to white
			i := (gridY*g.gridWidth + gridX) * 4
			g.pixels[i+0] = 255
			g.pixels[i+1] = 255
			g.pixels[i+2] = 255
			g.pixels[i+3] = 255
		}
	}

	// Write pixel array to image
	g.image.WritePixels(g.pixels)

	// Calculate scaling factors to maximise image
	sw, sh := screen.Bounds().Dx(), screen.Bounds().Dy()
	scaleX := float64(sw) / float64(g.gridWidth)
	scaleY := float64(sh) / float64(g.gridHeight)

	// Create scaling matrix
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scaleX, scaleY)

	screen.DrawImage(g.image, op)

	// Draw ui on top
	g.ui.Draw(screen)
}

// Set internal canvas size/rendering resolution
func (g *Game) Layout(_, _ int) (int, int) {
	return g.renderingWidth, g.renderingHeight
}
