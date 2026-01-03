package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
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

func handle1(file *os.File) {
	spaceSplitRe := regexp.MustCompile(`\s+`)
	intRe := regexp.MustCompile(`\d+`)

	var opArr []string
	var intArr [][]int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		columns := spaceSplitRe.Split(strings.TrimSpace(line), -1)

		if intRe.MatchString(columns[0]) {
			tempIntArr := make([]int, len(columns))

			for idx, col := range columns {
				colInt, _ := strconv.Atoi(col)
				tempIntArr[idx] = colInt
			}

			intArr = append(intArr, tempIntArr)
		} else {
			opArr = columns
		}
	}

	sumArr := make([]int, len(opArr))
	for idx, op := range opArr {
		sumArr[idx] = 0

		for _, intData := range intArr {
			currentIntData := intData[idx]

			if sumArr[idx] == 0 {
				sumArr[idx] += currentIntData
				continue
			}

			switch op {
			case "*":
				sumArr[idx] *= currentIntData
			case "+":
				sumArr[idx] += currentIntData
			default:
				log.Fatal("undefined op")
			}
		}
	}

	sum := 0
	for _, currentSum := range sumArr {
		sum += currentSum
	}

	fmt.Println("sum:", sum)
}

func handle2(file *os.File) {
	var inputArr []string
	lineLen := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		inputArr = append(inputArr, line)

		if lineLen == 0 {
			lineLen = len(line)
		}
	}

	var cacheStr strings.Builder
	var cacheInt []int
	sum := 0

	for idx := range lineLen {
		cacheStr.Reset()
		colIdx := lineLen - idx - 1
		operand := ""

		for _, line := range inputArr {
			currentChar := line[colIdx]

			if currentChar == '*' || currentChar == '+' {
				operand = string(currentChar)
			} else {
				cacheStr.WriteByte(currentChar)
			}
		}

		currentStr := strings.TrimSpace(cacheStr.String())
		currentInt, _ := strconv.Atoi(currentStr)
		cacheInt = append(cacheInt, currentInt)

		if operand == "" {
			continue
		}

		// operand found, do all calcs
		currentSum := 0
		for _, currentInt := range cacheInt {
			if currentSum == 0 {
				currentSum += currentInt
				continue
			}

			switch operand {
			case "*":
				currentSum *= currentInt
			case "+":
				currentSum += currentInt
			}
		}

		// purge cached ints
		cacheInt = make([]int, 0)
		sum += currentSum
	}

	fmt.Println("sum:", sum)
}
