package main

import (
	"os"
	"strconv"
	"fmt"
)

func main() {

	if len(os.Args[1:]) < 2 {
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

	var algorithm int = 1

	if len(os.Args) == 4 {
		algorithm, err = strconv.Atoi(os.Args[3])

		if err != nil {
			PrintUsage(os.Args[0])
			os.Exit(-1)
		}
	}

	var maze *Maze
	var start [2]int
	var paths PriorityQueue[Stack[[2]int]]

	switch algorithm {
	case 1:
		maze, start = GenerateMaze1(height, width)
	case 2:
		maze, start = GenerateMaze2(height, width)
	default:
		fmt.Println("\nUnsupported algorithm\n")
		os.Exit(-1)
	}
	
	paths = GetAllPaths(maze, start)

	path, _ := paths.Dequeue()

	MarkSolution(maze, path)

	PrintMaze(maze)

	os.Exit(0)
}