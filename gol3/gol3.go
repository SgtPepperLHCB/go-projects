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
var NEWB string = "+"
var HLTH string = "h"
var OVER string = "o"
var UNDR string = "u"
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
	neighbors int
}

func NewCell(x, y int, state string) Cell {
	return Cell{x, y, state, DEAD, 0}
}
func (cell *Cell) String() string {
	// return fmt.Sprintf("%1s", cell.state)
	return fmt.Sprintf("%1s:%1s:%d", cell.state, cell.nextState, cell.neighbors)
}
func (cell *Cell) isAlive() bool {
	return (cell.state == HLTH || cell.state == NEWB)
}
func (cell *Cell) isDead() bool {
	return cell != nil && cell.state != HLTH && cell.state != NEWB
}
func (cell *Cell) evolve() {
	cell.state = cell.nextState
	cell.nextState = DEAD
}
func (cell *Cell) evaluate(population int) string {
	var next = DEAD
	if cell.isAlive() {
		if population < 2 {
			next = UNDR
		} else if population == 2 || population == 3 {
			next = HLTH
		} else if population > 3 {
			next = OVER
		}

	} else {
		if population == 3 {
			next = NEWB
		}
	}

	return next
}

//
// World
//
type World struct {
	cells [][]Cell
}
type StringWorld [][]string

func NewWorld(size int) World {
	var world World
	world.cells = make([][]Cell, size)
	for x := 0; x < size; x++ {
		world.cells[x] = make([]Cell, size)
	}
	return world
}
func (world *World) create(size int) {
	world.cells = make([][]Cell, size)
	for x := 0; x < size; x++ {
		world.cells[x] = make([]Cell, size)
	}
}
func (world *World) populate(threshold int) {
	var size = len(world.cells)
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			if rand.Intn(100) > threshold {
				world.cells[x][y] = NewCell(x, y, NEWB)
			} else {
				world.cells[x][y] = NewCell(x, y, DEAD)
			}
		}
	}
}
func (world *World) addPattern(x, y int, pattern StringWorld) {
	for dx := range pattern {
		for dy := range pattern[dx] {
			world.cells[x+dx][y+dy] = NewCell(x+dx, y+dx, pattern[dx][dy])
		}
	}
}
func (world *World) neighbors() int {
	var pop int
	for x := range world.cells {
		for y := range world.cells[x] {
			// fmt.Printf("n:xy=%d,%d = %v,%v %d\n", x, y, world.cells[x][y].isAlive(), (x == 1 && y == 1), pop)
			if (x == 1) && (y == 1) {
			} else {
				if world.cells[x][y].isAlive() {
					pop++
				}
			}
		}
	}
	return pop
}
func (world *World) population() int {
	var pop int
	var n = len(world.cells)
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			if world.cells[x][y].isAlive() {
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
			fmt.Printf("%s ", w.cells[x][y].String())
		}
		fmt.Println()
	}
}
func (w *World) showString() string {
	var result string = ""
	var n = len(w.cells)
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			result += fmt.Sprintf("%s ", w.cells[x][y].String())
		}
		result += fmt.Sprintf("\n")
	}
	return result
}
func (world *World) evaluate(debug bool) {
	var n = len(world.cells)

	// Have each cell evaluate it's lot in life
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			var cell = world.cells[x][y]
			// Create a dead world
			var neighborhood = NewWorld(n)

			// Populate neighborhood
			var x0 = cell.x - 1
			var x1 = cell.x + 1
			var y0 = cell.y - 1
			var y1 = cell.y + 1
			// nx,nx neighborhood indices
			var nx = 0
			for x2 := x0; x2 <= x1; x2++ {
				var ny = 0
				for y2 := y0; y2 <= y1; y2++ {
					// wx,wy world indices
					var wx = x2
					var wy = y2
					if wx < 0 {
						wx = n - 1
					} else if wx >= n {
						wx = 0
					}
					if wy < 0 {
						wy = n - 1
					} else if wy >= n {
						wy = 0
					}
					neighborhood.cells[nx][ny] = world.cells[wx][wy]
					ny++
				}
				nx++
			}
			if debug {
				fmt.Printf("---[ neighborhood %d,%d : pop=%d,%d %s -> %s]---\n", x, y, neighborhood.population(), neighborhood.neighbors(), cell.state, cell.evaluate(neighborhood.neighbors()))
				neighborhood.show()
			}

			world.cells[x][y].nextState = cell.evaluate(neighborhood.neighbors())
			world.cells[x][y].neighbors = neighborhood.neighbors()
		}
	}
	if debug {
		fmt.Printf("---[ evaluated ]---\n")
		world.show()
	}
}
func (world *World) evolve() {
	for x := range world.cells {
		for y := range world.cells[x] {
			world.cells[x][y].evolve()
		}
	}
}

//
// Patterns
//
var GLIDER = StringWorld{
	{NEWB, NEWB, NEWB},
	{NEWB, DEAD, DEAD},
	{DEAD, NEWB, DEAD},
}
var TRIOMINO_V = StringWorld{
	{NEWB},
	{NEWB},
	{NEWB},
}

//
//main
//
func main() {
	// Define flags
	dimp := flag.Int("dim", 9, "Dimension for n x n world")
	genp := flag.Int("generations", -1, "Number of generations (-1=forever)")
	pausep := flag.Bool("pause", false, "Pause between steps")
	thresholdp := flag.Int("threshold", 50, "Initial seed (0=fully populated)")
	seedp := flag.Int("seed", -1, "Initial rng seed")
	debugp := flag.Bool("debug", false, "Enable debug output")

	if *seedp > 0 {
		rand.Seed(int64(*seedp))
	} else {
		rand.Seed(time.Now().UTC().UnixNano())
	}

	// Parse
	flag.Parse()

	var world = NewWorld(*dimp)
	world.populate(200)
	// world.addPattern(*dimp-len(GLIDER)*2, *dimp-len(GLIDER)*2, GLIDER)
	world.addPattern(2, 2, TRIOMINO_V)
	// world.populate(*thresholdp)
	scrReset()
	scrPosition(0, 0)
	scrEraseDisplay()
	fmt.Printf("--[ game-of-life: %d x %d, gens=%d, pause=%v, threshold=%d ]--", *dimp, *dimp, *genp, *pausep, *thresholdp)

	var iter = 0
	for true {
		var pop = world.population()
		scrPosition(2, 0)
		scrEraseLine()
		fmt.Printf("[ iter:%d of %d, pop:%d ]", iter, *genp, pop)
		scrPosition(3, 0)
		world.evaluate(*debugp)
		world.show()
		world.evolve()
		iter++
		time.Sleep(200 * time.Millisecond)
		if pop < 1 {
			fmt.Println("--[ dead world ]--")
			break
		}
		if *pausep {
			bufio.NewReader(os.Stdin).ReadBytes('\n')
		}
		if *genp > 0 && iter > *genp {
			break
		}
	}
	fmt.Printf("[ iter:%d of %d, pop:%d ]\n", iter, *genp, world.population())
}
