package main

import "testing"

func TestShouldReturnMessage(t *testing.T) {
	result, _, _, _ := concatWords("Hello", "Nice", "World")
	if result != "HelloNiceWorld" {
		t.Error("Failed Error!")
	}
}
