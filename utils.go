package main

import (
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func (g *Game) min() int {
	if g.boardHeight < g.boardWidth {
		return g.boardHeight - 1
	}
	return g.boardWidth - 1
}

func (g *Game) getRand() (int, int) {
	rand.Seed(time.Now().UnixNano())
	rand1 := rand.Intn(g.min() + 1)

	rand.Seed(time.Now().UnixNano())
	rand2 := rand.Intn(g.min() + 1)

	return rand2, rand1
}

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
