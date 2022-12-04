package day03

import ()

// a-z (1-26), A-Z (27-52)
func score(item byte) int{
	if item >= 'a' && item <= 'z' {
		return (int)(item - 'a' + 1)
	} else if (item >= 'A' && item <= 'Z') {
		return (int)(item - 'A' + 27)
	} else {
		return 0
	}
}

func solve1(lines []string) int {
	total := 0
	for _, line := range lines {
		var bins [256]byte
		for i, c := range line {
			var mask byte
			if (i < len(line) / 2) {
				mask = 0b01
			} else {
				mask = 0b10
			}
			bins[c] |= mask
		}

		for i, v := range bins {
			if v == 3 {
				total += score(byte(i))
				break
			}
		}
    }

	return total
}

func solve2(lines []string) int {
	total := 0
	for i := 0; i < len(lines); i += 3 {
		var bins [256]byte
		for s := 0; s < 3; s++ { // Would love to see it as an adhoc range
			mask := (byte)(0b1 << s)
			for _, c := range lines[i+s] {
				bins[c] |= mask
			}
		}

		for i, v := range bins {
			if v == 7 {
				total += score(byte(i))
				break
			}
		}
	}

	return total
}

func Solve(lines []string)  (int, int) {
	return solve1(lines), solve2(lines)
}