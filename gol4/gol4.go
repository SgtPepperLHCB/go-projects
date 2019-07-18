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

type Health int

const (
	DeadJim Health = iota + 1
	Healthy
	Newborn
	Overpop
	Undrpop
)

func (level Health) String() string {
	names := [...]string{
		"_",
		" ",
		"h",
		"+",
		" ",
		" ",
	}
	return names[level]
}
func (level Health) IsAlive() bool {
	return level == Healthy || level == Newborn
}
func (level Health) IsDead() bool {
	return level != Healthy && level != Newborn
}

var TRIONIMO_V = StringWorld{
	{Newborn},
	{Newborn},
	{Newborn},
}
var GLIDER = StringWorld{
	{Newborn, Newborn, Newborn},
	{Newborn, DeadJim, DeadJim},
	{DeadJim, Newborn, DeadJim},
}

//
// Cell
//
type Cell struct {
	health Health
	next   Health
}

func NewCell(level Health) Cell {
	return Cell{level, DeadJim}
}
func (cell *Cell) String(verbose ...bool) string {
	if len(verbose) > 0 {
		if verbose[0] {
			return fmt.Sprintf("%s:%s", cell.health.String(), cell.next.String())
		}
	}
	return fmt.Sprintf("%1s", cell.health.String())
}
func (cell *Cell) Evolve() {
	cell.health = cell.next
}
func (cell *Cell) IsAlive() bool {
	return cell.health.IsAlive()
}
func (cell *Cell) IsDead() bool {
	return cell.health.IsDead()
}

//
// World
//
type World struct {
	cells [][]Cell
}
type StringWorld [][]Health

func NewWorld(dimension int) World {
	var world World
	world.cells = make([][]Cell, dimension)
	for row := range world.cells {
		world.cells[row] = make([]Cell, dimension)
	}
	world.Populate()
	return world
}
func (world *World) Size() int {
	return len(world.cells)
}
func (world *World) NewCell(row, col int, level Health) {
	world.cells[row][col] = NewCell(level)
}
func (world *World) IsAliveCell(row, col int) bool {
	return world.cells[row][col].IsAlive()
}
func (world *World) Populate() {
	for row := range world.cells {
		for col := range world.cells[row] {
			world.cells[row][col] = NewCell(DeadJim)
		}
	}
}
func (world *World) PopulateWithThreshold(opt_threshold int, opt_max ...int) {
	var threshold = opt_threshold
	var m int = len(world.cells) * len(world.cells)
	if len(opt_max) > 0 {
		m = opt_max[0]
	}
	for row := range world.cells {
		for col := range world.cells[row] {
			if m < 1 {
				break
			}
			world.NewCell(row, col, DeadJim)
			if rand.Intn(100) > threshold {
				m--
				world.NewCell(row, col, Newborn)
			}
		}
	}
}
func (world *World) AddPattern(row, col int, pattern StringWorld) {
	for dr := range pattern {
		for dc := range pattern[dr] {
			world.cells[row+dr][col+dc] = NewCell(pattern[dr][dc])
		}
	}
}
func (world *World) Population() int {
	var pop int
	for row := range world.cells {
		for col := range world.cells[row] {
			if world.cells[row][col].IsAlive() {
				pop++
			}
		}
	}
	return pop
}
func (world *World) Neighbors(row, col int) int {
	var neighbors int
	for r := row - 1; r <= row+1; r++ {
		for c := col - 1; c <= col+1; c++ {
			var r0 = r
			var c0 = c
			if r0 < 0 {
				r0 = world.Size() - 1
			} else if r0 >= world.Size() {
				r0 = 0
			}
			if c0 < 0 {
				c0 = world.Size() - 1
			} else if c0 >= world.Size() {
				c0 = 0
			}
			if r0 != row || c0 != col {
				// fmt.Printf("n:rc=%d,%d %d,%d %v\n", r, c, r0, c0, world.IsAliveCell(r0, c0))
				if world.IsAliveCell(r0, c0) {
					neighbors++
				}
			}
		}
	}
	return neighbors
}
func (world *World) NextState(row, col int) Health {
	var next Health = DeadJim
	var neighbors = world.Neighbors(row, col)
	if world.IsAliveCell(row, col) {
		if neighbors < 2 {
			next = Undrpop
		} else if neighbors >= 2 && neighbors <= 3 {
			next = Healthy
		} else if neighbors > 3 {
			next = Overpop
		}
	} else {
		if neighbors == 3 {
			next = Newborn
		}
	}
	return next
}
func (world *World) Evaluate() {
	for row := range world.cells {
		for col := range world.cells[row] {
			world.cells[row][col].next = world.NextState(row, col)
		}
	}

}
func (world *World) Evolve() {
	world.Evaluate()
	for row := range world.cells {
		for col := range world.cells[row] {
			world.cells[row][col].Evolve()
		}
	}
}
func (world *World) Show(verbose ...bool) {
	var v bool
	if len(verbose) > 0 {
		v = verbose[0]
	}
	for row := range world.cells {
		for col := range world.cells[row] {
			fmt.Printf("%s ", world.cells[row][col].String(v))
		}
		fmt.Println()
	}
}

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

func main() {
	// Define command-line options
	dimp := flag.Int("dim", 9, "Dimension for n x n world")
	genp := flag.Int("generations", -1, "Number of generations (-1=forever)")
	maxp := flag.Int("max", 100, "Maximum number of cells")
	pausep := flag.Bool("pause", false, "Pause between steps")
	thresholdp := flag.Int("threshold", 50, "Initial seed (0=fully populated)")
	seedp := flag.Int("seed", -1, "Initial rng seed")
	// debugp := flag.Bool("debug", false, "Enable debug output")

	if *seedp > 0 {
		rand.Seed(int64(*seedp))
	} else {
		rand.Seed(time.Now().UTC().UnixNano())
	}

	// Parse
	flag.Parse()

	var world = NewWorld(*dimp)
	world.PopulateWithThreshold(*thresholdp, *maxp)
	// world.AddPattern(1, *dimp-3, TRIONIMO_V)
	world.AddPattern(*dimp-3, *dimp-3, GLIDER)

	scrReset()
	scrPosition(0, 0)
	scrEraseDisplay()
	fmt.Printf("--[ game-of-life: %d x %d, gens=%d, pause=%v, threshold=%d, max=%d ]--", *dimp, *dimp, *genp, *pausep, *thresholdp, *maxp)
	var iter = 0
	for true {
		var pop = world.Population()
		scrPosition(2, 0)
		scrEraseLine()
		fmt.Printf("[ iter:%d of %d, pop:%d ]", iter, *genp, pop)
		scrPosition(3, 0)
		world.Evolve()
		world.Show()
		iter++
		time.Sleep(100 * time.Millisecond)
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
}
