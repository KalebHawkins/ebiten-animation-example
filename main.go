package main

import (
	"bytes"
	"image"
	"log"
	"time"

	_ "embed"
	"image/color"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	scrWidth  = 320 // Screen Width
	scrHeight = 240 // Screen Height
)

// Embed the sprite sheet providing a byteslice of the image file.
//
//go:embed images/fox_sprite_sheet.png
var foxSpriteSheet []byte

type Game struct {
	foxSprite *sprite
	ticker    time.Ticker
	animation int
}

// Each update frame we increment the Game struct's inc variable.
func (g *Game) Update() error {
	g.foxSprite.animations[g.animation].Update()
	return nil
}

func (g *Game) Draw(scr *ebiten.Image) {
	scr.Fill(color.RGBA{0, 255, 128, 255})
	op := &ebiten.DrawImageOptions{}
	// Move the origin of the sprite to the center of the frame.
	op.GeoM.Translate(-float64(g.foxSprite.animations[g.animation].frameWidth)/2, -float64(g.foxSprite.animations[g.animation].frameHeight)/2)
	// Move the sprite to the center of the screen.
	op.GeoM.Translate(scrWidth/2, scrHeight/2)
	// op.ColorM.Translate(255, 0, 0, 128)

	g.foxSprite.animations[g.animation].Play(g.foxSprite.image, scr, op)
}

func (g *Game) Layout(outWidth, outHeight int) (int, int) {
	return scrWidth, scrHeight
}

func main() {
	// Decode our embeded image with image decode.
	// If this file is a png the remember to import `_ "image/png"`
	img, _, err := image.Decode(bytes.NewBuffer(foxSpriteSheet))
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the Game struct.
	g := &Game{
		foxSprite: &sprite{
			image: ebiten.NewImageFromImage(img),
			animations: []*Animation{
				{
					startX:      0,
					startY:      0,
					frameWidth:  32,
					frameHeight: 32,
					frameCount:  5,
					frameSpeed:  8,
				},
				{
					startX:      0,
					startY:      32,
					frameWidth:  32,
					frameHeight: 32,
					frameCount:  14,
					frameSpeed:  8,
				},
				{
					startX:      0,
					startY:      64,
					frameWidth:  32,
					frameHeight: 32,
					frameCount:  8,
					frameSpeed:  4,
				},
				{
					startX:      0,
					startY:      96,
					frameWidth:  32,
					frameHeight: 32,
					frameCount:  11,
					frameSpeed:  8,
				},
				{
					startX:      0,
					startY:      128,
					frameWidth:  32,
					frameHeight: 32,
					frameCount:  5,
					frameSpeed:  8,
				},
				{
					startX:      0,
					startY:      160,
					frameWidth:  32,
					frameHeight: 32,
					frameCount:  6,
					frameSpeed:  8,
				},
				{
					startX:      0,
					startY:      192,
					frameWidth:  32,
					frameHeight: 32,
					frameCount:  7,
					frameSpeed:  8,
				},
			},
		},
		ticker:    *time.NewTicker(time.Second * 5),
		animation: 0,
	}

	// Set the window title.
	ebiten.SetWindowTitle("Animation Example")

	go func() {
		for range g.ticker.C {
			if g.animation == len(g.foxSprite.animations)-1 {
				g.animation = 0
			} else {
				g.animation += 1
			}
		}
	}()

	// Start the game.
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
