package main

import (
	"github.com/16point7/tiles/game"
	"github.com/gdamore/tcell"
)

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

	s.SetContent(x, y, lb, nil, bgStyle)
	x++

	for t := 0; t < game.Side-1; t++ {
		for k := 0; k < cellWidth; k++ {
			s.SetContent(x, y, fill, nil, bgStyle)
			x++
		}
		s.SetContent(x, y, div, nil, bgStyle)
		x++
	}
	for k := 0; k < cellWidth; k++ {
		s.SetContent(x, y, fill, nil, bgStyle)
		x++
	}

	s.SetContent(x, y, rb, nil, bgStyle)
}

func drawData(s tcell.Screen, y int, row []uint) {
	x := 0

	s.SetContent(x, y, tcell.RuneVLine, nil, bgStyle)
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

		var tileStyle tcell.Style
		switch row[t] {
		case 0:
			tileStyle = bgStyle
		case 2 << 0:
			tileStyle = style0
		case 2 << 1:
			tileStyle = style1
		case 2 << 2:
			tileStyle = style2
		case 2 << 3:
			tileStyle = style3
		case 2 << 4:
			tileStyle = style4
		case 2 << 5:
			tileStyle = style5
		case 2 << 6:
			tileStyle = style6
		case 2 << 7:
			tileStyle = style7
		case 2 << 8:
			tileStyle = style8
		case 2 << 9:
			tileStyle = style9
		case 2 << 10:
			tileStyle = style10
		case 2 << 11:
			tileStyle = style11
		case 2 << 12:
			tileStyle = style12
		case 2 << 13:
			tileStyle = style13
		case 2 << 14:
			tileStyle = style14
		case 2 << 15:
			tileStyle = style15
		case 2 << 16:
			tileStyle = style16
		}

		i := 0

		for i < left {
			s.SetContent(x, y, ' ', nil, tileStyle)
			x++
			i++
		}

		for len > 0 {
			len--
			s.SetContent(x, y, str[len], nil, tileStyle)
			x++
			i++
		}

		for i < cellWidth {
			s.SetContent(x, y, ' ', nil, tileStyle)
			x++
			i++
		}

		s.SetContent(x, y, tcell.RuneVLine, nil, bgStyle)
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
		s.SetContent(x, y, str[len], nil, bgStyle)
		x++
	}
}

func drawStatus(s tcell.Screen, y int, gameOver bool) {
	if gameOver {
		s.SetContent(0, y, 'G', nil, bgStyle)
		s.SetContent(1, y, 'a', nil, bgStyle)
		s.SetContent(2, y, 'm', nil, bgStyle)
		s.SetContent(3, y, 'e', nil, bgStyle)
		s.SetContent(4, y, ' ', nil, bgStyle)
		s.SetContent(5, y, 'O', nil, bgStyle)
		s.SetContent(6, y, 'v', nil, bgStyle)
		s.SetContent(7, y, 'e', nil, bgStyle)
		s.SetContent(8, y, 'r', nil, bgStyle)
		s.SetContent(9, y, '!', nil, bgStyle)
	}
}
