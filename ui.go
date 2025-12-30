package main

import (
	"image/color"

	e_image "github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

// Create floating window with stats about current simulation
func createStatsWindow() (*widget.Window, *widget.Text) {
	titleFace := loadFont(22)
	bodyFace := loadFont(20)

	// Create title container and widget
	title := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			e_image.NewNineSliceColor(color.NRGBA{70, 70, 70, 255}),
		),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	title.AddChild(widget.NewText(
		widget.TextOpts.Text("Statistics", &titleFace, color.White),
		widget.TextOpts.Padding(&widget.Insets{Left: 4, Top: 1}),
	))

	// Create text widget. Content is modified later in game's Update()
	statsText := widget.NewText(
		widget.TextOpts.Text("", &bodyFace, color.White),
		widget.TextOpts.Padding(&widget.Insets{Left: 4, Top: 1}),
	)

	// Create content container for text widget
	content := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			e_image.NewNineSliceColor(color.NRGBA{40, 40, 40, 220}),
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
