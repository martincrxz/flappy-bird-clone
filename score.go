package main

import (
	"fmt"
	"strconv"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var scoreDst sdl.Rect = sdl.Rect{X: width - 60, Y: 10, W: 50, H: 72}

type score struct {
	points int
	font   *ttf.Font
}

func newScore() (*score, error) {
	font, err := ttf.OpenFont("res/font/Anton-Regular.ttf", 56)
	if err != nil {
		return nil, fmt.Errorf("could not open font, %v", err)
	}

	return &score{
		points: 0,
		font:   font,
	}, nil
}

func (sc *score) paint(r *sdl.Renderer) error {
	surface, err := sc.font.RenderUTF8Solid(strconv.Itoa(sc.points), sdl.Color{R: 255, G: 255, B: 255, A: 255})
	if err != nil {
		return fmt.Errorf("could not render points, %v", err)
	}
	defer surface.Free()

	texture, err := r.CreateTextureFromSurface(surface)
	if err != nil {
		return fmt.Errorf("could not create texture, %v", err)
	}
	defer texture.Destroy()

	err = r.Copy(texture, nil, &scoreDst)
	if err != nil {
		return fmt.Errorf("could not copy texture, %v", err)
	}
	return nil
}

func (sc *score) destroy() {
	sc.font.Close()
}

func (sc *score) addPoint() {
	sc.points++
}
