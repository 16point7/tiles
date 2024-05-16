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
		drawState(s, style, g.Board, g.Score)
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
			case tcell.KeyRune:
				if ev.Rune() == 'r' {
					g.Reset()
					s.Clear()
				}
			}
		}

		if gameOver {
			return
		}
	}
}

func drawState(s tcell.Screen, style tcell.Style, board [][]uint, score uint) {
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
	y++

	drawScore(s, style, y, score)
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

var (
	base    = tcell.StyleDefault.Bold(true)
	style0  = base.Background(tcell.Color246)
	style1  = base.Background(tcell.Color240)
	style2  = base.Background(tcell.Color39)
	style3  = base.Background(tcell.Color33)
	style4  = base.Background(tcell.Color185)
	style5  = base.Background(tcell.Color179)
	style6  = base.Background(tcell.Color202)
	style7  = base.Background(tcell.Color196)
	style8  = base.Background(tcell.Color124)
	style9  = base.Background(tcell.Color88)
	style10 = base.Background(tcell.Color52)
	style11 = base.Background(tcell.Color52)
	style12 = base.Background(tcell.Color52)
	style13 = base.Background(tcell.Color52)
	style14 = base.Background(tcell.Color52)
	style15 = base.Background(tcell.Color52)
	style16 = base.Background(tcell.Color52)
)

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

		var numStyle tcell.Style
		switch row[t] {
		case 0:
			numStyle = style
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

		s.SetContent(x, y, tcell.RuneVLine, nil, style)
		x++
	}
}

func drawScore(s tcell.Screen, style tcell.Style, y int, score uint) {
	s.SetContent(0, y, 'S', nil, style)
	s.SetContent(1, y, 'c', nil, style)
	s.SetContent(2, y, 'o', nil, style)
	s.SetContent(3, y, 'r', nil, style)
	s.SetContent(4, y, 'e', nil, style)
	s.SetContent(5, y, ':', nil, style)
	s.SetContent(6, y, ' ', nil, style)

	x := 7

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

	for len > 0 {
		len--
		s.SetContent(x, y, str[len], nil, style)
		x++
	}
}
