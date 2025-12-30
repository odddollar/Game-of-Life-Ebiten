package main

import (
	"bytes"
	"fmt"
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
Randomise simulation: <r>
Clear simulation: <c>
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

// Create floating window with width/height inputs and a Set button.
func createSizeWindow(onSet func(width, height int)) (*widget.Window, *widget.TextInput, *widget.TextInput) {
	// Create title container and widget
	title := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			e_image.NewNineSliceColor(titleBarColor),
		),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	title.AddChild(widget.NewText(
		widget.TextOpts.Text("Set Grid Size", &titleFace, titleTextColor),
		widget.TextOpts.Padding(&widget.Insets{Left: 4, Top: 1}),
	))

	// Only allow integers
	intValidation := func(newText string) (bool, *string) {
		if newText == "" {
			return true, nil
		}
		for _, r := range newText {
			if r < '0' || r > '9' {
				return false, nil
			}
		}
		return true, nil
	}

	// Helper that returns both row container and TextInput widget
	makeLabeledInput := func(label string, inputWidth int) (*widget.Container, *widget.TextInput) {
		// Create input widget
		input := widget.NewTextInput(
			widget.TextInputOpts.WidgetOpts(
				widget.WidgetOpts.MinSize(inputWidth, 0),
			),
			widget.TextInputOpts.Face(&bodyFace),
			widget.TextInputOpts.Padding(widget.NewInsetsSimple(4)),
			widget.TextInputOpts.Image(&widget.TextInputImage{
				Idle:     e_image.NewNineSliceColor(color.NRGBA{80, 90, 110, 255}),
				Disabled: e_image.NewNineSliceColor(color.NRGBA{80, 90, 110, 255}),
			}),
			widget.TextInputOpts.Color(&widget.TextInputColor{
				Idle:  bodyTextColor,
				Caret: bodyTextColor,
			}),
			widget.TextInputOpts.Validation(intValidation),
		)

		// Create row container and label widget
		row := widget.NewContainer(
			widget.ContainerOpts.Layout(widget.NewRowLayout(
				widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
				widget.RowLayoutOpts.Spacing(8),
			)),
		)
		row.AddChild(widget.NewText(
			widget.TextOpts.Text(label, &bodyFace, bodyTextColor),
			widget.TextOpts.Padding(&widget.Insets{Left: 2, Top: 4}),
		))
		row.AddChild(input)

		return row, input
	}

	// Make labelled inputs
	widthRow, widthInput := makeLabeledInput("Width:", 100)
	heightRow, heightInput := makeLabeledInput("Height:", 100)

	// Create set button widget
	setButton := widget.NewButton(
		widget.ButtonOpts.Image(&widget.ButtonImage{
			Idle:    e_image.NewNineSliceColor(color.NRGBA{70, 100, 150, 255}),
			Hover:   e_image.NewNineSliceColor(color.NRGBA{90, 120, 170, 255}),
			Pressed: e_image.NewNineSliceColor(color.NRGBA{50, 80, 130, 255}),
		}),
		widget.ButtonOpts.Text("Set", &bodyFace, &widget.ButtonTextColor{
			Idle: bodyTextColor,
		}),
		widget.ButtonOpts.TextPadding(&widget.Insets{
			Left:   16,
			Right:  16,
			Top:    4,
			Bottom: 4,
		}),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			var w, h int
			fmt.Sscanf(widthInput.GetText(), "%d", &w)
			fmt.Sscanf(heightInput.GetText(), "%d", &h)

			// Run passed in function
			onSet(w, h)
		}),
	)

	// Create main content layout
	content := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			e_image.NewNineSliceColor(windowBgColor),
		),
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Spacing(6),
			widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(6)),
		)),
	)
	content.AddChild(widthRow)
	content.AddChild(heightRow)
	content.AddChild(setButton)

	// Create window
	win := widget.NewWindow(
		widget.WindowOpts.TitleBar(title, 26),
		widget.WindowOpts.Contents(content),
		widget.WindowOpts.Draggable(),
	)

	return win, widthInput, heightInput
}
