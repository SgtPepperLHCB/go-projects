package main

import "testing"

func TestCanCreateNewCell(t *testing.T) {
	var cell = NewCell(DeadJim)
	if cell.IsAlive() {
		t.Error(
			"expected", true,
			"actual", cell.IsAlive(),
		)
	}
}
