package main

import (
	"time"

	"github.com/nsf/termbox-go"
)

func main() {
	g := NewGame()

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

	render(g)

	for {
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey {
				switch {
				case ev.Key == termbox.KeyArrowLeft:
					g.direction = "left"

					if g.checkMove() {
						g.moveLeft()
						g.checkFood()
					}
				case ev.Key == termbox.KeyArrowRight:
					g.direction = "right"

					if g.checkMove() {
						g.moveRight()
						g.checkFood()
					}
				case ev.Key == termbox.KeyArrowUp:
					g.direction = "up"

					if g.checkMove() {
						g.moveUp()
						g.checkFood()
					}
				case ev.Key == termbox.KeyArrowDown:
					g.direction = "down"

					if g.checkMove() {
						g.moveDown()
						g.checkFood()
					}

				case ev.Ch == 'q' || ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC || ev.Key == termbox.KeyCtrlD:
					return
				}
			}
		default:
			render(g)
			time.Sleep(10 * time.Millisecond)
		}
	}
}
