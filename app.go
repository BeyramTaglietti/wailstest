package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"path/filepath"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) ReadFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(data)
}

func (a *App) ToBlackAndWhite(path string) string {
	file, err := os.Open(path)
	if err != nil {
		return ""
	}
	defer file.Close()

	_, _, err = image.DecodeConfig(file)
	if err != nil {
		return ""
	}

	file.Seek(0, 0)

	img, _, err := image.Decode(file)
	if err != nil {
		return ""
	}

	grayImg := image.NewGray(img.Bounds())
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			originalColor := img.At(x, y)
			grayColor := color.GrayModel.Convert(originalColor)
			grayImg.Set(x, y, grayColor)
		}
	}

	originalExt := filepath.Ext(path)
	outputPath := strings.TrimSuffix(path, originalExt) + "_bw.jpg"

	output, err := os.Create(outputPath)
	if err != nil {
		return ""
	}
	defer output.Close()

	// Encode with JPEG quality setting
	opt := jpeg.Options{
		Quality: 90,
	}
	err = jpeg.Encode(output, grayImg, &opt)
	if err != nil {
		return ""
	}

	return outputPath
}
