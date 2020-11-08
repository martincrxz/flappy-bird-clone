package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	width  = 800
	height = 600
)

var titleDst sdl.Rect = sdl.Rect{X: width/2 - 300, Y: 100, W: 600, H: 400}

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Printf("could not initialize SDL, %v\n", err)
	}
	defer sdl.Quit()

	if err := ttf.Init(); err != nil {
		fmt.Printf("could not initialize ttf, %v\n", err)
	}
	defer ttf.Quit()

	window, renderer, err := sdl.CreateWindowAndRenderer(width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Printf("could not create window, %v\n", err)
	}
	defer window.Destroy()

	scene, err := newScene(renderer)
	if err != nil {
		fmt.Printf("could not create new sceen, %v", err)
	}
	defer scene.destroy()

	sdl.PumpEvents()

	err = drawTitle(renderer, "Flappy Bird")
	if err != nil {
		fmt.Printf("could not print title, %v\n", err)
	}

	time.Sleep(2 * time.Second)

	events := make(chan sdl.Event)
	sceenErrc := scene.run(events, renderer)
	runtime.LockOSThread()

	for {
		select {
		case events <- sdl.WaitEvent():
		case err = <-sceenErrc:
			if err != nil {
				fmt.Printf("runtime error, %v\n", err)
			}
			return
		}
	}
}

func drawTitle(r *sdl.Renderer, text string) error {
	r.Clear()

	font, err := ttf.OpenFont("res/font/Lobster-Regular.ttf", 148)
	if err != nil {
		return fmt.Errorf("could not open font, %v", err)
	}
	defer font.Close()

	surface, err := font.RenderUTF8Solid(text, sdl.Color{R: 255, G: 255, B: 255, A: 255})
	if err != nil {
		return fmt.Errorf("could not render title, %v", err)
	}
	defer surface.Free()

	texture, err := r.CreateTextureFromSurface(surface)
	if err != nil {
		return fmt.Errorf("could not create texture, %v", err)
	}
	defer texture.Destroy()

	err = r.Copy(texture, nil, &titleDst)
	if err != nil {
		return fmt.Errorf("could not copy texture, %v", err)
	}

	r.Present()

	return nil
}
