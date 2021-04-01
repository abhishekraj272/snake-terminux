package main

import (
	"github.com/abhishekraj272/snake-terminux/termsnake"
	"github.com/nsf/termbox-go"
)

func main() {
	g := termsnake.NewGame()

	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	defer termbox.Close()

	eventQueue := make(chan termbox.Event)

	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	g.Render()

	for {
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey {
				switch {
				case ev.Key == termbox.KeyArrowLeft:
					g.Direction = "left"

					if g.CheckMove() && g.State == 1 {
						g.Round++
						g.MoveLeft()
						g.CheckFood()
					} else {
						g.State = 0
					}
				case ev.Key == termbox.KeyArrowRight:
					g.Direction = "right"

					if g.CheckMove() && g.State == 1 {
						g.Round++
						g.MoveRight()
						g.CheckFood()
					} else {
						g.State = 0
					}
				case ev.Key == termbox.KeyArrowUp:
					g.Direction = "up"

					if g.CheckMove() && g.State == 1 {
						g.Round++
						g.MoveUp()
						g.CheckFood()
					} else {
						g.State = 0
					}
				case ev.Key == termbox.KeyArrowDown:
					g.Direction = "down"

					if g.CheckMove() && g.State == 1 {
						g.Round++
						g.MoveDown()
						g.CheckFood()
					} else {
						g.State = 0
					}

				case ev.Ch == 'q' || ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC || ev.Key == termbox.KeyCtrlD:
					return
				}
			}
		default:
			g.Render()
		}
	}
}
