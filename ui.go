package main

import (
	"image/color"

	e_image "github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

func createStatsWindow() (*widget.Window, *widget.Text) {
	titleFace, _ := loadFont(22)
	bodyFace, _ := loadFont(20)

	statsText := widget.NewText(
		widget.TextOpts.Text("", &bodyFace, color.White),
		widget.TextOpts.Padding(&widget.Insets{Left: 4, Top: 1}),
	)

	content := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			e_image.NewNineSliceColor(color.NRGBA{40, 40, 40, 220}),
		),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	content.AddChild(statsText)

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

	win := widget.NewWindow(
		widget.WindowOpts.Contents(content),
		widget.WindowOpts.TitleBar(title, 26),
		widget.WindowOpts.Draggable(),
	)

	return win, statsText
}
