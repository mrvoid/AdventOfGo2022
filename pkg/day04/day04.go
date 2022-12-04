package day04

import (
	"strings"
	"strconv"
)

type section struct {
	a int
	b int
}

func contains(big section, small section) bool{
	return big.a <= small.a && big.b >= small.b
}

func overlap(big section, small section) bool{
	return  big.a <= small.a && big.b >= small.a ||
			big.a <= small.b && big.b >= small.b
}

func parseSections(line string) (section, section) {
	sec_def := strings.Split(line, ",")
	sec1_ab := strings.Split(sec_def[0], "-")
	sec2_ab := strings.Split(sec_def[1], "-")

	// Looks ugly
	a1, _ := strconv.Atoi(sec1_ab[0])
	b1, _ := strconv.Atoi(sec1_ab[1])

	a2, _ := strconv.Atoi(sec2_ab[0])
	b2, _ := strconv.Atoi(sec2_ab[1])

	return section{a1, b1}, section{a2, b2}
}

func solve1(lines []string) int {
	total := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		sec1, sec2 := parseSections(line)
		if contains(sec1, sec2) || contains(sec2, sec1) {
			total ++;
		}
    }

	return total
}

func solve2(lines []string) int {
	total := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		sec1, sec2 := parseSections(line)
		if (contains(sec1, sec2) || contains(sec2, sec1)) || (overlap(sec1, sec2) || overlap(sec2, sec1)) {
			total ++;
		}
	}

	return total
}

func Solve(lines []string)  (int, int) {
	return solve1(lines), solve2(lines)
}