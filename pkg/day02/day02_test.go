package day02

import (
    "testing"
    // "regexp"
    utils "advent2022/utils"
)

func TestScore(t *testing.T) {
	if score(0, 0) != 4 {
		t.Fatalf(`Rock vs Rock should have given 4 instead of %d`, score(0, 0))
	}
	if score(0, 1) != 1 {
		t.Fatalf(`Rock vs Paper should have given 1 instead of %d`, score(0, 1))
	}
	if score(0, 2) != 7 {
		t.Fatalf(`Rock vs Scissors should have given 7 instead of %d`, score(0, 2))
	}
}

func TestAns1(t *testing.T) {
	lines := utils.ReadLines("./testdata/input_test.txt")

    ans1, _ := Solve(lines)
    want := 15

    if ans1 != want {
        t.Fatalf(`Solution has given %d and should have given %d`, ans1, want )
    }
}

func TestAns2(t *testing.T) {
    lines := utils.ReadLines("./testdata/input_test.txt")

    _, ans2 := Solve(lines)
    want := 12

    if ans2 != want {
        t.Fatalf(`Solution has given %d and should have given %d`, ans2, want )
    }
}
