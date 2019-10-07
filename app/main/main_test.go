package main

import (
	"fmt"
	"testing"
)

// TestHealth checks app health
func TestHealth(t *testing.T) {

	expected := "You're good to...GO! :D"
	result := health()

	if result != expected {
		t.Error("Health check failed, expected %w, got %w", expected, result)
	}
}

// TestInvert - it should return a csv where rows and columns are reversed
func TestInvert(t *testing.T) {
	input := []string{"a", "b", "c", "d"}
	expected := fmt.Printf("%s\n", "a, c, b, d")

	result := "hola"
}
