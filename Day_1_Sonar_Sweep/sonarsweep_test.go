package main

import "testing"

func TestFoo(t *testing.T) {
	res := CheckIncreased([]int{1, 2})
	if res != 0 {
		t.Fatal("not 0")
	}
}
