package main

import (
	"testing"
)

func TestEvenAndOdd(t *testing.T) {
	even := "even"
	odd := "odd"
	actual := evenOrOdd(1)
	if isEqual(odd, actual) == false {
		t.Errorf("Expected %v, but got %v", odd, actual)
	}
	actual = evenOrOdd(2)
	if isEqual(even, actual) == false {
		t.Errorf("Expected %v, but got %v", odd, actual)
	}
	actual = evenOrOdd(3)
	if isEqual(even, actual) == false {
		t.Errorf("Expected %v, but got %v", even, actual)
	}
}

func isEqual(expected string, actual string) bool {
	if actual != expected {
		return false
	}
	return true
}
