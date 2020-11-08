package main

import (
	"fmt"
	"sync"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type pipe struct {
	mu sync.RWMutex
	h  int32
	w  int32
	x  int32
	y  int32
}

type pipes struct {
	mu      sync.RWMutex
	texture *sdl.Texture
	speed   float64
	pipes   []*pipe
}

func newPipe() (*pipe, error) {
	return &pipe{
		h: 100,
		w: 50,
		x: 800,
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
	for _, p := range ps.pipes {
		p.update(ps.speed)
	}
}

func (p *pipe) paint(r *sdl.Renderer, tex *sdl.Texture) error {
	rect := &sdl.Rect{W: p.w, H: p.h, X: p.x, Y: 600 - p.h}
	flip := sdl.FLIP_NONE

	if err := r.CopyEx(tex, nil, rect, 0, nil, flip); err != nil {
		return fmt.Errorf("could not copy pipe, %v", err)
	}

	return nil
}

func (ps *pipes) paint(r *sdl.Renderer, time uint64) error {
	var ret error = nil
	for _, p := range ps.pipes {
		ret = p.paint(r, ps.texture)
	}
	return ret
}

func (ps *pipes) destroy() {
	ps.mu.Lock()
	ps.texture.Destroy()
	ps.mu.Unlock()
}
