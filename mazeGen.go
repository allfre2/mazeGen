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

	if len(os.Args) > 3 {
		algorithm, err = strconv.Atoi(os.Args[3])

		if err != nil {
			PrintUsage(os.Args[0])
			os.Exit(-1)
		}
	}

	var dificulty int = 0

	if len(os.Args) > 4 {
		dificulty, err = strconv.Atoi(os.Args[4])

		if err != nil {
			PrintUsage(os.Args[0])
			os.Exit(-1)
		}

		if dificulty > 99 { dificulty = 99 }

		if dificulty < 0 { dificulty = 0 }
	}
	
	var showSolution bool = false

	if len(os.Args) > 5 && os.Args[5] == "-s" {
		showSolution = true
	}

	var maze *Maze
	var start [2]int

	switch algorithm {
	case 1:
		maze, start = GenerateMaze1(height, width)
	case 2:
		maze, start = GenerateMaze2(height, width)
	default:
		fmt.Println("\nUnsupported algorithm\n")
		os.Exit(-1)
	}
	
	path := GetPath(maze, start, dificulty)

	MarkSolution(maze, path)

	PrintMaze(maze, showSolution)

	os.Exit(0)
}