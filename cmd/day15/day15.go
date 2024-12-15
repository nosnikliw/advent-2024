package day15

import (
	"advent2024/cmd/geometry"
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

var directions = map[string]geometry.Vector{
	"^": {X: 0, Y: -1},
	">": {X: 1, Y: 0},
	"v": {X: 0, Y: 1},
	"<": {X: -1, Y: 0},
}

func LoadFile(fileName string) (warehouse [][]string, moves []geometry.Vector, position geometry.Vector) {
	readFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	loadedMap := false
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if !loadedMap {
			if line == "" {
				loadedMap = true
				continue
			}
			warehouse = append(warehouse, strings.Split(line, ""))
		} else {
			mvs := strings.Split(line, "")
			for _, m := range mvs {
				moves = append(moves, directions[m])
			}
		}
	}
	for i, row := range warehouse {
		for j, v := range row {
			if v == "@" {
				position.X = j
				position.Y = i
			}
		}
	}
	return
}

func MakeWide(warehouse [][]string) (wider [][]string, position geometry.Vector) {
	for _, row := range warehouse {
		r := []string{}
		for _, v := range row {
			if v == "O" {
				r = append(r, "[", "]")
			} else if v == "#" {
				r = append(r, "#", "#")
			} else {
				r = append(r, v, ".")
			}
		}
		wider = append(wider, r)
	}
	for i, row := range wider {
		for j, v := range row {
			if v == "@" {
				position.X = j
				position.Y = i
			}
		}
	}
	return
}

func Move(warehouse *[][]string, direction geometry.Vector, position geometry.Vector) (newPosition geometry.Vector) {
	if tryMove(warehouse, direction, position) {
		(*warehouse)[position.Y][position.X] = "."
		newPosition = position.Add(direction)
		(*warehouse)[newPosition.Y][newPosition.X] = "@"
	} else {
		newPosition = position
	}
	return
}

func MoveWide(warehouse *[][]string, direction geometry.Vector, position geometry.Vector) (newPosition geometry.Vector) {
	// fmt.Println(position, direction)
	if direction.Y == 0 {
		newPosition = Move(warehouse, direction, position)
		return
	}
	if canMoveVertically(warehouse, direction, position) {
		moveVertically(warehouse, direction, []geometry.Vector{position})
		(*warehouse)[position.Y][position.X] = "."
		newPosition = position.Add(direction)
		(*warehouse)[newPosition.Y][newPosition.X] = "@"
	} else {
		newPosition = position
	}
	return
}

func canMoveVertically(warehouse *[][]string, direction geometry.Vector, position geometry.Vector) bool {
	next := position.Add(direction)
	nextVal := (*warehouse)[next.Y][next.X]
	if nextVal == "." {
		return true
	}
	if nextVal == "#" {
		return false
	}
	if nextVal == "[" {
		pair := geometry.Vector{X: next.X + 1, Y: next.Y}
		return canMoveVertically(warehouse, direction, next) && canMoveVertically(warehouse, direction, pair)
	}
	if nextVal == "]" {
		pair := geometry.Vector{X: next.X - 1, Y: next.Y}
		return canMoveVertically(warehouse, direction, next) && canMoveVertically(warehouse, direction, pair)
	}
	panic(nextVal)
}

func moveVertically(warehouse *[][]string, direction geometry.Vector, positions []geometry.Vector) {
	nextRow := getNextRow(warehouse, direction, positions)

	if len(nextRow) > 0 {
		moveVertically(warehouse, direction, nextRow)
	}

	for _, p := range positions {
		newPos := p.Add(direction)
		val := (*warehouse)[p.Y][p.X]
		(*warehouse)[newPos.Y][newPos.X] = val
	}

	for _, p := range nextRow {
		old := p.Subtract(direction)
		if !slices.Contains(positions, old) {
			(*warehouse)[p.Y][p.X] = "."
		}
	}

	//moveVertically(warehouse, direction, nextRow)

	// nextVal := (*warehouse)[next.Y][next.X]
	// currentVal := (*warehouse)[position.Y][position.X]
	// if nextVal == "." {
	// 	(*warehouse)[next.Y][next.X] = currentVal
	// }
	// if nextVal == "#" {
	// 	panic("invalid move")
	// }
	// if nextVal == "[" {
	// 	pair := geometry.Vector{X: next.X + 1, Y: next.Y}
	// 	moveVertically(warehouse, direction, next)
	// 	moveVertically(warehouse, direction, pair)
	// }
	// if nextVal == "]" {
	// 	pair := geometry.Vector{X: next.X - 1, Y: next.Y}
	// 	moveVertically(warehouse, direction, next)
	// 	moveVertically(warehouse, direction, pair)
	// }
}

func getNextRow(warehouse *[][]string, direction geometry.Vector, positions []geometry.Vector) (nextRow []geometry.Vector) {
	for _, p := range positions {
		next := p.Add(direction)
		nextVal := (*warehouse)[next.Y][next.X]
		// fmt.Println(nextVal, next)
		if nextVal == "#" {
			panic("invalid move")
		}
		if nextVal == "[" {
			nextRow = append(nextRow, next)
			nextRow = append(nextRow, next.Add(directions[">"]))
		}
		if nextVal == "]" {
			nextRow = append(nextRow, next)
			nextRow = append(nextRow, next.Add(directions["<"]))
		}
	}
	// fmt.Println(nextRow)
	nextRow = distinct(nextRow)
	// fmt.Println(nextRow)
	return
}

func distinct[T comparable](list []T) []T {
	dist := map[T]bool{}
	for _, v := range list {
		dist[v] = true
	}
	keys := make([]T, 0, len(dist))
	for k := range dist {
		keys = append(keys, k)
	}
	return keys
}

func tryMove(warehouse *[][]string, direction geometry.Vector, position geometry.Vector) bool {
	next := position.Add(direction)
	nextVal := (*warehouse)[next.Y][next.X]
	currentVal := (*warehouse)[position.Y][position.X]
	if nextVal == "." {
		(*warehouse)[next.Y][next.X] = currentVal
		return true
	}
	if nextVal == "#" {
		return false
	}
	if tryMove(warehouse, direction, next) {
		(*warehouse)[next.Y][next.X] = currentVal
		return true
	}
	return false
}
