package day14

import (
	"advent2024/cmd/geometry"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Robot struct {
	Position geometry.Vector
	Velocity geometry.Vector
}

func (r Robot) NewPosition(seconds int, width int, height int) geometry.Vector {
	p := r.Position.Add(r.Velocity.Scale(seconds))
	return geometry.Vector{
		X: wrap(p.X, width),
		Y: wrap(p.Y, height),
	}
}

func wrap(val int, size int) int {
	mod := val % size
	if mod < 0 {
		return mod + size
	}
	return mod
}

func LoadFile(fileName string) (robots []Robot) {
	readFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fields := strings.Fields(fileScanner.Text())
		robots = append(robots, Robot{
			Position: readVector(fields[0]),
			Velocity: readVector(fields[1]),
		})
	}
	return
}

func readVector(s string) geometry.Vector {
	values := strings.Split(strings.Split(s, "=")[1], ",")
	x, _ := strconv.Atoi(values[0])
	y, _ := strconv.Atoi(values[1])
	return geometry.Vector{
		X: x,
		Y: y,
	}
}

func GetGrid(robots []Robot, width int, height int, time int) [][]int {
	grid := make([][]int, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]int, width)
	}
	for _, r := range robots {
		p := r.NewPosition(time, width, height)
		grid[p.Y][p.X]++
	}
	return grid
}
