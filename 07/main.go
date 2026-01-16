package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
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

func handle1(file *os.File) {
	scanner := bufio.NewScanner(file)
	inputArr := make([]string, 0)
	sourceIdx := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()

		if len(sourceIdx) == 0 {
			currentSourceIdx := strings.IndexRune(line, 'S')
			sourceIdx = append(sourceIdx, currentSourceIdx)
			continue
		}

		inputArr = append(inputArr, line)
	}

	splitCounter := 0

	/*
		retro: i could have done this in the loop above directly, rather than
		wasting precious memory!
	*/
	for _, line := range inputArr {
		// copy source idx, and nuke the original
		copySourceIdx := make([]int, len(sourceIdx))
		copy(copySourceIdx, sourceIdx)
		sourceIdx = make([]int, 0)

		// use a map to keep track of unique values
		sourceMap := make(map[int]bool)

		// loop through the source
		for _, source := range copySourceIdx {
			switch line[source] {
			case '^':
				// if we hit a splitter, increment, and add two new source
				splitCounter++
				sourceMap[source-1] = true
				sourceMap[source+1] = true
			case '.':
				// if we don't hit a splitter, carry source down
				sourceMap[source] = true

				/*
					retro: i actually forgot to carry source down, until i print debug
					the diagram!
				*/
			}
		}

		// convert the map back to array
		for source := range sourceMap {
			sourceIdx = append(sourceIdx, source)
		}

		// if we never hit any splitter, means all beams continue
		if len(sourceIdx) == 0 {
			sourceIdx = copySourceIdx
		}

		// debug
		// for idx := range len(line) {
		// 	if line[idx] == '^' {
		// 		fmt.Print("^")
		// 	} else if slices.Contains(sourceIdx, idx) {
		// 		fmt.Print("|")
		// 	} else {
		// 		fmt.Print(".")
		// 	}
		// }
		// fmt.Println()
	}

	fmt.Println("counter:", splitCounter)
}

func handle2(_ *os.File) {
	fmt.Println("dnf")
}
