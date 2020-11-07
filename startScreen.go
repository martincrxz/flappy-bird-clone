package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/ttf"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type selectedOption string

const (
	startNewGame selectedOption = "startNewGame"
	quitGame                    = "quitGame"
)

var startNewGameDst sdl.Rect = sdl.Rect{X: 300 - 80, Y: 300, W: 160, H: 48}
var quitGameDst sdl.Rect = sdl.Rect{X: 300 - 50, Y: 400, W: 100, H: 48}

type startScreen struct {
	background         *sdl.Texture
	option             selectedOption
	whiteStartNewGame  *sdl.Texture
	yellowStartNewGame *sdl.Texture
	whiteQuitGame      *sdl.Texture
	yellowQuitGame     *sdl.Texture
}

func newStartScreen(r *sdl.Renderer) (*startScreen, error) {
	background, err := img.LoadTexture(r, "res/img/background.png")
	if err != nil {
		return nil, fmt.Errorf("could not load background, %v", err)
	}

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
		background:         background,
		option:             startNewGame,
		whiteStartNewGame:  wsngTexture,
		yellowStartNewGame: ysngTexture,
		whiteQuitGame:      wqgTexture,
		yellowQuitGame:     yqgTexture}, nil
}

func (ss *startScreen) destroy() {
	ss.background.Destroy()
	ss.whiteStartNewGame.Destroy()
	ss.yellowStartNewGame.Destroy()
	ss.whiteQuitGame.Destroy()
	ss.yellowQuitGame.Destroy()
}

func (ss *startScreen) handleEvent(event sdl.Event) bool {
	switch t := event.(type) {
	case *sdl.QuitEvent:
		return true
	case *sdl.KeyboardEvent:
		if t.Keysym.Sym == sdl.K_UP && t.State == sdl.PRESSED {
			if ss.option == startNewGame {
				ss.option = quitGame
			} else {
				ss.option = startNewGame
			}
		}
		if t.Keysym.Sym == sdl.K_DOWN && t.State == sdl.PRESSED {
			if ss.option == startNewGame {
				ss.option = quitGame
			} else {
				ss.option = startNewGame
			}
		}
	default:
	}
	return false
}

func (ss *startScreen) run(events chan sdl.Event, renderer *sdl.Renderer) <-chan error {
	errc := make(chan error)
	go func() {
		defer close(errc)
		done := false
		for !done {
			select {
			case event := <-events:
				done = ss.handleEvent(event)
				ss.paint(renderer)
			}
		}
	}()
	return errc
}

func (ss *startScreen) paint(r *sdl.Renderer) error {
	r.Clear()
	if err := r.Copy(ss.background, nil, nil); err != nil {
		return fmt.Errorf("could not copy background, %v", err)
	}

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

	r.Present()
	return nil
}
