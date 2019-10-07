package main

import (
	"fmt"
	"testing"
	"github.com/dinophile/routes"
)

// TestHealth checks app health
// So I wrote this test back when I wasn't exporting my routes, and health was just a function within main
// And here's where I hit a wall, dig into http request testing or hand in the tech test?
// TODO: continue to do some research and will add once I learn more!
// Usually with a larger project I like to map out my test cases and do TDD when I'm more comfortable, but in this case 
// I rushed to implement the main program as a priority!

func TestHealth(t *testing.T) {

	expected := "You're good to...GO! :D"
	result := health()

	if result != expected {
		t.Error("Health check failed, expected %w, got %w", expected, result)
	}
}



