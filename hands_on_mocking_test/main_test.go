package main

import "testing"

func TestAdd(t *testing.T) {
	total := Add(5, 4)

	if total != 9 {
		t.Errorf("Error occured, expected %d, got %d", 9, total)
	}

}
func TestSubtract(t *testing.T) {
	total := Subtract(10, 9)
	if total != 1 {
		t.Errorf("Error occured, expected %d, got %d", 1, total)
	}
}

func TestDoMaths(t *testing.T) {
	total := DoMaths(10, 10, Add)
	if total != 20 {
		t.Errorf("Error occured, expected %d, got %d", 20, total)
	}
}
