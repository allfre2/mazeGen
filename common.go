package main

import (
	"fmt"
)

const empty rune = ' '
const wall rune = '█' //'𝦦'

type Maze struct {
	Grid [][]rune
	width int
	height int
}

func newMaze(x int, y int) *Maze {
	maze := Maze {width: x, height: y}

	maze.Grid = make([][]rune, maze.width)

	for i := range maze.width {
		maze.Grid[i] = make([]rune, maze.height)
		for j := range maze.height {
			maze.Grid[i][j] = wall
		}
	}

	return &maze
}

func PrintMaze (maze *Maze) {
	fmt.Printf("\n")
	for i := range maze.width {
		for j := range maze.height {
			fmt.Printf("%c", maze.Grid[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func PrintUsage(binName string) {
	fmt.Println("Usage:")
	fmt.Println(binName, "[height] [width]")
}