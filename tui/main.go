package main

import (
	"fmt"
	"os"

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
	defer s.Fini()

	style := tcell.StyleDefault
	s.SetStyle(style)

	g := tiles.NewGame()

	for {
		drawBoard(s, style, g.Board)
		s.Show()

		var gameOver bool

		switch ev := s.PollEvent().(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyLeft:
				gameOver = g.MoveLeft()
			case tcell.KeyRight:
				gameOver = g.MoveRight()
			case tcell.KeyUp:
				gameOver = g.MoveUp()
			case tcell.KeyDown:
				gameOver = g.MoveDown()
			case tcell.KeyEscape:
				gameOver = true
			}
		}

		if gameOver {
			return
		}
	}
}

func drawBoard(s tcell.Screen, style tcell.Style, board [][]uint) {
	y := 0

	drawLine(s, style, y, tcell.RuneULCorner, tcell.RuneHLine, tcell.RuneTTee, tcell.RuneURCorner)
	y++

	for j := 0; j < tiles.Side-1; j++ {
		drawData(s, style, y, board[j])
		y++
		drawLine(s, style, y, tcell.RuneLTee, tcell.RuneHLine, tcell.RunePlus, tcell.RuneRTee)
		y++
	}
	drawData(s, style, y, board[tiles.Side-1])
	y++

	drawLine(s, style, y, tcell.RuneLLCorner, tcell.RuneHLine, tcell.RuneBTee, tcell.RuneLRCorner)
}

const cellWidth = 9

func drawLine(s tcell.Screen, style tcell.Style, y int, lb, fill, div, rb rune) {
	x := 0

	s.SetContent(x, y, lb, nil, style)
	x++

	for t := 0; t < tiles.Side-1; t++ {
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

func drawData(s tcell.Screen, style tcell.Style, y int, row []uint) {
	x := 0

	s.SetContent(x, y, tcell.RuneVLine, nil, style)
	x++

	for t := 0; t < tiles.Side; t++ {
		str := [cellWidth]rune{}
		len := 0

		num := row[t]
		for {
			str[len] = '0' + rune((num % 10))
			num /= 10
			len++
			if num == 0 {
				break
			}
		}

		var left int
		if len%2 == 0 {
			left = cellWidth/2 - (len/2 - 1)
		} else {
			left = cellWidth/2 - (len-1)/2
		}

		i := 0

		for i < left {
			s.SetContent(x, y, ' ', nil, style)
			x++
			i++
		}

		for len > 0 {
			len--
			s.SetContent(x, y, str[len], nil, style)
			x++
			i++
		}

		for i < cellWidth {
			s.SetContent(x, y, ' ', nil, style)
			x++
			i++
		}

		s.SetContent(x, y, tcell.RuneVLine, nil, style)
		x++
	}
}
