package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	width := int32(600)
	height := int32(600)

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Printf("could not initialize SDL, %v", err)
	}
	defer sdl.Quit()

	window, _, err := sdl.CreateWindowAndRenderer(width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Printf("could not create window, %v", err)
	}
	defer window.Destroy()

	sdl.PumpEvents()

	time.Sleep(2 * time.Second)
}
