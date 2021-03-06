package main

import (
	"fmt"
	"sync"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	birdWidth  = 50
	birdHeigth = 43
	initY      = 300
	initX      = 10
	gravity    = 0.1
)

type bird struct {
	mu       sync.RWMutex
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
		x:        initX,
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
	rect := &sdl.Rect{W: birdWidth, H: birdHeigth, X: b.x, Y: 600 - b.y - int32(birdHeigth/2)}
	texIndex := (t / 10) % uint64(len(b.textures))
	if err := r.Copy(b.textures[texIndex], nil, rect); err != nil {
		return fmt.Errorf("could not copy bird, %v", err)
	}
	return nil
}

func (b *bird) jump() {
	if b.speed < 0 {
		b.speed = 3
	} else if b.speed > 6 {
		b.speed = 6
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
	b.x = initX
	b.dead = false
}

func (b *bird) move(m int) {
	b.x += int32(m)
}

func (b *bird) hit(p *pipe) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.x+birdWidth < p.x {
		return
	}
	if b.x > p.x+p.w {
		return
	}
	if !p.inverted && b.y-birdHeigth/2 > p.h {
		return
	}
	if p.inverted && height-p.h > b.y+birdHeigth/2 {
		return
	}
	b.dead = true
}
