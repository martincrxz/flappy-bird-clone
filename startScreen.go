package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/ttf"

	"github.com/veandco/go-sdl2/sdl"
)

type selectedOption string

const (
	startNewGame selectedOption = "startNewGame"
	quitGame                    = "quitGame"
)

var startNewGameDst sdl.Rect = sdl.Rect{X: width/2 - 80, Y: 300, W: 160, H: 40}
var quitGameDst sdl.Rect = sdl.Rect{X: width/2 - 50, Y: 400, W: 100, H: 40}

type startScreen struct {
	option             selectedOption
	whiteStartNewGame  *sdl.Texture
	yellowStartNewGame *sdl.Texture
	whiteQuitGame      *sdl.Texture
	yellowQuitGame     *sdl.Texture
	scene              *scene
	running            bool
}

func newStartScreen(r *sdl.Renderer) (*startScreen, error) {
	font, err := ttf.OpenFont("res/font/Anton-Regular.ttf", 48)
	if err != nil {
		return nil, fmt.Errorf("could not open font, %v", err)
	}
	defer font.Close()

	wsngSurface, err := font.RenderUTF8Solid("Start new game", sdl.Color{R: 255, G: 255, B: 255, A: 255})
	if err != nil {
		return nil, fmt.Errorf("could not render text, %v", err)
	}
	defer wsngSurface.Free()

	ysngSurface, err := font.RenderUTF8Solid("Start new game", sdl.Color{R: 255, G: 255, B: 0, A: 255})
	if err != nil {
		return nil, fmt.Errorf("could not render text, %v", err)
	}
	defer ysngSurface.Free()

	wqgSurface, err := font.RenderUTF8Solid("Quit game", sdl.Color{R: 255, G: 255, B: 255, A: 255})
	if err != nil {
		return nil, fmt.Errorf("could not render text, %v", err)
	}
	defer wqgSurface.Free()

	yqgSurface, err := font.RenderUTF8Solid("Quit game", sdl.Color{R: 255, G: 255, B: 0, A: 255})
	if err != nil {
		return nil, fmt.Errorf("could not render text, %v", err)
	}
	defer yqgSurface.Free()

	wsngTexture, err := r.CreateTextureFromSurface(wsngSurface)
	if err != nil {
		return nil, fmt.Errorf("could not create texture, %v", err)
	}

	ysngTexture, err := r.CreateTextureFromSurface(ysngSurface)
	if err != nil {
		return nil, fmt.Errorf("could not create texture, %v", err)
	}

	wqgTexture, err := r.CreateTextureFromSurface(wqgSurface)
	if err != nil {
		return nil, fmt.Errorf("could not create texture, %v", err)
	}

	yqgTexture, err := r.CreateTextureFromSurface(yqgSurface)
	if err != nil {
		return nil, fmt.Errorf("could not create texture, %v", err)
	}

	return &startScreen{
		option:             startNewGame,
		whiteStartNewGame:  wsngTexture,
		yellowStartNewGame: ysngTexture,
		whiteQuitGame:      wqgTexture,
		yellowQuitGame:     yqgTexture,
		running:            true}, nil
}

func (ss *startScreen) destroy() {
	ss.whiteStartNewGame.Destroy()
	ss.yellowStartNewGame.Destroy()
	ss.whiteQuitGame.Destroy()
	ss.yellowQuitGame.Destroy()
}

func (ss *startScreen) paint(r *sdl.Renderer) error {
	if ss.option == startNewGame {
		if err := r.Copy(ss.yellowStartNewGame, nil, &startNewGameDst); err != nil {
			return fmt.Errorf("could not copy text, %v", err)
		}

		if err := r.Copy(ss.whiteQuitGame, nil, &quitGameDst); err != nil {
			return fmt.Errorf("could not copy text, %v", err)
		}
	} else {
		if err := r.Copy(ss.whiteStartNewGame, nil, &startNewGameDst); err != nil {
			return fmt.Errorf("could not copy text, %v", err)
		}

		if err := r.Copy(ss.yellowQuitGame, nil, &quitGameDst); err != nil {
			return fmt.Errorf("could not copy text, %v", err)
		}
	}
	return nil
}
