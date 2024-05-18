package main

import "github.com/gdamore/tcell"

var (
	bgStyle = tcell.StyleDefault.Background(tcell.NewHexColor(0xbbada0)).Foreground(tcell.NewHexColor(0x000000))

	bold  = tcell.StyleDefault.Bold(true)
	dark  = bold.Foreground(tcell.NewHexColor(0x776e65))
	light = bold.Foreground(tcell.NewHexColor(0xf9f6f2))

	style0  = dark.Background(tcell.NewHexColor(0xeee4da))
	style1  = dark.Background(tcell.NewHexColor(0xede0c8))
	style2  = light.Background(tcell.NewHexColor(0xf2b179))
	style3  = light.Background(tcell.NewHexColor(0xf59563))
	style4  = light.Background(tcell.NewHexColor(0xf67c5f))
	style5  = light.Background(tcell.NewHexColor(0xf65e3b))
	style6  = light.Background(tcell.NewHexColor(0xedcf72))
	style7  = light.Background(tcell.NewHexColor(0xedcc61))
	style8  = light.Background(tcell.NewHexColor(0xedc850))
	style9  = light.Background(tcell.NewHexColor(0xedc53f))
	style10 = light.Background(tcell.NewHexColor(0xedc22e))
	style11 = light.Background(tcell.NewHexColor(0xef9096))
	style12 = light.Background(tcell.NewHexColor(0xe75a63))
	style13 = light.Background(tcell.NewHexColor(0xe1443a))
	style14 = light.Background(tcell.NewHexColor(0xa6cae2))
	style15 = light.Background(tcell.NewHexColor(0x6aaad6))
	style16 = light.Background(tcell.NewHexColor(0x077bb9))
)
