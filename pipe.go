package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type pipe struct {
	h       int32
	w       int32
	x       int32
	y       int32
	texture *sdl.Texture
	speed   float64
}

func newPipe(r *sdl.Renderer) (*pipe, error) {
	t, e := img.LoadTexture(r, "res/img/pipe.png")
	if e != nil {
		return nil, fmt.Errorf("could not load pipe, %v", e)
	}

	return &pipe{
		h:       100,
		w:       50,
		x:       800,
		texture: t,
		speed:   -0.1,
	}, nil
}

func (p *pipe) update(t uint64) {
	p.x += int32(float64(t) * p.speed)
}

func (p *pipe) paint(r *sdl.Renderer, time uint64) error {
	rect := &sdl.Rect{W: p.w, H: p.h, X: p.x, Y: 600 - p.h}
	flip := sdl.FLIP_NONE

	if err := r.CopyEx(p.texture, nil, rect, 0, nil, flip); err != nil {
		return fmt.Errorf("could not copy pipe, %v", err)
	}

	return nil
}
