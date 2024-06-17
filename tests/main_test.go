package test

import (
	sla "module"
	"testing"
)

// TestAdd tests the Add function
func TestAdd(t *testing.T) {
	result := sla.Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}

// TestSubtract tests the Subtract function
func TestSubtract(t *testing.T) {
	result := sla.Subtract(5, 3)
	expected := 2
	if result != expected {
		t.Errorf("Subtract(5, 3) = %d; want %d", result, expected)
	}
}
