package main

import (
	"math/rand/v2"
)

// Iterative randomized Prim's algorithm
func GenerateMaze(width int, height int) *Maze {
	maze := newMaze(width, height)

	var wallList [][2]int 
	wallCount := 0

	maze.Grid[0][0] = empty
	walls, count := GetWalls(maze, 0, 0)

	for i := range count {
		wallList = append(wallList, walls[i])
		wallCount++
	}

	for wallCount > 0 {
		randomWall := rand.IntN(wallCount)

		wall := wallList[randomWall]

		visitedCount := GetVisitedCount(maze, wall[0], wall[1])

		if visitedCount == 1 {
			maze.Grid[wall[0]][wall[1]] = empty

			adjWalls, adjCount := GetWalls(maze, wall[0], wall[1])

			for i := range adjCount {
				wallList = append(wallList, adjWalls[i])
				wallCount++
			}
		}

		wallList = append(wallList[:randomWall], wallList[randomWall+1:]...)
		wallCount--
	}

	return maze
}

func GetWalls(maze *Maze, x int, y int) ([][2]int, int) {
	
	var points [][2]int
	count := 0

	if x - 1 >= 0 && maze.Grid[x-1][y] == wall {
		points = append(points, [2]int {x-1, y})
		count++
	}

	if x + 1 < maze.width && maze.Grid[x+1][y] == wall {
		points = append(points, [2]int {x+1, y})
		count++
	}

	if y - 1 >= 0 && maze.Grid[x][y-1] == wall {
		points = append(points, [2]int {x, y-1})
		count++
	}

	if y + 1 < maze.height && maze.Grid[x][y+1] == wall {
		points = append(points, [2]int {x, y+1})
		count++
	}

	return points, count
}

func GetVisitedCount(maze *Maze, x int, y int) int {
	count := 0

	if x - 1 >= 0 {
		if maze.Grid[x-1][y] == empty {
			count++
		}
	}

	if x + 1 < maze.width {
		if maze.Grid[x+1][y] == empty {
			count++
		}
	}

	if y - 1 >= 0 {
		if maze.Grid[x][y-1] == empty {
			count++
		}
	}

	if y + 1 < maze.height {
		if maze.Grid[x][y+1] == empty {
			count++
		}
	}

	return count
}