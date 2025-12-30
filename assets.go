package main

import (
	"embed"
	"image"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed assets/icon.png
var icon embed.FS

// Load window icon from embedded assets
func loadWindowIcon() image.Image {
	_, iconImg, _ := ebitenutil.NewImageFromFileSystem(icon, "assets/icon.png")

	return iconImg
}
