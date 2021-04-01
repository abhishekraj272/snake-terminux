package termsnake

import (
	"fmt"
	"strings"

	"github.com/nsf/termbox-go"
)

const backgroundColor = termbox.ColorBlue
const boardColor = termbox.ColorBlack
const instructionsColor = termbox.ColorWhite

const defaultMarginWidth = 2
const defaultMarginHeight = 1
const titleStartX = defaultMarginWidth
const titleStartY = defaultMarginHeight
const titleHeight = 1
const titleEndY = titleStartY + titleHeight
const boardStartX = defaultMarginWidth
const boardStartY = titleEndY + defaultMarginHeight
const cellWidth = 2

const title = "SNAKE Game By Abhishek Raj (abhishekraj272@gmail.com)"

// gameDetails stores template of the score board
var gameDetails = []string{
	"Round: %v",
	"Score: %v",
	"Length: %v",
	"",
	"GAME OVER!",
}

// Render renders the board & score board on the console using termbox
func (g *Game) Render() {
	termbox.Clear(backgroundColor, backgroundColor)
	termboxPrint(titleStartX, titleStartY, instructionsColor, backgroundColor, title)

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

	instructStartX, instructStartY := g.getBoardEnd()

	for y, instruction := range gameDetails {
		if strings.HasPrefix(instruction, "Round:") {
			instruction = fmt.Sprintf(instruction, g.Round)
		} else if strings.HasPrefix(instruction, "Score:") {
			instruction = fmt.Sprintf(instruction, g.score)
		} else if strings.HasPrefix(instruction, "Length:") {
			instruction = fmt.Sprintf(instruction, g.score)
		} else if strings.HasPrefix(instruction, "GAME OVER") && g.State != 0 {
			instruction = ""
		}

		termboxPrint(instructStartX, instructStartY+y, instructionsColor, backgroundColor, instruction)
	}
	termbox.Flush()
}

// termboxPrint is used to print score board
func termboxPrint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}
