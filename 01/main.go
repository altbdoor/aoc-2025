package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

func main() {
	// open input file relative from script
	_, filename, _, _ := runtime.Caller(0)
	inputPath := filepath.Join(filepath.Dir(filename), "input.txt")
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
	initIdx := 50
	zeroCounter := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		moveDir := line[:1]
		moveAmt, _ := strconv.Atoi(line[1:])

		nextIdx := initIdx

		if moveDir == "L" {
			nextIdx -= moveAmt

			// loop around if less than zero
			if nextIdx < 0 {
				/*
					retro: this is incorrect to display the `nextIdx`, because a -20 will
					not return any hundreds, when in fact, it already has passed one
					cycle of hundred
				*/
				idxHundreds := (nextIdx * -1) / 100
				nextIdx += 100 * idxHundreds
			}
		} else {
			nextIdx += moveAmt
			nextIdx = nextIdx % 100
		}

		fmt.Printf("rotating %s by %d, from %d to %d\n", moveDir, moveAmt, initIdx, nextIdx)
		initIdx = nextIdx

		if initIdx == 0 {
			zeroCounter += 1
		}
	}

	fmt.Printf("zero counter: %d", zeroCounter)
}

/*
references:

- https://old.reddit.com/r/adventofcode/comments/1pbolbe/help_for_day_1_part_2/
*/
func handle2(file *os.File) {
	initIdx := 50
	zeroCounter := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		moveDir := line[:1]
		moveAmt, _ := strconv.Atoi(line[1:])

		idxHundreds := 0

		// every 100 loop will bounce through zero
		idxHundreds += moveAmt / 100
		moveAmt = moveAmt % 100

		fmt.Printf("start=%d, ", initIdx)
		fmt.Printf("move=%s, ", line)

		if idxHundreds > 0 {
			fmt.Printf("[100=%d], ", idxHundreds)
		}

		// nothing to do if no more movement
		if moveAmt == 0 {
			fmt.Printf("no move, stop")
			zeroCounter += idxHundreds
			continue
		}

		// prepare a future var
		nextIdx := initIdx

		// update next index
		if moveDir == "L" {
			nextIdx -= moveAmt
		} else {
			nextIdx += moveAmt
		}

		// initIdx can be from 0-99
		// moveAmt can be from -99 to 99
		// nextIdx can be from -99 to 198
		fmt.Printf("next=%d, ", nextIdx)

		if nextIdx == 0 || nextIdx == 100 || nextIdx == -100 {
			// exact zeroes gets a click
			fmt.Printf("[zero], ")
			idxHundreds++
			nextIdx = 0

		} else if nextIdx < 0 {
			if initIdx > 0 {
				// from a positive to a negative
				idxHundreds++
				nextIdx += 100
				fmt.Printf("[adjust=%d], ", nextIdx)

			} else {
				// bring negatives back to positives
				nextIdx += 100
				fmt.Printf("adjust=%d, ", nextIdx)
			}

		} else if nextIdx > 100 {
			// over 100 gets a click
			idxHundreds++
			nextIdx -= 100
			fmt.Printf("[adjust=%d], ", nextIdx)
		}

		fmt.Printf("stop=%d", nextIdx)

		// guard values within range
		if nextIdx < 0 || nextIdx >= 100 {
			fmt.Println()
			log.Fatalf("Error: stop=%d", nextIdx)
		}

		initIdx = nextIdx
		zeroCounter += idxHundreds
		fmt.Println()
		// fmt.Scanln()
	}

	fmt.Printf("zero counter: %d", zeroCounter)
	fmt.Println()
}
