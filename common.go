package main

import (
	"fmt"
)

const empty rune = ' '
const wall rune = '█'
const stepSquare rune = '.'
const startSquare = 'S'
const goalSquare = 'X'

type Maze struct {
	Grid [][]rune
	width int
	height int
}

func newMaze(x int, y int) *Maze {
	maze := Maze {width: x, height: y}

	maze.Grid = makeNew(maze.width, maze.height, wall)

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

func PrintMaze(maze *Maze, showSolution bool) {
	for i := range (maze.height + 2) {
		fmt.Printf("%c", wall)
		i++
	}
	fmt.Printf("\n")
	for i := range maze.width {
		fmt.Printf("%c", wall)
		for j := range maze.height {
			if maze.Grid[i][j] == startSquare {
				fmt.Printf("\033[1m%c\033[0m", startSquare)
			} else if maze.Grid[i][j] == goalSquare {
				fmt.Printf("\033[1m%c\033[0m", goalSquare)
			} else if maze.Grid[i][j] == stepSquare {
				if showSolution {
					fmt.Printf("%c", stepSquare)
				} else {
					fmt.Printf("%c", empty)
				}
			} else {
				fmt.Printf("%c", maze.Grid[i][j])
			}
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

func MarkSolution(maze *Maze, path Stack[[2]int]) {
	start := path.Pop()

	maze.Grid[start[0]][start[1]] = startSquare

	for path.Size() > 1 {
		
		step := path.Pop()

		maze.Grid[step[0]][step[1]] = stepSquare
	}

	goal := path.Pop()

	maze.Grid[goal[0]][goal[1]] = goalSquare
}

func PrintUsage(binName string) {
	fmt.Println("Usage:")
	fmt.Println(binName, "[height] [width] [algorithm] [dificulty] [-s]")
	fmt.Println("\nalgorithms:")
	fmt.Println("1: Prim's (default)")
	fmt.Println("2: DFS")
	fmt.Println("\ndificulty: the higher the number the easier the puzzle. default = 0 (hardest)")
	fmt.Println("\n-s: Show Solution (optional)")
}