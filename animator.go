package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Animator interface{}

type Animation struct {
	startX      int
	startY      int
	frameWidth  int
	frameHeight int
	frameCount  int
	frameSpeed  int
	incrementor int
}

// NewAnimation(sX, sY, fGapX, fGapY, fWidth, fHeight, fCount, fSpeed int) is a convience function for create a new Animation.
func NewAnimation(sX, sY, fGapX, fGapY, fWidth, fHeight, fCount, fSpeed int) *Animation {
	return &Animation{
		startX:      sX,
		startY:      sY,
		frameWidth:  fWidth,
		frameHeight: fHeight,
		frameCount:  fCount,
		frameSpeed:  fSpeed,
		incrementor: 0,
	}
}

// Update the incrementor each update frame.
func (a *Animation) Update() {
	a.incrementor++
}

// Play is used to play the animation.
func (a *Animation) Play(src *ebiten.Image, dst *ebiten.Image, op *ebiten.DrawImageOptions) {
	// This section provides the speed of the animation.
	//
	// For example:
	// g.inc = 1
	// frameSpeed = 14
	// frameCount = 5
	//
	// (1 / 14) % 5 = 0
	// (2 / 14) % 5 = 0
	// ...
	// (70 / 14) % 5 = 1
	// ...
	// (140 / 14) % 5 = 2
	// ... and so on.
	//
	// This denotes that the animation frame will change every 70 frames.
	// At 60 frames per second the animation would update every 70 / 60 = 1.16666 seconds.
	i := (a.incrementor / a.frameSpeed) % a.frameCount
	sx, sy := a.startX+i*a.frameWidth, a.startY

	// The draw subimage specified the rectangle of the spritesheet to draw. This plugs in the
	// the values from the previous calculations.

	// Taking our (140/14) % 5 example.
	// i = 2
	//
	// sx = 0 + 2 * 32 == 2 * 32 == 64
	// sy = 32
	// The we add the frame width and frame height to sx and sy.
	dst.DrawImage(src.SubImage(image.Rect(sx, sy, sx+a.frameWidth, sy+a.frameHeight)).(*ebiten.Image), op)
}
