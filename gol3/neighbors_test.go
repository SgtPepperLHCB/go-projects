package main

import (
	"fmt"
	"testing"
)

func compare(n0, n1 World) bool {
	var size = len(n0.cells)
	var result bool = true
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			result = result && (n0.cells[x][y].state == n1.cells[x][y].state)
		}
	}

	return result
}

func newWorld(s StringWorld) World {
	var size = len(s)
	var world = NewWorld(size)
	for x := range s {
		for y := range s[x] {
			world.cells[x][y] = NewCell(x, y, s[x][y])
		}
	}
	return world
}

var D string = DEAD
var H string = HLTH
var N string = NEWB
var O string = OVER
var U string = UNDR

func TestAddPattern(t *testing.T) {
	var data = StringWorld{
		{N, N, D, N},
		{D, D, D, D},
		{D, D, D, D},
		{N, D, D, N},
	}
	fmt.Println("---[ pattern ]---")
	var size = len(data)
	var world = NewWorld(size * 4)
	world.populate(200)
	world.addPattern(size*3, size*3, data)
	world.show()
}
func TestNeighbors(t *testing.T) {
	var data = [][]StringWorld{
		//t0
		{
			{
				{D, D, D, D, D},
				{D, D, N, D, D},
				{D, D, N, D, D},
				{D, D, N, D, D},
				{D, D, D, D, D},
			},
			{
				{D, D, D, D, D},
				{D, H, H, D, D},
				{D, D, D, H, D},
				{D, H, H, D, D},
				{D, D, D, D, D},
			},
		},
		//t1
		{
			{
				{N, N, D, N},
				{D, D, D, D},
				{D, D, D, D},
				{N, D, D, N},
			},
			{
				{O, H, N, U},
				{N, D, D, D},
				{D, D, D, D},
				{U, D, D, U},
			},
		},
		//t2
		{
			{
				{D, N, N},
				{D, D, D},
				{D, N, N},
			},
			{
				{N, H, N},
				{N, D, D},
				{D, D, D},
			},
		},
		//t2
		{
			{
				{D, N, D},
				{D, N, D},
				{D, N, D},
			},
			{
				{N, H, N},
				{N, D, D},
				{D, D, D},
			},
		},
	}
	for i := range data {
		var n = newWorld(data[i][0])
		var e = newWorld(data[i][1])

		fmt.Println("-[ before ]-")
		n.show()

		for j := 0; j < 5; j++ {
			n.evolve(false)
			fmt.Printf("-[ evolved %d ]-", j)
			n.show()
		}

		fmt.Println("-[ expected ]-")
		e.show()
		fmt.Printf("same ? %v\n\n", compare(n, e))
		// if !compare(n, e) {
		// 	t.Error(
		// 		"\nexpected\n", e.showString(),
		// 		"\nactual\n", n.showString(),
		// 	)
		// }
	}

}
