package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type scene struct {
	time        uint64
	background  *sdl.Texture
	startScreen *startScreen
	bird        *bird
	playing     bool
}

func newScene(r *sdl.Renderer) (*scene, error) {
	background, err := img.LoadTexture(r, "res/img/background.png")
	if err != nil {
		return nil, fmt.Errorf("could not load background, %v", err)
	}

	startScreen, err := newStartScreen(r)
	if err != nil {
		return nil, fmt.Errorf("could not create start screen, %v", err)
	}

	bird, err := newBird(r)
	if err != nil {
		return nil, fmt.Errorf("could not create bird, %v: ", err)
	}

	return &scene{
		background:  background,
		startScreen: startScreen,
		bird:        bird,
		playing:     false}, nil
}

func (s *scene) update() {
	if s.playing {
		s.time++
		s.bird.update()
	}
}

func (s *scene) paint(r *sdl.Renderer) error {
	r.Clear()
	if err := r.Copy(s.background, nil, nil); err != nil {
		return fmt.Errorf("could not copy background, %v", err)
	}

	if s.playing {
		if err := s.bird.paint(r, s.time); err != nil {
			return fmt.Errorf("could not paint bird, %v: ", err)
		}
	} else {
		if err := s.startScreen.paint(r); err != nil {
			return fmt.Errorf("could not paint start screen, %v", err)
		}
	}

	r.Present()
	return nil
}

func (s *scene) run(events chan sdl.Event, renderer *sdl.Renderer) <-chan error {
	errc := make(chan error)
	go func() {
		defer close(errc)
		tick := time.Tick(10 * time.Millisecond)
		done := false
		for !done {
			select {
			case event := <-events:
				done = s.handleEvent(event)
				if !s.playing {
					if err := s.paint(renderer); err != nil {
						errc <- err
					}
				}
			case <-tick:
				if s.playing {
					s.update()
					if s.bird.isDead() {
						s.playing = false
					}
					if err := s.paint(renderer); err != nil {
						errc <- err
					}
				}
			}
		}
	}()
	return errc
}

func (s *scene) handleEvent(event sdl.Event) bool {
	switch t := event.(type) {
	case *sdl.QuitEvent:
		return true
	case *sdl.KeyboardEvent:
		if s.playing {
			if t.Keysym.Sym == sdl.K_SPACE && t.State == sdl.PRESSED {
				s.bird.jump()
			}
		} else {
			if t.Keysym.Sym == sdl.K_UP && t.State == sdl.PRESSED {
				if s.startScreen.option == startNewGame {
					s.startScreen.option = quitGame
				} else {
					s.startScreen.option = startNewGame
				}
			}
			if t.Keysym.Sym == sdl.K_DOWN && t.State == sdl.PRESSED {
				if s.startScreen.option == startNewGame {
					s.startScreen.option = quitGame
				} else {
					s.startScreen.option = startNewGame
				}
			}
			if t.Keysym.Sym == sdl.K_RETURN && t.State == sdl.PRESSED {
				if s.startScreen.option == startNewGame {
					s.playing = true
					s.bird.revive()
				} else {
					return true
				}
			}
		}
	default:
	}
	return false
}

func (s *scene) restart() {
}

func (s *scene) destroy() {
	s.background.Destroy()
}

func (s *scene) setStartScreen(startScreen *startScreen) {
	s.startScreen = startScreen
}
