package utils

import (
    "os"
	"io"
    "bufio"
)

func ReadLines(filename string) []string {
	file, _ := os.Open(filename)
    defer file.Close()

	reader := bufio.NewReader(file)
    var lines []string
	
    for {
        line, _, err := reader.ReadLine()
		lines = append(lines, string(line))

		if err == io.EOF {
            break
        }
    }

	return lines
}