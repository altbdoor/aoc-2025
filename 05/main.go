package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	// open input file relative from script
	inputPath := "input.txt"
	if len(os.Args) > 2 {
		inputPath = os.Args[2]
	}

	_, filename, _, _ := runtime.Caller(0)
	inputPath = filepath.Join(filepath.Dir(filename), inputPath)
	file, _ := os.Open(inputPath)
	defer file.Close()

	// conditionally run handler
	handlerStr := os.Args[1]
	switch handlerStr {
	case "1":
		handle1(file)
	case "2":
		handle2(file)
	default:
		log.Fatal("invalid arg")
	}
}

type Pair struct {
	start int64
	end   int64
}

func handle1(file *os.File) {
	scanner := bufio.NewScanner(file)

	checkRangeDone := false
	goodCounter := 0
	var allowedPair []Pair

	timer := time.Now()

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			timeTaken := time.Since(timer).Seconds()
			timeTakenStr := fmt.Sprintf("%.4f", timeTaken)

			fmt.Println("finish checking ranges in", timeTakenStr, "seconds")
			checkRangeDone = true
			continue
		}

		if !checkRangeDone {
			rangeData := strings.SplitN(line, "-", 2)
			rangeStart, _ := strconv.ParseInt(rangeData[0], 10, 64)
			rangeEnd, _ := strconv.ParseInt(rangeData[1], 10, 64)

			pair := Pair{start: rangeStart, end: rangeEnd}
			allowedPair = append(allowedPair, pair)
		} else {
			lineInt, _ := strconv.ParseInt(line, 10, 64)

			for _, pair := range allowedPair {
				if lineInt >= pair.start && lineInt <= pair.end {
					goodCounter++
					break
				}
			}
		}
	}

	fmt.Println("fresh:", goodCounter)
}

/*
retro: honestly not sure what went *right*. i got the comparator, and merging
logic (or so i believe?). the answer was still incorrect, so i checked reddit,
saw some clues about sorting, proceeded to add it, and got the result.

references:
- https://www.reddit.com/r/adventofcode/comments/1pemdwd/2025_day_5_solutions/
*/
func handle2(file *os.File) {
	scanner := bufio.NewScanner(file)
	var allowedPair []Pair

	var lowestIdx int64 = 0
	var highestIdx int64 = 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		rangeData := strings.SplitN(line, "-", 2)
		rangeStart, _ := strconv.ParseInt(rangeData[0], 10, 64)
		rangeEnd, _ := strconv.ParseInt(rangeData[1], 10, 64)

		if lowestIdx == 0 {
			lowestIdx = rangeStart
		} else if rangeStart < lowestIdx {
			lowestIdx = rangeStart
		}

		if rangeEnd > highestIdx {
			highestIdx = rangeEnd
		}

		pair := Pair{start: rangeStart, end: rangeEnd}
		allowedPair = append(allowedPair, pair)
	}

	sort.Slice(allowedPair, func(a, b int) bool {
		return allowedPair[a].start < allowedPair[b].start
	})

	for true {
		var compressedPair []Pair
		hasIntersect := false

		for _, pair := range allowedPair {
			compressedPairLen := len(compressedPair)

			if compressedPairLen == 0 {
				compressedPair = append(compressedPair, pair)
				continue
			}

			intersectIdx := -1
			for idx := range compressedPairLen {
				checkPair := compressedPair[idx]

				if (pair.start >= checkPair.start && pair.start <= checkPair.end) ||
					(pair.end >= checkPair.start && pair.end <= checkPair.end) {
					intersectIdx = idx
					break
				}
			}

			if intersectIdx == -1 {
				compressedPair = append(compressedPair, pair)
				continue
			}

			hasIntersect = true
			fmt.Println("pair", pair)
			fmt.Println("before", compressedPair[intersectIdx])

			if pair.start < compressedPair[intersectIdx].start {
				compressedPair[intersectIdx].start = pair.start
			}

			if pair.end > compressedPair[intersectIdx].end {
				compressedPair[intersectIdx].end = pair.end
			}

			fmt.Println("after", compressedPair[intersectIdx])
		}

		sort.Slice(allowedPair, func(a, b int) bool {
			return allowedPair[a].start < allowedPair[b].start
		})
		allowedPair = compressedPair

		if !hasIntersect {
			break
		}
	}

	fmt.Println(strings.Repeat("=", 40))
	var sum int64 = 0
	for _, pair := range allowedPair {
		fmt.Println(pair)
		sum += pair.end - pair.start + 1
	}

	fmt.Println("sum:", sum)

}
