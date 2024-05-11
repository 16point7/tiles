package main

import (
	"fmt"
	"os"
	"time"

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

	s.SetContent(0, 0, 'T', nil, style)

	s.Show()
	time.Sleep(3 * time.Second)

	s.Fini()
}
