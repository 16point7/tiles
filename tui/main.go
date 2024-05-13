package main

import (
	"fmt"
	"os"
	"time"

	"github.com/16point7/tiles"
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

	style := tcell.StyleDefault
	s.SetStyle(style)

	drawBoard(s, style)

	s.Show()
	time.Sleep(10 * time.Second)

	s.Fini()
}

func drawBoard(s tcell.Screen, style tcell.Style) {
	y := 0

	drawLine(s, style, y, tcell.RuneULCorner, tcell.RuneHLine, tcell.RuneTTee, tcell.RuneURCorner)
	y++

	for j := 0; j < tiles.Length-1; j++ {
		drawData(s, style, y)
		y++
		drawLine(s, style, y, tcell.RuneLTee, tcell.RuneHLine, tcell.RunePlus, tcell.RuneRTee)
		y++
	}
	drawData(s, style, y)
	y++

	drawLine(s, style, y, tcell.RuneLLCorner, tcell.RuneHLine, tcell.RuneBTee, tcell.RuneLRCorner)
}

const cellWidth = 9

func drawLine(s tcell.Screen, style tcell.Style, y int, lb, fill, div, rb rune) {
	x := 0

	s.SetContent(x, y, lb, nil, style)
	x++

	for c := 0; c < tiles.Length-1; c++ {
		for k := 0; k < cellWidth; k++ {
			s.SetContent(x, y, fill, nil, style)
			x++
		}
		s.SetContent(x, y, div, nil, style)
		x++
	}
	for k := 0; k < cellWidth; k++ {
		s.SetContent(x, y, fill, nil, style)
		x++
	}

	s.SetContent(x, y, rb, nil, style)
}

func drawData(s tcell.Screen, style tcell.Style, y int) {
	x := 0

	s.SetContent(x, y, tcell.RuneVLine, nil, style)
	x++

	for c := 0; c < tiles.Length; c++ {
		for i := 0; i < cellWidth/2; i++ {
			s.SetContent(x, y, ' ', nil, style)
			x++
		}
		s.SetContent(x, y, 'D', nil, style)
		x++
		for i := 0; i < cellWidth/2; i++ {
			s.SetContent(x, y, ' ', nil, style)
			x++
		}
		s.SetContent(x, y, tcell.RuneVLine, nil, style)
		x++
	}
}
