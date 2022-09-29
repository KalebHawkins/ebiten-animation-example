package main

import "github.com/hajimehoshi/ebiten/v2"

type Sprite interface{}

type sprite struct {
	image      *ebiten.Image
	animations []*Animation
}
