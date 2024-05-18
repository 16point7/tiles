package main

import (
	"fmt"
	"os"

	"github.com/16point7/tiles/game"
	"github.com/gdamore/tcell"
)

func main() {
	s, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if err = s.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	defer s.Fini()

	s.SetStyle(bgStyle)

	g := game.New()

	for {
		drawState(s, g.Board, g.Score, g.GameOver)
		s.Show()

		switch ev := s.PollEvent().(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyLeft:
				g.MoveLeft()
			case tcell.KeyRight:
				g.MoveRight()
			case tcell.KeyUp:
				g.MoveUp()
			case tcell.KeyDown:
				g.MoveDown()
			case tcell.KeyEscape:
				return
			case tcell.KeyRune:
				if ev.Rune() == 'r' {
					g.Reset()
					s.Clear()
				}
			}
		}
	}
}
