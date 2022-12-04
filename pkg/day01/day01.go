package day01

import (
    "os"
	"io"
    "strconv"
    "bufio"
    "sort"
)

func Solve(file *os.File)  (int, int) {
	reader := bufio.NewReader(file)
    highscore := 0
    var elves []int

    total := 0
    for {
        
        line, _, err := reader.ReadLine()
        
        if len(line) == 0 {
            if total > highscore {
                highscore = total
            }
            elves = append(elves, total)
            total = 0
        } else {
            calories, _ := strconv.Atoi(string(line))
            total += calories
        }

        if err == io.EOF {
            if total > highscore {
                highscore = total
            }
            elves = append(elves, total)
            break
        }
    }
    sort.Ints(elves)

	return highscore, elves[len(elves) - 1] + elves[len(elves) - 2] + elves[len(elves) - 3]
}