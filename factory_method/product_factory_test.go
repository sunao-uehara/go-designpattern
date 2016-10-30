package main

import (
	"testing"
)

func TestRegister(t *testing.T) {
	// redeclare p1
	Register("p1", NewP1())

	// sad path
	// Register("invalid", nil)
}