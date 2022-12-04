package day01

import (
    "testing"
    "os"
    // "regexp"
)

func TestAns1(t *testing.T) {
    file, _ := os.Open("./testdata/input_test.txt")
    defer file.Close()

    ans1, _ := Solve(file)
    want := 24000

    if ans1 != want {
        t.Fatalf(`Solution has given %d and should have given %d`, ans1, want )
    }
}

func TestAns2(t *testing.T) {
    file, _ := os.Open("./testdata/input_test.txt")
    defer file.Close()

    _, ans2 := Solve(file)
    want := 45000

    if ans2 != want {
        t.Fatalf(`Solution has given %d and should have given %d`, ans2, want )
    }
}
