package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Cell string
type World [][]Cell

var RESET string = "\u001b[0m"
var CLEAR string = "\u001b[2J"
var HOME0 string = "\u001b[1;0H"
var HOME1 string = "\u001b[2;0H"
var HEALTHY Cell = "*"
var NEWBORN Cell = "^"
var UNDER Cell = ""
var OVER Cell = ""

func minmax(v, x, y int) int {
	if v < x {
		return x
	} else if v > y {
		return y
	} else {
		return v
	}
}

func newWorld(dim int) World {
	world := make(World, dim)
	for r := 0; r < dim; r++ {
		world[r] = make([]Cell, dim)
		for c := 0; c < dim; c++ {
			if rand.Intn(100) > 95 {
				world[r][c] = NEWBORN
			} else {
				world[r][c] = ""
			}
		}
	}
	return world
}

func printWorld(world World) {
	dim := len(world)
	for r := 0; r < dim; r++ {
		for c := 0; c < dim; c++ {
			fmt.Printf("%1s ", world[r][c])
		}
		fmt.Println()
	}
}

func deepCopy(world World) World {
	var dim = len(world)
	var newWorld = make(World, dim)
	for r := 0; r < dim; r++ {
		newWorld[r] = make([]Cell, dim)
		for c := 0; c < dim; c++ {
			newWorld[r][c] = world[r][c]
		}
	}

	return newWorld
}

func stepWorld(world World) World {
	d := len(world)
	var newWorld = deepCopy(world)
	for r := 0; r < d; r++ {
		for c := 0; c < d; c++ {
			var state = nextState(world, r, c)
			newWorld[r][c] = state
		}
	}

	return newWorld
}

func population(world World, row, col, r0, r1, c0, c1 int) int {
	var pop = 0
	for r := r0; r <= r1; r++ {
		for c := c0; c <= c1; c++ {
			if r == row && c == col {
			} else {
				if isAlive(world[r][c]) {
					pop++
				}
			}
		}
	}

	return pop
}
func isAlive(cell Cell) bool {
	return cell == HEALTHY || cell == NEWBORN
}

func nextState(world World, row, col int) Cell {
	var max = len(world[row]) - 1
	var val = world[row][col]

	var r0 = minmax(row-1, 0, max)
	var r1 = minmax(row+1, 0, max)
	var c0 = minmax(col-1, 0, max)
	var c1 = minmax(col+1, 0, max)
	var pop = population(world, row, col, r0, r1, c0, c1)
	if !isAlive(val) && pop == 3 {
		val = NEWBORN
	} else if pop < 2 {
		val = UNDER
	} else if pop >= 2 && pop <= 3 {
		val = HEALTHY
	} else if pop > 3 {
		val = OVER
	}

	return val
}

func worldPopulation(world World) int {
	var d = len(world)
	var pop = 0
	for r := 0; r < d; r++ {
		for c := 0; c < d; c++ {
			if isAlive(world[r][c]) {
				pop++
			}
		}
	}
	return pop
}

//
func main2() {
	rand.Seed(time.Now().UTC().UnixNano())

	// Define flags
	dimp := flag.Int("dim", 9, "Dimension for n x n world")
	iterp := flag.Int("iter", 9, "Iterations")
	pausep := flag.Bool("pause", false, "Pause between steps")

	// Parse
	flag.Parse()

	fmt.Printf("%s%s--[ game-of-life: %d x %d ]--\n", CLEAR, HOME0, *dimp, *dimp)
	world := newWorld(*dimp)

	for iter := 0; iter < *iterp; iter++ {
		fmt.Printf("%s%s[%2d / %2d] pop:%4d\n", "", HOME1, iter, *iterp, worldPopulation(world))
		printWorld(world)
		time.Sleep(200 * time.Millisecond)
		if worldPopulation(world) < 1 {
			fmt.Println("--[ dead world ]--")
			break
		}
		if *pausep {
			bufio.NewReader(os.Stdin).ReadBytes('\n')
		}
		world = stepWorld(world)
	}
}
