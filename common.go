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

func makeNew[T any](x int, y int, val T) [][]T {
	arr := make([][]T, x)
	for i := range x {
		arr[i] = make([]T, y)
		for j := range y {
			arr[i][j] = val
		}
	}

	return arr
}

func PrintMaze (maze *Maze) {
	for i := range (maze.height + 2) {
		fmt.Printf("%c", wall)
		i++
	}
	fmt.Printf("\n")
	for i := range maze.width {
		fmt.Printf("%c", wall)
		for j := range maze.height {
			fmt.Printf("%c", maze.Grid[i][j])
		}
		fmt.Printf("%c", wall)
		fmt.Printf("\n")
	}
	for i := range (maze.height + 2) {
		fmt.Printf("%c", wall)
		i++
	}
	fmt.Printf("\n")
}

func PrintUsage(binName string) {
	fmt.Println("Usage:")
	fmt.Println(binName, "[height] [width] [algorithm]")
	fmt.Println("\nalgorithms:")
	fmt.Println("1: Prim's")
	fmt.Println("2: DFS")
}