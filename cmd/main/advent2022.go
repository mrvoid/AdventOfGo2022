package main

import (
    "fmt"
    "os"
    "flag"
    "errors"
    "advent2022/day01"
    "advent2022/day02"
    utils "advent2022/utils"
//    "advent2022/day03"
//    "advent2022/day04"
//    "advent2022/day05"
//    "advent2022/day06"
//    "advent2022/day07"
//    "advent2022/day08"
//    "advent2022/day09"
//    "advent2022/day10"
//    "advent2022/day11"
//    "advent2022/day12"
//    "advent2022/day13"
//    "advent2022/day14"
//    "advent2022/day15"
//    "advent2022/day16"
//    "advent2022/day17"
//    "advent2022/day18"
//    "advent2022/day19"
//    "advent2022/day20"
//    "advent2022/day21"
//    "advent2022/day22"
//    "advent2022/day23"
//    "advent2022/day24"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

var (
    day *int
    input  *string
)

func init() {
    day = flag.Int("d", 0, "Day to solve")
    input = flag.String("i", "", "Input file")
}

func printUsage() {
    fmt.Println("Usage: advent2022 -d <day> -i <input_file>")
}

func verifyArgs() {
    if len(*input) == 0 {
        fmt.Println("No input given")
        printUsage()
        os.Exit(3)
    }

    if *day < 0 || *day > 24 {
        fmt.Println("Day outside valid range. Enter day between 1 and 24")
        printUsage()
        os.Exit(1)
    }

    if _, err := os.Stat(*input); errors.Is(err, os.ErrNotExist) {
        fmt.Println("File not found.")
        printUsage()
        os.Exit(2)
    }
}

func main() {
    flag.Parse()
    verifyArgs()
    fmt.Println("Solving for day ", *day)

    file, err := os.Open(*input)
    check(err)
    defer file.Close()

    switch num := *day; num {
    case 1:
        answer1, answer2 := day01.Solve(file)
        fmt.Println(answer1)
        fmt.Println(answer2)
    case 2:
        answer1, answer2 := day02.Solve(utils.ReadLines(*input))
        fmt.Println(answer1)
        fmt.Println(answer2)
    default:
        fmt.Println("Day not implemented")
    }
}