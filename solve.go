package main

import (
	"math"
)

// TODO:
// A* to find solutions to the maze
func Solve(maze *Maze, start [2]int, goal [2]int) Stack[[2]int] {
	var path Stack[[2]int]

	return path
}

func GetAllPaths(maze *Maze, start [2]int) PriorityQueue[Stack[[2]int]] {

	var paths PriorityQueue[Stack[[2]int]]

	cells := GetSortedCells(maze, start)

	for cells.Any() {
		goal, distance := cells.Dequeue()
		path := Solve(maze, start, goal)
		paths.Queue(path, distance)
	}

	return paths
}

func Distance(a [2]int, b [2]int) float64 {
	return math.Sqrt(math.Pow(math.Abs(float64(b[0] - a[0])), 2) + math.Pow(math.Abs(float64(b[1] - a[0])), 2))
}

func GetSortedCells(maze *Maze, start [2]int) PriorityQueue[[2]int] {

	var queue PriorityQueue[[2]int]

	for i := range maze.width {
		for j := range maze.height {
			if maze.Grid[i][j] == empty {
				queue.Queue([2]int {i, j}, Distance(start, [2]int {i, j}))
			}
		}
	}

	return queue
}

