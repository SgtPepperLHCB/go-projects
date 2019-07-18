package main

import (
	"testing"
)

func TestHealthEnum(t *testing.T) {
	var h Health = Newborn

	if !h.IsAlive() {
		t.Error(
			"expected", true,
			"actual", h.IsAlive(),
		)
	}
	if h.IsDead() {
		t.Error(
			"expected", false,
			"actual", h.IsDead(),
		)
	}
	if h != Newborn {
		t.Error(
			"expected", Newborn,
			"got", h,
		)
	}
	if h.String() != "+" {
		t.Error("expected", "+", "got", h.String())
	}
}
