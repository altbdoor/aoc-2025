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
	// prepare grid with first line
	grid := []string{""}
	colLen := 0

	// scan through file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fixedLine := "." + line + "."
		grid = append(grid, fixedLine)

		if colLen == 0 {
			colLen = len(fixedLine)
		}
	}

	// pad last line
	grid = append(grid, strings.Repeat(".", colLen))
	grid[0] = strings.Repeat(".", colLen)
	rowLen := len(grid)

	// accessible count
	totalCount := 0

	for currentRowIdx := range rowLen {
		if currentRowIdx == 0 || currentRowIdx+1 == rowLen {
			continue
		}

		beforeRowIdx := currentRowIdx - 1
		afterRowIdx := currentRowIdx + 1

		for currentColIdx := range colLen {
			if currentColIdx == 0 || currentColIdx+1 == colLen {
				continue
			}

			// skip if this is not paper
			if grid[currentRowIdx][currentColIdx] != '@' {
				fmt.Print(".")
				continue
			}

			beforeColIdx := currentColIdx - 1
			afterColIdx := currentColIdx + 1

			surroundChars := []byte{
				// row before
				grid[beforeRowIdx][beforeColIdx],
				grid[beforeRowIdx][currentColIdx],
				grid[beforeRowIdx][afterColIdx],

				// current row
				grid[currentRowIdx][beforeColIdx],
				grid[currentRowIdx][afterColIdx],

				// row after
				grid[afterRowIdx][beforeColIdx],
				grid[afterRowIdx][currentColIdx],
				grid[afterRowIdx][afterColIdx],
			}

			totalValid := strings.Count(string(surroundChars), "@")
			if totalValid < 4 {
				fmt.Print("x")
				totalCount++
			} else {
				fmt.Printf("%c", grid[currentRowIdx][currentColIdx])
			}
		}

		fmt.Println()
	}

	fmt.Print("total count: ", totalCount)
}

type GridCheckResult struct {
	nextGrid     []string
	currentCount int
}

func handle2(file *os.File) {
	// prepare grid with first line
	grid := []string{""}
	colLen := 0

	// scan through file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fixedLine := "." + line + "."
		grid = append(grid, fixedLine)

		if colLen == 0 {
			colLen = len(fixedLine)
		}
	}

	// pad last line
	grid = append(grid, strings.Repeat(".", colLen))
	grid[0] = strings.Repeat(".", colLen)

	totalCount := 0
	result := checkLines(grid)
	totalCount += result.currentCount

	for result.currentCount != 0 {
		result = checkLines(result.nextGrid)
		totalCount += result.currentCount
	}

	fmt.Print("total count: ", totalCount)
}

func checkLines(grid []string) GridCheckResult {
	colLen := len(grid[0])
	rowLen := len(grid)

	nextGrid := []string{}
	nextGrid = append(nextGrid, grid[0])

	totalCount := 0

	for currentRowIdx := range rowLen {
		var nextRow strings.Builder

		if currentRowIdx == 0 || currentRowIdx+1 == rowLen {
			continue
		}

		beforeRowIdx := currentRowIdx - 1
		afterRowIdx := currentRowIdx + 1

		for currentColIdx := range colLen {
			if currentColIdx == 0 || currentColIdx+1 == colLen {
				continue
			}

			// skip if this is not paper
			if grid[currentRowIdx][currentColIdx] != '@' {
				nextRow.WriteString(".")
				continue
			}

			beforeColIdx := currentColIdx - 1
			afterColIdx := currentColIdx + 1

			surroundChars := []byte{
				// row before
				grid[beforeRowIdx][beforeColIdx],
				grid[beforeRowIdx][currentColIdx],
				grid[beforeRowIdx][afterColIdx],

				// current row
				grid[currentRowIdx][beforeColIdx],
				grid[currentRowIdx][afterColIdx],

				// row after
				grid[afterRowIdx][beforeColIdx],
				grid[afterRowIdx][currentColIdx],
				grid[afterRowIdx][afterColIdx],
			}

			totalValid := strings.Count(string(surroundChars), "@")
			if totalValid < 4 {
				nextRow.WriteString(".")
				totalCount++
			} else {
				nextRow.WriteByte(grid[currentRowIdx][currentColIdx])
			}
		}

		nextRowStr := nextRow.String()
		nextGrid = append(nextGrid, "."+nextRowStr+".")
	}

	nextGrid = append(nextGrid, grid[0])
	result := GridCheckResult{
		nextGrid:     nextGrid,
		currentCount: totalCount,
	}

	return result
}
