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
	totalJolt := 0

	for scanner.Scan() {
		line := scanner.Text()
		numList := make([]int, len(line))

		maxInt := 0
		maxIdx := -1
		lastIdx := len(line) - 1

		for idx, ch := range line {
			chInt, _ := strconv.Atoi(string(ch))
			numList[idx] = chInt

			// max value must not be last item
			if chInt > maxInt && idx != lastIdx {
				maxInt = chInt
				maxIdx = idx
			}
		}

		if maxIdx == -1 {
			log.Fatal("unable to find first max value")
		}

		maxJolt := 0

		for i := maxIdx + 1; i < len(line); i++ {
			currentMaxJolt := maxInt*10 + numList[i]
			fmt.Printf("%d, ", currentMaxJolt)

			if currentMaxJolt > maxJolt {
				maxJolt = currentMaxJolt
			}
		}

		if maxJolt == 0 {
			log.Fatal("unable to find max jolt")
		}

		fmt.Printf("max=%d \n", maxJolt)
		totalJolt += maxJolt
	}

	fmt.Println(totalJolt)
}

func handle2(file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(strconv.Atoi(line))
	}
}
