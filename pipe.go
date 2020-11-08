package main

import (
	"fmt"
	"math/rand"
	"sync"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type pipe struct {
	mu       sync.RWMutex
	h        int32
	w        int32
	x        int32
	y        int32
	inverted bool
}

type pipes struct {
	mu      sync.RWMutex
	texture *sdl.Texture
	speed   float64
	pipes   []*pipe
}

func newPipe() (*pipe, error) {
	return &pipe{
		h:        100 + rand.Int31n(300),
		w:        50,
		x:        800,
		inverted: rand.Float32() < 0.5,
	}, nil
}

func newPipes(r *sdl.Renderer) (*pipes, error) {
	t, e := img.LoadTexture(r, "res/img/pipe.png")
	if e != nil {
		return nil, fmt.Errorf("could not load pipe, %v", e)
	}
	return &pipes{
		texture: t,
		speed:   -2,
		pipes:   nil,
	}, nil
}

func (p *pipe) update(s float64) {
	p.mu.Lock()
	p.x += int32(s)
	p.mu.Unlock()
}

func (ps *pipes) update(t uint64) {
	if t%100 == 0 {
		p, _ := newPipe()
		ps.pipes = append(ps.pipes, p)
	}
	var rem []*pipe
	for _, p := range ps.pipes {
		p.update(ps.speed)
		if p.x+p.w > 0 {
			rem = append(rem, p)
		}
	}
}

func (p *pipe) paint(r *sdl.Renderer, tex *sdl.Texture) error {
	p.mu.RLock()
	rect := &sdl.Rect{W: p.w, H: p.h, X: p.x, Y: 600 - p.h}
	flip := sdl.FLIP_NONE

	if p.inverted {
		rect.Y = 0
		flip = sdl.FLIP_VERTICAL
	}

	if err := r.CopyEx(tex, nil, rect, 0, nil, flip); err != nil {
		return fmt.Errorf("could not copy pipe, %v", err)
	}

	p.mu.RUnlock()
	return nil
}

func (ps *pipes) paint(r *sdl.Renderer, time uint64) error {
	var ret error = nil
	for _, p := range ps.pipes {
		ret = p.paint(r, ps.texture)
	}
	return ret
}

func (ps *pipes) restart() {
	ps.pipes = nil
}

func (ps *pipes) destroy() {
	ps.mu.Lock()
	ps.texture.Destroy()
	ps.mu.Unlock()
}

func (ps *pipes) hit(bird *bird) {
	for _, p := range ps.pipes {
		p.mu.RLock()
		bird.hit(p)
		p.mu.RUnlock()
	}
}
