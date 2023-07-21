package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

func drawText(s *tcell.Screen, x1, y1, x2, y2 int, style *tcell.Style, text string) {
	row := y1
	col := x1
	for _, r := range text {
		(*s).SetContent(col, row, r, nil, *style)
		col++
		if col > x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}

func drawBox(s *tcell.Screen, x1, y1, x2, y2 int, style *tcell.Style, text string) {
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	if x2 < x1 {
		x1, x2 = x2, x1
	}

	for row := y1; row <= y2; row++ {
		for col := x1; col <= x2; col++ {
			(*s).SetContent(col, row, ' ', nil, *style)
		}
	}

	for col := x1; col <= x2; col++ {
		(*s).SetContent(col, y1, tcell.RuneHLine, nil, *style)
		(*s).SetContent(col, y2, tcell.RuneHLine, nil, *style)
	}
	for row := y1 + 1; row < y2; row++ {
		(*s).SetContent(x1, row, tcell.RuneVLine, nil, *style)
		(*s).SetContent(x2, row, tcell.RuneVLine, nil, *style)
	}

	if y1 != y2 && x1 != x2 {
		(*s).SetContent(x1, y1, tcell.RuneULCorner, nil, *style)
		(*s).SetContent(x2, y1, tcell.RuneURCorner, nil, *style)
		(*s).SetContent(x1, y2, tcell.RuneLLCorner, nil, *style)
		(*s).SetContent(x2, y2, tcell.RuneLRCorner, nil, *style)
	}

	drawText(s, x1+1, y1+1, x2-1, y2-1, style, text)
}

func main() {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}

	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	style := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	boxStyle := tcell.StyleDefault.Background(tcell.ColorBlue).Foreground(tcell.ColorWhite)

	s.SetStyle(style)

	// Draw initial boxes
	drawBox(&s, 1, 1, 42, 7, &boxStyle, "Click and drag to draw a box")
	drawBox(&s, 5, 9, 32, 14, &boxStyle, "Press C to reset")

	defer func() {
		s.Fini()
		os.Exit(0)
	}()

	for {
		s.Show()

		switch ev := s.PollEvent().(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape, tcell.KeyCtrlC:
				return
			}
		}
	}
}
