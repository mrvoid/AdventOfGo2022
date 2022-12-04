package day02

import (
    //"fmt"
)

// Rock 0, Paper 1, Scissors 2
func score(me int, him int) int{
	s := me + 1
	
	if me == him {
		s += 3
	} else if (me + 1) % 3 == him {
		s += 0 // Haha
	} else {
		s += 6
	}
	return s
}

func solve1(lines []string) int {
	total := 0
	for _, line := range lines {
		total += score((int)(line[2] - 'X'), (int)(line[0]-'A'))
    }

	return total
}

func solve2(lines []string) int {
	total := 0
	for _, line := range lines {
		switch strat := line[2]; strat {
		case 'X':
			total += score((int)((line[0]-'A') + 2) % 3, (int)(line[0]-'A'))
		case 'Y':
			total += score((int)(line[0]-'A'), (int)(line[0]-'A'))
		case 'Z':
			total += score( (int)((line[0]-'A') + 1) % 3, (int)(line[0]-'A'))
		}
    }

	return total
}

func Solve(lines []string)  (int, int) {
	return solve1(lines), solve2(lines)
}