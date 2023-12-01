package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func main() {
    startTime := time.Now()

    lines, err := readLines("../inputs/input.txt")
    if err != nil {
        fmt.Println("Error reading input file:", err)
        return
    }

    total := computeTotal(lines)
    fmt.Println(total)

    endTime := time.Now()
    fmt.Println("Program execution time:", endTime.Sub(startTime))

    printMemoryUsage()
}

// computeTotal computes the total sum of numbers in the input lines.
func computeTotal(lines []string) int {
    re := regexp.MustCompile("[0-9]|zero|one|two|three|four|five|six|seven|eight|nine")
    total := 0
    for _, line := range lines {
        digitSlice := re.FindAllString(line, -1)
        if len(digitSlice) > 0 {
            first := digitSlice[0]
            last := digitSlice[len(digitSlice)-1]
            if len(first) > 1 {
                first = wordToNum(first)
            }
            if len(last) > 1 {
                last = wordToNum(last)
            }
            num, _ := strconv.Atoi(first + last)
            total += num
            fmt.Println("First: ", first, " | Last: ", last, " | Num: ", num, " | Total: ", total)
        }
    }
    return total
}

// readLines reads the lines from the given file.
func readLines(filename string) ([]string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    reader := bufio.NewReader(file)
    for {
        line, err := reader.ReadString('\n')
        if err != nil && err != io.EOF {
            return nil, err
        }
        lines = append(lines, strings.TrimSpace(line))
        if err == io.EOF {
            break
        }
    }

    return lines, nil
}

// wordToNum converts a word representation of a number to its corresponding numerical value.
func wordToNum(word string) string {
    wordNumMap := map[string]string{
        "zero":  "0",
        "one":   "1",
        "two":   "2",
        "three": "3",
        "four":  "4",
        "five":  "5",
        "six":   "6",
        "seven": "7",
        "eight": "8",
        "nine":  "9",
    }

    return wordNumMap[word]
}

// printMemoryUsage prints the current memory usage.
func printMemoryUsage() {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    fmt.Printf("Memory Usage: %v bytes\n", m.Alloc)
}
