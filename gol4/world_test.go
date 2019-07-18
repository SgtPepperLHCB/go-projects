package main

import (
	"fmt"
	"testing"
)

func TestCanCreateNewWorld(t *testing.T) {
	var world = NewWorld(5)
	if world.Size() != 5 {
		t.Error(
			"expected size", 5,
			"actual", world.Size(),
		)
	}
}

func TestCanPopulateWorld(t *testing.T) {
	const SIZE = 5
	var world = NewWorld(SIZE)

	if world.Population() != 0 {
		t.Error(
			"expected", 0,
			"actual", world.Population(),
		)
	}

	var pop = 5
	world.PopulateWithThreshold(0, pop)
	if world.Population() != pop {
		t.Error(
			"expected pop", pop,
			"actual pop", world.Population(),
		)
	}

	pop = 15
	world.PopulateWithThreshold(0, pop)
	if world.Population() != pop {
		t.Error(
			"expected pop", pop,
			"actual pop", world.Population(),
		)
	}

	world.Populate()
	if world.Population() > 0 {
		t.Error(
			"expected pop", 0,
			"actual pop", world.Population(),
		)
	}
	pop = 3
	world.AddPattern(1, 1, TRIONIMO_V)
	if world.Population() != pop {
		t.Error(
			"expected pop", pop,
			"actual pop", world.Population(),
		)
	}
}

func TestNeighbors(t *testing.T) {
	const SIZE = 5
	var world = NewWorld(SIZE)

	world.AddPattern(1, 2, TRIONIMO_V)
	// var show = false
	// world.Show(show)
	fmt.Println()
	var neighbors int
	// Test in center (no wrapping)
	neighbors = world.Neighbors(2, 2)
	if neighbors != 2 {
		t.Error(
			"expected", 2,
			"actual", neighbors,
		)
	}

	// Test in corners (wrapping)
	neighbors = world.Neighbors(0, 0)
	if neighbors != 0 {
		t.Error(
			"expected", 0,
			"actual", neighbors,
		)
	}
	neighbors = world.Neighbors(SIZE-1, 0)
	if neighbors != 0 {
		t.Error(
			"expected", 0,
			"actual", neighbors,
		)
	}
	neighbors = world.Neighbors(0, SIZE-1)
	if neighbors != 0 {
		t.Error(
			"expected", 0,
			"actual", neighbors,
		)
	}
	neighbors = world.Neighbors(SIZE-1, SIZE-1)
	if neighbors != 0 {
		t.Error(
			"expected", 0,
			"actual", neighbors,
		)
	}

	// Top/bottom/left/right middle edge
	neighbors = world.Neighbors(0, SIZE/2)
	if neighbors != 1 {
		t.Error(
			"expected", 1,
			"actual", neighbors,
		)
	}
	neighbors = world.Neighbors(SIZE/2, 0)
	if neighbors != 0 {
		t.Error(
			"expected", 0,
			"actual", neighbors,
		)
	}
	neighbors = world.Neighbors(SIZE/2, SIZE-1)
	if neighbors != 0 {
		t.Error(
			"expected", 0,
			"actual", neighbors,
		)
	}
	neighbors = world.Neighbors(SIZE/2, SIZE-1)
	if neighbors != 0 {
		t.Error(
			"expected", 0,
			"actual", neighbors,
		)
	}
}
