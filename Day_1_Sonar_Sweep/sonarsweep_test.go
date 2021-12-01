package main

import "testing"

func TestSecondGreaterThanFirst(t *testing.T) {
	res := CheckIncreased([]int{1, 2})
	if res != 1 {
		t.Fatal("expected increase count to be 1")
	}
}

func TestSecondSameAsFirst(t *testing.T) {
	res := CheckIncreased([]int{1, 1})
	if res != 0 {
		t.Fatal("expected increase count to be 0")
	}
}

func TestSecondLessThanFirst(t *testing.T) {
	res := CheckIncreased([]int{12, 1})
	if res != 0 {
		t.Fatal("expected increase count to be 0")
	}
}

func TestThirdGreaterThatFirst(t *testing.T) {
	res := CheckIncreased([]int{12, 1, 14})
	if res != 1 {
		t.Fatal("expected increase count to be 0")
	}
}

func TestExample(t *testing.T) {
	//res := CheckIncreased([]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263})
	res := CheckIncreased([]int{199, 200, 208, 210})
	if res != 3 {
		t.Fatalf("expected increase count to be 4 but was %d", res)
	}
}

func TestWave(t *testing.T) {
	res := CheckIncreased([]int{12, 1, 14, 10, 42})
	if res != 2 {
		t.Fatal("expected increase count to be 0")
	}
}
