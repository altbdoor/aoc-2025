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
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(strconv.Atoi(line))
	}
}

func handle2(file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(strconv.Atoi(line))
	}
}
