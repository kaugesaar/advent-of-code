package pathfinding

import (
	"github.com/kaugesaar/advent-of-code/utils/matrix"
)

var directions = [][]int{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
}

// Bfs find the shortage way
func Bfs(m matrix.IntMatrix, start matrix.Pos, target matrix.Pos) int {
	rows := len(m)
	cols := len(m[0])
	queue := []matrix.Pos{start}
	dist := make([][]int, rows)

	for i := 0; i < rows; i++ {
		dist[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			dist[i][j] = -1
		}
	}

	dist[start.Y][start.X] = 0

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr == target {
			return dist[target.Y][target.X]
		}

		for _, direction := range directions {
			next := matrix.Pos{Y: direction[0] + curr.Y, X: direction[1] + curr.X}

			if next.Y < 0 || next.Y == rows || next.X < 0 ||
				next.X == cols || dist[next.Y][next.X] != -1 ||
				m[next.Y][next.X]-m[curr.Y][curr.X] > 1 {
				continue
			}

			dist[next.Y][next.X] = dist[curr.Y][curr.X] + 1
			queue = append(queue, next)
		}
	}

	return -1
}
