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

func countMovesFromPosition(current coordinate, max coordinate, accumulator int) int {
	if max.x == current.x && max.y == current.y {
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

func main() {
	fmt.Println(countMovesFromPosition(coordinate{x: 0, y: 0}, coordinate{x: 2, y: 2}, 0))

}
