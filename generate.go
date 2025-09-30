package main

import (
	"math/rand/v2"
)

// Iterative randomized Prim's algorithm
func GenerateMaze1(width int, height int) (*Maze, [2]int) {

	maze := newMaze(width, height)

	start := [2]int {0, 0}

	var wallList [][2]int 
	wallCount := 0

	maze.Grid[0][0] = empty
	walls, count := GetWalls(maze, start)

	for i := range count {
		wallList = append(wallList, walls[i])
		wallCount++
	}

	for wallCount > 0 {
		randomWall := rand.IntN(wallCount)

		wall := wallList[randomWall]

		visitedCount := GetVisitedCount(maze, wall)

		if visitedCount == 1 {
			maze.Grid[wall[0]][wall[1]] = empty

			adjWalls, adjCount := GetWalls(maze, wall)

			for i := range adjCount {
				wallList = append(wallList, adjWalls[i])
				wallCount++
			}
		}

		wallList = append(wallList[:randomWall], wallList[randomWall+1:]...)
		wallCount--
	}

	return maze, start
}

// Iterative DFS using a stack
func GenerateMaze2(width int, height int) (*Maze, [2]int) {
	maze := newMaze(width, height)

	var stack Stack[[2]int]
	
	start := [2]int {width/2, height/2}

	maze.Grid[start[0]][start[1]] = empty

	stack.Push(start)

	for stack.Size() > 0 {
		current := stack.Pop()
		
		walls, count := GetWalls(maze, current)

		if count > 0 {
			cell := walls[rand.IntN(count)]

			visited := GetVisitedCount(maze, cell)

			if (IsEdge(maze, cell[0], cell[1]) && visited == 1) || visited < 3 {
				stack.Push(current)
				maze.Grid[cell[0]][cell[1]] = empty
				stack.Push(cell)
			}
		}
	}

	return maze, start
}

func GetNeighbours(maze *Maze, cell [2]int) [][2]int {
	var points [][2]int

	if cell[0] - 1 >= 0 {
		points = append(points, [2]int {cell[0]-1, cell[1]})
	}

	if cell[0] + 1 < maze.width {
		points = append(points, [2]int {cell[0]+1, cell[1]})
	}

	if cell[1] - 1 >= 0 {
		points = append(points, [2]int {cell[0], cell[1]-1})
	}

	if cell[1] + 1 < maze.height {
		points = append(points, [2]int {cell[0], cell[1]+1})
	}

	return points
}

func GetWalls(maze *Maze, cell [2]int) ([][2]int, int) {
	
	adj := GetNeighbours(maze, cell)
	var points [][2]int
	count := 0

	for i := range len(adj) {
		if maze.Grid[adj[i][0]][adj[i][1]] == wall {
			points = append(points, adj[i])
			count++
		}
	}

	return points, count
}

func GetAdjacent(maze *Maze, cell [2]int) ([][2]int, int) {
	
	adj := GetNeighbours(maze, cell)
	var points [][2]int
	count := 0

	for i := range len(adj) {
		if maze.Grid[adj[i][0]][adj[i][1]] == empty {
			points = append(points, adj[i])
			count++
		}
	}

	return points, count
}

func GetVisitedCount(maze *Maze, cell [2]int) int {

	adj := GetNeighbours(maze, cell)

	count := 0

	for i := range len(adj) {
		if maze.Grid[adj[i][0]][adj[i][1]] == empty {
			count++
		}
	}

	return count
}

func IsEdge(maze *Maze, x int, y int) bool {
	return x == 0 || y == 0 || x == maze.width - 1 || y == maze.height -1
}