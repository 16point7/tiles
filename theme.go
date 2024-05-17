package main

import "github.com/gdamore/tcell"

var (
	defStyle = tcell.StyleDefault.Background(tcell.NewHexColor(0x8C7E78)).Foreground(tcell.NewHexColor(0xF2F2F2))

	s = tcell.StyleDefault.Bold(true).Foreground(tcell.NewHexColor(0xF2F2F2))

	style0 = s.Background(tcell.NewHexColor(0xE2CEBE)) // 2
	style1 = s.Background(tcell.NewHexColor(0xE2CEBE)) // 4

	style2 = s.Background(tcell.NewHexColor(0xC8A2C8)) // 8
	style3 = s.Background(tcell.NewHexColor(0xC8A2C8)) // 16
	style4 = s.Background(tcell.NewHexColor(0xC8A2C8)) // 32
	style5 = s.Background(tcell.NewHexColor(0xC8A2C8)) // 64

	style6  = s.Background(tcell.NewHexColor(0xFAD7A0)) // 128
	style7  = s.Background(tcell.NewHexColor(0xFAD7A0)) // 256
	style8  = s.Background(tcell.NewHexColor(0xFAD7A0)) // 512
	style9  = s.Background(tcell.NewHexColor(0xFAD7A0)) // 1024
	style10 = s.Background(tcell.NewHexColor(0xFAD7A0)) // 2048

	style11 = s.Background(tcell.NewHexColor(0xE68A8A)) // 4096
	style12 = s.Background(tcell.NewHexColor(0xE68A8A)) // 8192
	style13 = s.Background(tcell.NewHexColor(0xE68A8A)) // 16384

	style14 = s.Background(tcell.NewHexColor(0x6699CC)) // 32768
	style15 = s.Background(tcell.NewHexColor(0x6699CC)) // 65536
	style16 = s.Background(tcell.NewHexColor(0x6699CC)) // 131072
)
