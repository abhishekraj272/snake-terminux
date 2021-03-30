package main

import (
	"math/rand"
	"time"
)

func (g *Game) min() int {
	if g.boardHeight < g.boardWidth {
		return g.boardHeight
	}
	return g.boardWidth
}

func (g *Game) getRand() (int, int) {
	rand.Seed(time.Now().UnixNano())
	rand1 := rand.Intn(g.min()-0+1) + g.min()

	rand.Seed(time.Now().UnixNano())
	rand2 := rand.Intn(g.min()-0+1) + g.min()

	return rand2, rand1
}

func (g *Game) getTail() (int, int) {

}
