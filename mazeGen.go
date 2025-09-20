package main

import (
	"os"
	"strconv"
)

func main() {

	if len(os.Args[1:]) != 2 {
		PrintUsage(os.Args[0])
		os.Exit(-1)
	}

	height, err := strconv.Atoi(os.Args[1])

	if err != nil {
		PrintUsage(os.Args[0])
		os.Exit(-1)
	}

	width, err := strconv.Atoi(os.Args[2])

	if err != nil {
		PrintUsage(os.Args[0])
		os.Exit(-1)
	}

	maze := GenerateMaze(height, width)
	PrintMaze(maze)
}