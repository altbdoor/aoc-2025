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

	fmt.Println(">>", totalJolt)
}

func handle2(file *os.File) {
	scanner := bufio.NewScanner(file)
	totalJolt := 0

	for scanner.Scan() {
		line := scanner.Text()
		lineLen := len(line)

		MAX_DIGIT := 12
		lastMaxIdx := -1

		maxJolt := ""

		for i := range MAX_DIGIT {
			lastLineIdx := lineLen + 1 - (MAX_DIGIT - i)
			iterSubset := ""

			if i == 0 {
				iterSubset = line[0:lastLineIdx]
			} else {
				/*
					retro: should be able to check, if the text remaining length is equal
					to whatever that's required during the loop, we can exit early
				*/
				iterSubset = line[lastMaxIdx+1 : lastLineIdx]
			}

			firstMaxInt := 0
			firstMaxStr := ""
			currentLastMaxIdx := lastMaxIdx + 1

			for idx, ch := range iterSubset {
				chStr := string(ch)
				chInt, _ := strconv.Atoi(chStr)

				if chInt > firstMaxInt {
					firstMaxInt = chInt
					firstMaxStr = chStr

					if i == 0 {
						lastMaxIdx = idx
					} else {
						lastMaxIdx = idx + currentLastMaxIdx
					}
				}
			}

			if firstMaxStr == "" {
				log.Fatal("failed to find max int")
			}

			maxJolt += firstMaxStr
		}

		if len(maxJolt) != 12 {
			log.Fatal("max jolt is not 12 char")
		}

		maxJoltInt, _ := strconv.Atoi(maxJolt)
		totalJolt += maxJoltInt
	}

	fmt.Println(">>", totalJolt)
}
