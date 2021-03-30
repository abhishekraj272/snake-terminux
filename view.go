package main

import (
	"github.com/nsf/termbox-go"
)

const backgroundColor = termbox.ColorBlue
const boardColor = termbox.ColorBlack
const instructionsColor = termbox.ColorYellow

const defaultMarginWidth = 2
const defaultMarginHeight = 1
const titleStartX = defaultMarginWidth
const titleStartY = defaultMarginHeight
const titleHeight = 1
const titleEndY = titleStartY + titleHeight
const boardStartX = defaultMarginWidth
const boardStartY = titleEndY + defaultMarginHeight
const cellWidth = 2

const title = "Nostalgic SNAKE Game"

func render(g *Game) {
	termbox.Clear(backgroundColor, backgroundColor)
	tbprint(titleStartX, titleStartY, instructionsColor, backgroundColor, title)
	for i := 0; i < g.boardHeight; i++ {
		for j := 0; j < g.boardWidth; j++ {
			var cellColor termbox.Attribute
			switch g.board[i][j] {
			case 0:
				cellColor = termbox.ColorBlack
			case 1:
				cellColor = termbox.ColorCyan
			case 2:
				cellColor = termbox.ColorMagenta
			}
			for x := 0; x < cellWidth; x++ {
				termbox.SetCell(boardStartX+cellWidth*j+x, boardStartY+i, ' ', cellColor, cellColor)
			}
		}
	}
	termbox.Flush()
}

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}
