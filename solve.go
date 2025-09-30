package main

import (
	"math"
)

// TODO:
// A* to find solutions of the maze
func Solve(maze *Maze, start [2]int, goal [2]int) Stack[[2]int] {
	var frontier PriorityQueue[[2]int]
	cameFrom := makeNew(maze.width, maze.height, [2]int {-1, -1})
	cost := makeNew(maze.width, maze.height, 0.0)
	
	frontier.Queue(start, 0.0)
	
	for frontier.Any() {
		current, _ := frontier.Dequeue()
		
		if cellEquals(current, goal) {
			break
		}
		
		neighbors, _ := GetAdjacent(maze, current)
		
		for _, next := range neighbors {
			newCost := cost[current[0]][current[1]] + 1
			
			if cost[next[0]][next[1]] == 0.0 || newCost < cost[next[0]][next[1]] {
				cost[next[0]][next[1]] = newCost
				priority := newCost + Distance(next, goal)
				frontier.Queue(next, priority)
				cameFrom[next[0]][next[1]] = current
			}
		}
	}
	
	// Reconstruct path
	var path Stack[[2]int]

	current := goal

	for current[0] != start[0] || current[1] != start[1] {
		path.Push(current)
		current = cameFrom[current[0]][current[1]]
	}

	path.Push(start)

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

func cellEquals(a [2]int, b [2]int) bool {
	return a[0] == b[0] && a[1] == b[1]
}