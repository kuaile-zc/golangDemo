package main

import (
	"fmt"
)

func getMaze() [][]int {
	return [][]int{
		{0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0},
		{0, 1, 0, 1, 0},
		{1, 1, 1, 0, 0},
		{0, 1, 0, 0, 1},
		{0, 1, 0, 0, 0},
	}

}

type point struct {
	i int
	j int
}

func (p point) addPoint(i, j, row, lin int) (point, bool) {
	if p.i+i < 0 || p.i+i >= row || p.j+j < 0 || p.j+j >= lin {
		return point{0, 0}, false
	}
	return point{p.i + i, p.j + j}, true
}

var dirs = [4]point{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
}

var startPoint = point{0, 0}

func walk(maze [][]int) int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	stack := []point{{0, 0}}

	for len(stack) > 0 {
		print(steps)
		pointVal := stack[0]
		stack = stack[1:]

		for _, p := range dirs {
			if addPoint, b := pointVal.addPoint(p.i, p.j, len(maze), len(maze[0]));
				b &&
					steps[addPoint.i][addPoint.j] == 0 &&
					maze[addPoint.i][addPoint.j] == 0 &&
					addPoint != startPoint {
				steps[addPoint.i][addPoint.j] = steps[pointVal.i][pointVal.j] + 1
				stack = append(stack, addPoint)
			}
		}
	}

	return 0
}

func print(print [][]int) {
	for _, row := range print {
		for _, value := range row {
			fmt.Printf("%3d", value)
		}
		fmt.Println()
	}
	fmt.Println("==========================================")
}

func main() {
	maze := getMaze()
	print(maze)

	walk(maze)
}
