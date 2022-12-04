package day03

import (
    "testing"
    utils "advent2022/utils"
)

func TestScore(t *testing.T) {
    m := map[byte]int {
        'a': 1,
        'c': 3,
        'z': 26,
        'A': 27,
        'D': 30,
        'Z': 52,
    }

    for key, value := range m {
        have := score(key)
        want := value
        
        if have != want {
            t.Fatalf(`Was expecting %d for %c and instead got %d`, want, key, have)
        }
    }
}

func TestAns1(t *testing.T) {
	lines := utils.ReadLines("./testdata/input_test.txt")

    ans1 := solve1(lines)
    want := 157

    if ans1 != want {
        t.Fatalf(`Solution has given %d and should have given %d`, ans1, want )
    }
}

func TestAns2(t *testing.T) {
    lines := utils.ReadLines("./testdata/input_test.txt")

    ans2 := solve2(lines)
    want := 70

    if ans2 != want {
        t.Fatalf(`Solution has given %d and should have given %d`, ans2, want )
    }
}
