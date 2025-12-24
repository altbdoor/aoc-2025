package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
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

type Handle1Coord struct {
	start int
	stop  int
}

func handle1(file *os.File) {
	scanner := bufio.NewScanner(file)
	var coordList []Handle1Coord

	for scanner.Scan() {
		line := scanner.Text()
		currentText := strings.Split(line, ",")

		for _, text := range currentText {
			cleanedText := strings.TrimSpace(text)
			if cleanedText == "" {
				continue
			}

			cleanedArr := strings.SplitN(cleanedText, "-", 2)
			startVal, _ := strconv.Atoi(cleanedArr[0])
			stopVal, _ := strconv.Atoi(cleanedArr[1])

			coord := Handle1Coord{
				start: startVal,
				stop:  stopVal,
			}
			coordList = append(coordList, coord)
		}
	}

	sumOfBrokenIds := 0

	for _, coord := range coordList {
		fmt.Printf("group %d to %d\n  >> ", coord.start, coord.stop)
		checkInt := coord.start - 1

		for checkInt != coord.stop {
			checkInt += 1

			checkStr := strconv.Itoa(checkInt)
			checkStrLen := len(checkStr)

			if checkStrLen%2 != 0 {
				continue
			}

			midPoint := checkStrLen / 2
			left := checkStr[:midPoint]
			right := checkStr[midPoint:]

			isDifferent := false
			for idx, chr := range left {
				rightChr := rune(right[idx])

				if chr != rightChr {
					isDifferent = true
					break
				}
			}

			if !isDifferent {
				fmt.Printf("%d, ", checkInt)
				sumOfBrokenIds += checkInt
			}
		}

		fmt.Println()
		fmt.Println()
	}

	fmt.Printf("sum: %d", sumOfBrokenIds)
	fmt.Println()
}

func handle2(file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(strconv.Atoi(line))
	}
}
