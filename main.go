package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/gen2brain/raylib-go/raylib"
)

var screenWidth int32 = 1024
var screenHeight int32 = 768
var zoom float32 = 1

var assets []rl.Texture2D
var assetNames []string
var currAssetIndex = 0

func main() {
	// setting log to just show message
	log.SetFlags(0)
	log.Println("Starting glimpse...")

	// load args
	args := os.Args

	if len(args) < 2 {
		log.Println("no dir or file provided")
		return
	}

	if len(args) > 2 {
		log.Println("too many args provided, only provided a single dir or file")
	}

	// figure out if file or dir
	pathname := args[1]

	fileInfo, err := os.Stat(pathname)
	if err != nil {
		log.Fatalln("error:", err)
	}

	// init the raylib window before loading anything
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.SetTraceLogLevel(rl.LogNone)
	rl.InitWindow(screenWidth, screenHeight, "glimpse")
	defer rl.CloseWindow()

	if fileInfo.IsDir() {
		loadDirAssets(pathname)
	} else {
		loadFileAsset(pathname)
	}

	for !rl.WindowShouldClose() {
		update()
		render()
	}
}

func loadDirAssets(dirPath string) {
	log.Printf("loading dir assets from \"%s\" ...", dirPath)
}

func loadFileAsset(filePath string) {
	log.Printf("loading file asset from \"%s\" ...", filePath)

	asset := rl.LoadTexture(filePath)

	if asset.ID == 0 {
		log.Printf("error: failed to load asset from \"%s\"", filePath)
		return
	}

	assets = append(assets, asset)
	assetNames = append(assetNames, filepath.Base(filePath))
}

func update() {
	zoom += rl.GetMouseWheelMove() * 0.1

	// change screen vars based on actual size of window
	screenWidth = int32(rl.GetScreenWidth())
	screenHeight = int32(rl.GetScreenHeight())
}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Gray)

	// draw app title
	fontSize := int32(30)
	titleText := "glimpse"
	textPos := rl.Vector2{
		X: 20,
		Y: 20,
	}

	rl.DrawTextEx(
		rl.GetFontDefault(),
		titleText,
		textPos,
		float32(fontSize),
		1.5,
		rl.White,
	)

	// draw curr asset
	if len(assets) > 0 {
		currAsset := assets[currAssetIndex]

		// assetPos := rl.Vector2{
		// 	X: float32(screenWidth)/2 - float32(currAsset.Width)/2,
		// 	Y: float32(screenHeight)/2 - float32(currAsset.Height)/2,
		// }

		scaledW := float32(currAsset.Width) * zoom
		scaledH := float32(currAsset.Height) * zoom
		assetPos := rl.Vector2{
			X: float32(screenWidth)/2 - scaledW/2,
			Y: float32(screenHeight)/2 - scaledH/2,
		}

		rl.DrawTextureEx(currAsset, assetPos, 0, zoom, rl.White)

		fontSize := int32(50)
		assetText := assetNames[currAssetIndex]
		textSize := rl.MeasureTextEx(rl.GetFontDefault(), assetText, float32(fontSize), 1.5)
		textPos := rl.Vector2{
			X: float32(screenWidth)/2 - textSize.X/2,
			Y: 50,
		}

		rl.DrawTextEx(
			rl.GetFontDefault(),
			assetText,
			textPos,
			float32(fontSize),
			1.5,
			rl.White,
		)
	}

	rl.EndDrawing()
}
