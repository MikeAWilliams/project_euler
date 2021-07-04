package main

import "fmt"

type coordinate struct {
	x, y int
}

func (c coordinate) moveRight() coordinate {
	return coordinate{x: c.x + 1, y: c.y}
}

func (c coordinate) moveDown() coordinate {
	return coordinate{x: c.x, y: c.y + 1}
}

func (c coordinate) equal(other coordinate) bool {
	return c.x == other.x && c.y == other.y
}

func countMovesFromPosition(current coordinate, max coordinate, accumulator int) int {
	if max.equal(current) {
		return accumulator + 1
	}
	if current.x < max.x {
		accumulator = countMovesFromPosition(current.moveRight(), max, accumulator)
	}
	if current.y < max.y {
		accumulator = countMovesFromPosition(current.moveDown(), max, accumulator)
	}
	return accumulator
}

func countMoves(max coordinate) int {
	result := 0
	current := coordinate{x: 0, y: 0}
	pivots := []coordinate{current}
	for len(pivots) > 0 {
		current, pivots = pivots[len(pivots)-1], pivots[0:len(pivots)-1]
		if max.x > current.x {
			next := current.moveRight()
			if max.equal(next) {
				result += 1
			} else {
				pivots = append(pivots, next)
			}
		}
		if max.y > current.y {
			next := current.moveDown()
			if max.equal(next) {
				result += 1
			} else {
				pivots = append(pivots, next)
			}
		}
	}

	return result
}

func countMovesFromEnd(size int) int64 {
	grid := [][]int64{}
	for j := 0; j <= size; j++ {
		grid = append(grid, make([]int64, size+1))
	}

	for i := 0; i < size; i++ {
		grid[i][size] = 1
		grid[size][i] = 1
	}
	for i := size - 1; i >= 0; i-- {
		for j := size - 1; j >= 0; j-- {
			grid[i][j] = grid[i+1][j] + grid[i][j+1]
		}
	}
	return grid[0][0]
}

func main() {
	fmt.Println(countMoves(coordinate{x: 2, y: 2}))
	fmt.Println(countMovesFromEnd(2))
	fmt.Println(countMoves(coordinate{x: 10, y: 10}))
	fmt.Println(countMovesFromEnd(10))
	fmt.Println(countMovesFromEnd(20))
}
