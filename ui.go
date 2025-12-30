package main

import (
	"bytes"
	"image/color"

	e_image "github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/goregular"
)

// ImGui-like colours
var (
	titleBarColor  = color.NRGBA{55, 95, 165, 255}
	windowBgColor  = color.NRGBA{28, 40, 65, 235}
	titleTextColor = color.NRGBA{240, 245, 250, 255}
	bodyTextColor  = color.NRGBA{215, 225, 235, 255}
)

// Fonts
var (
	titleFace text.Face
	bodyFace  text.Face
)

// Load fonts
func loadFonts() {
	// Load generic font with size
	loadFont := func(size float64) text.Face {
		s, _ := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
		return &text.GoTextFace{Source: s, Size: size}
	}

	titleFace = loadFont(22)
	bodyFace = loadFont(20)
}

// Create floating window with stats about current simulation
func createStatsWindow() (*widget.Window, *widget.Text) {
	// Create title container and widget
	title := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			e_image.NewNineSliceColor(titleBarColor),
		),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	title.AddChild(widget.NewText(
		widget.TextOpts.Text("Statistics", &titleFace, titleTextColor),
		widget.TextOpts.Padding(&widget.Insets{Left: 4, Top: 1}),
	))

	// Create text widget. Content is modified later in game's Update()
	statsText := widget.NewText(
		widget.TextOpts.Text("", &bodyFace, bodyTextColor),
		widget.TextOpts.Padding(&widget.Insets{Left: 4, Top: 1}),
	)

	// Create content container for text widget
	content := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			e_image.NewNineSliceColor(windowBgColor),
		),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	content.AddChild(statsText)

	// Create window
	win := widget.NewWindow(
		widget.WindowOpts.TitleBar(title, 26),
		widget.WindowOpts.Contents(content),
		widget.WindowOpts.Draggable(),
	)

	return win, statsText
}

// Create floating window with control scheme
func createControlsWindow() *widget.Window {
	// Create title container and widget
	title := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			e_image.NewNineSliceColor(titleBarColor),
		),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	title.AddChild(widget.NewText(
		widget.TextOpts.Text("Controls", &titleFace, titleTextColor),
		widget.TextOpts.Padding(&widget.Insets{Left: 4, Top: 1}),
	))

	// Create content container for text widget
	content := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			e_image.NewNineSliceColor(windowBgColor),
		),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	content.AddChild(widget.NewText(
		widget.TextOpts.Text(
			`Toggle UI: <p>
Pause simulation: <space>
Randomise grid: <r>
Clear grid: <c>
Increase simulation speed: <up arrow>
Decrease simulation speed: <down arrow>
Draw alive cells: <left mouse button>
Draw dead cells: <right mouse button>`,
			&bodyFace, bodyTextColor),
		widget.TextOpts.Padding(&widget.Insets{Left: 4, Top: 1}),
	))

	// Create window
	win := widget.NewWindow(
		widget.WindowOpts.TitleBar(title, 26),
		widget.WindowOpts.Contents(content),
		widget.WindowOpts.Draggable(),
	)

	return win
}
