package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var RESET string = "\u001b[0m"
var HOME0 string = "\u001b[1;0H"
var HOME1 string = "\u001b[2;0H"
var NEWBORN string = "n"
var HEALTHY string = "h"
var OVERPOP string = "o"
var UNDERPOP string = "u"
var DEAD string = "-"

//
// Screen
//
var SCR_RESET string = "\u001b[0m"
var SCR_ERASE_DISPLAY string = "\u001b[2J"
var SCR_ERASE_LINE string = "\u001b[K"
var SCR_CURSOR_POSITION = "\u001b[%d;%dH"

func scrEraseDisplay() {
	fmt.Printf(SCR_ERASE_DISPLAY)
}
func scrEraseLine() {
	fmt.Printf(SCR_ERASE_LINE)
}
func scrPosition(line, column int) {
	fmt.Printf(SCR_CURSOR_POSITION, line, column)
}
func scrReset() {
	fmt.Printf(SCR_RESET)
}

func minmax(v, x, y int) int {
	if v < x {
		return x
	} else if v > y {
		return y
	} else {
		return v
	}
}

//
// Cell
//
type Cell struct {
	x         int
	y         int
	state     string
	nextState string
}

func NewCell(x, y int, state string) Cell {
	return Cell{x, y, state, DEAD}
}

func (cell *Cell) isAlive() bool {
	return (cell.state == HEALTHY || cell.state == NEWBORN)
}
func (cell *Cell) isDead() bool {
	return cell != nil && cell.state != HEALTHY && cell.state != NEWBORN
}
func (cell *Cell) evolve() {
	cell.state = cell.nextState
	cell.nextState = DEAD
}
func (cell *Cell) prepareToEvolve(neighborhood World) {
	var pop = neighborhood.population()
	if !cell.isAlive() && pop == 3 {
		cell.nextState = NEWBORN
	} else if pop < 2 {
		cell.nextState = DEAD
	} else if pop >= 2 && pop <= 3 {
		cell.nextState = HEALTHY
	} else if pop > 3 {
		cell.nextState = DEAD
	}
}

//
// World
//
type World struct {
	cells [][]Cell
}

func (world *World) create(n, threshold int) {
	world.cells = make([][]Cell, n)
	for x := 0; x < n; x++ {
		world.cells[x] = make([]Cell, n)
		for y := 0; y < n; y++ {
			if rand.Intn(100) > threshold {
				world.cells[x][y] = NewCell(x, y, NEWBORN)
			} else {
				world.cells[x][y] = NewCell(x, y, DEAD)
			}
		}
	}
}
func (w *World) population() int {
	var pop int
	var n = len(w.cells)
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			if w.cells[x][y].isAlive() {
				pop++
			}
		}
	}

	return pop
}
func (w *World) show() {
	var n = len(w.cells)
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			fmt.Printf("%1s ", w.cells[x][y].state)
		}
		fmt.Println()
	}
}
func (w *World) evolve() {
	var n = len(w.cells)

	// Have each cell evaluate it's lot in life
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			var cell = w.cells[x][y]
			//Create a dead world
			var neighborhood = World{nil}
			neighborhood.create(3, 101)
			var x0 = cell.x - 1
			var x1 = cell.x + 1
			var y0 = cell.y - 1
			var y1 = cell.y + 1
			var xxx = 0
			for xx := x0; xx <= x1; xx++ {
				var yyy = 0
				for yy := y0; yy <= y1; yy++ {
					if xx >= 0 && xx < n {
						if yy >= 0 && yy < n {
							neighborhood.cells[xxx][yyy] = w.cells[xx][yy]
						}
					}
					yyy++
				}
				xxx++
			}
			w.cells[x][y].prepareToEvolve(neighborhood)
		}
	}
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			w.cells[x][y].evolve()
		}
	}
}

//
//main
//
func main() {
	// Define flags
	dimp := flag.Int("dim", 9, "Dimension for n x n world")
	genp := flag.Int("generations", 9, "Number of generations")
	pausep := flag.Bool("pause", false, "Pause between steps")
	thresholdp := flag.Int("threshold", 50, "Initial populating random threshold (0=fully populated, >100=none)")
	seedp := flag.Int("seed", -1, "Initial rng seed")

	if *seedp > 0 {
		rand.Seed(int64(*seedp))
	} else {
		rand.Seed(time.Now().UTC().UnixNano())
	}

	// Parse
	flag.Parse()

	var world = World{nil}
	world.create(*dimp, *thresholdp)
	scrReset()
	scrPosition(0, 0)
	scrEraseDisplay()
	fmt.Printf("--[ game-of-life: %d x %d, gens=%d, pause=%v, seed=%d ]--", *dimp, *dimp, *genp, *pausep, *seedp)

	var iter = 0
	for true {
		var pop = world.population()
		scrPosition(2, 0)
		scrEraseLine()
		fmt.Printf("[ iter:%d of %d, pop:%d ]", iter, *genp, pop)
		scrPosition(3, 0)
		world.show()
		time.Sleep(200 * time.Millisecond)
		if pop < 1 {
			fmt.Println("--[ dead world ]--")
			break
		}
		if *pausep {
			bufio.NewReader(os.Stdin).ReadBytes('\n')
		}
		world.evolve()
		iter++
	}
	fmt.Printf("[ iter:%d of %d, pop:%d ]\n", iter, *genp, world.population())

	/*
		for iter := 0; iter < *genp; iter++ {
		 	var pop = world.population()
		 	scrPosition(2, 0)
		 	scrEraseLine()
		 	fmt.Printf("[ iter:%d of %d, pop:%d ]", iter, *genp, pop)
		 	scrPosition(3, 0)
		 	world.show()
		 	time.Sleep(200 * time.Millisecond)
		 	if pop < 1 {
		 		fmt.Println("--[ dead world ]--")
		 		break
		 	}
		 	if *pausep {
		 		bufio.NewReader(os.Stdin).ReadBytes('\n')
		 	}
		 	world.evolve()
		 }
		 fmt.Printf("[ iter:%d of %d, pop:%d ]\n", *genp, *genp, world.population())
	*/
}
