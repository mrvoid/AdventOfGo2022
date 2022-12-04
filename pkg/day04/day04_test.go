package day04

import (
    "testing"
    utils "advent2022/utils"
)

func TestContainment(t *testing.T) {
    have := contains(section{1,10}, section{9,10})
    want := true
    if have != want {
        t.Fatalf(`Was expecting %t for contains(section{1,10}, section{9,10}) and instead got %t`, want, have)
    }

}

func TestAns1(t *testing.T) {
	lines := utils.ReadLines("./testdata/input_test.txt")

    ans1 := solve1(lines)
    want := 2

    if ans1 != want {
        t.Fatalf(`Solution has given %d and should have given %d`, ans1, want )
    }
}

func TestAns2(t *testing.T) {
    lines := utils.ReadLines("./testdata/input_test.txt")

    ans2 := solve2(lines)
    want := 4

    if ans2 != want {
        t.Fatalf(`Solution has given %d and should have given %d`, ans2, want )
    }
}
