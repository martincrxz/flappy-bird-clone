package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	birdWidth  = 50
	birdHeigth = 43
	initY      = 300
	gravity    = 0.1
)

type bird struct {
	textures []*sdl.Texture
	speed    float64
	y, x     int32
	dead     bool
}

func newBird(r *sdl.Renderer) (*bird, error) {

	var textures []*sdl.Texture
	for i := 1; i <= 4; i++ {
		path := fmt.Sprintf("res/img/bird_frame_%d.png", i)
		texture, err := img.LoadTexture(r, path)
		if err != nil {
			return nil, fmt.Errorf("could not load bird, %v", err)
		}
		textures = append(textures, texture)
	}

	return &bird{
		textures: textures,
		y:        initY,
		dead:     false}, nil
}

func (b *bird) update() {
	b.y += int32(b.speed)
	b.speed -= gravity
	if b.y < 0 {
		b.dead = true
	}
}

func (b *bird) paint(r *sdl.Renderer, t uint64) error {
	rect := &sdl.Rect{W: birdWidth, H: birdHeigth, X: 10, Y: 600 - b.y - int32(birdHeigth/2)}
	texIndex := (t / 10) % uint64(len(b.textures))
	if err := r.Copy(b.textures[texIndex], nil, rect); err != nil {
		return fmt.Errorf("could not copy bird, %v", err)
	}
	return nil
}

func (b *bird) jump() {
	if b.speed < 0 {
		b.speed = 3
	} else {
		b.speed += 2
	}
}

func (b *bird) isDead() bool {
	return b.dead
}

func (b *bird) revive() {
	b.speed = 0
	b.y = initY
	b.dead = false
}
