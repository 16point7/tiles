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

	s.SetStyle(defStyle)

	g := game.New()
	g.Board = [][]uint{
		{0, 2, 4, 8},
		{16, 32, 64, 128},
		{256, 512, 1024, 2048},
		{4096, 8192, 16384, 32768},
	}

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

func drawState(s tcell.Screen, board [][]uint, score uint, gameOver bool) {
	y := 0

	drawLine(s, y, tcell.RuneULCorner, tcell.RuneHLine, tcell.RuneTTee, tcell.RuneURCorner)
	y++

	for j := 0; j < game.Side-1; j++ {
		drawData(s, y, board[j])
		y++
		drawLine(s, y, tcell.RuneLTee, tcell.RuneHLine, tcell.RunePlus, tcell.RuneRTee)
		y++
	}
	drawData(s, y, board[game.Side-1])
	y++

	drawLine(s, y, tcell.RuneLLCorner, tcell.RuneHLine, tcell.RuneBTee, tcell.RuneLRCorner)
	y++

	drawScore(s, y, score)
	y++

	drawStatus(s, y, gameOver)
}

const cellWidth = 9

func drawLine(s tcell.Screen, y int, lb, fill, div, rb rune) {
	x := 0

	s.SetContent(x, y, lb, nil, defStyle)
	x++

	for t := 0; t < game.Side-1; t++ {
		for k := 0; k < cellWidth; k++ {
			s.SetContent(x, y, fill, nil, defStyle)
			x++
		}
		s.SetContent(x, y, div, nil, defStyle)
		x++
	}
	for k := 0; k < cellWidth; k++ {
		s.SetContent(x, y, fill, nil, defStyle)
		x++
	}

	s.SetContent(x, y, rb, nil, defStyle)
}

func drawData(s tcell.Screen, y int, row []uint) {
	x := 0

	s.SetContent(x, y, tcell.RuneVLine, nil, defStyle)
	x++

	for t := 0; t < game.Side; t++ {
		str := [cellWidth]rune{}
		len := 0

		num := row[t]
		for num != 0 {
			str[len] = '0' + rune((num % 10))
			num /= 10
			len++
		}

		var left int
		if len%2 == 0 {
			left = cellWidth/2 - (len/2 - 1)
		} else {
			left = cellWidth/2 - (len-1)/2
		}

		var numStyle tcell.Style
		switch row[t] {
		case 0:
			numStyle = defStyle
		case 2 << 0:
			numStyle = style0
		case 2 << 1:
			numStyle = style1
		case 2 << 2:
			numStyle = style2
		case 2 << 3:
			numStyle = style3
		case 2 << 4:
			numStyle = style4
		case 2 << 5:
			numStyle = style5
		case 2 << 6:
			numStyle = style6
		case 2 << 7:
			numStyle = style7
		case 2 << 8:
			numStyle = style8
		case 2 << 9:
			numStyle = style9
		case 2 << 10:
			numStyle = style10
		case 2 << 11:
			numStyle = style11
		case 2 << 12:
			numStyle = style12
		case 2 << 13:
			numStyle = style13
		case 2 << 14:
			numStyle = style14
		case 2 << 15:
			numStyle = style15
		case 2 << 16:
			numStyle = style16
		}

		i := 0

		for i < left {
			s.SetContent(x, y, ' ', nil, numStyle)
			x++
			i++
		}

		for len > 0 {
			len--
			s.SetContent(x, y, str[len], nil, numStyle)
			x++
			i++
		}

		for i < cellWidth {
			s.SetContent(x, y, ' ', nil, numStyle)
			x++
			i++
		}

		s.SetContent(x, y, tcell.RuneVLine, nil, defStyle)
		x++
	}
}

func drawScore(s tcell.Screen, y int, score uint) {
	str := [cellWidth]rune{}
	len := 0
	for {
		str[len] = '0' + rune((score % 10))
		score /= 10
		len++
		if score == 0 {
			break
		}
	}

	for x := 0; len > 0; {
		len--
		s.SetContent(x, y, str[len], nil, defStyle)
		x++
	}
}

func drawStatus(s tcell.Screen, y int, gameOver bool) {
	if gameOver {
		s.SetContent(0, y, 'G', nil, defStyle)
		s.SetContent(1, y, 'a', nil, defStyle)
		s.SetContent(2, y, 'm', nil, defStyle)
		s.SetContent(3, y, 'e', nil, defStyle)
		s.SetContent(4, y, ' ', nil, defStyle)
		s.SetContent(5, y, 'O', nil, defStyle)
		s.SetContent(6, y, 'v', nil, defStyle)
		s.SetContent(7, y, 'e', nil, defStyle)
		s.SetContent(8, y, 'r', nil, defStyle)
		s.SetContent(9, y, '!', nil, defStyle)
	}
}
