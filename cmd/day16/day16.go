package day16

import (
	"advent2024/cmd/geometry"
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

func LoadFile(fileName string) (maze [][]string) {
	readFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		maze = append(maze, strings.Split(fileScanner.Text(), ""))
	}
	return
}

var North geometry.Vector = geometry.Vector{X: 0, Y: -1}
var East geometry.Vector = geometry.Vector{X: 1, Y: 0}
var South geometry.Vector = geometry.Vector{X: 0, Y: 1}
var West geometry.Vector = geometry.Vector{X: -1, Y: 0}

func RunMaze(maze [][]string) (int, int) {
	start := []geometry.Vector{findStart(maze)}
	end := findEnd(maze)
	progress := make([][]int, len(maze))
	for i := 0; i < len(progress); i++ {
		progress[i] = make([]int, len(maze[i]))
	}

	paths := []Path{}

	progress[end.Y][end.X] = math.MaxInt

	traverse(maze, &progress, start, East, 0, &paths, end)
	traverse(maze, &progress, start, North, 1000, &paths, end)
	traverse(maze, &progress, start, South, 1000, &paths, end)
	traverse(maze, &progress, start, West, 2000, &paths, end)

	bestScore := progress[end.Y][end.X]

	bestPaths := filter(paths, func(p Path) bool { return p.Score == bestScore })

	positions := []geometry.Vector{}
	for _, p := range bestPaths {
		positions = append(positions, p.Path...)
	}
	posCount := len(distinct(positions))

	for i, line := range maze {
		for j, _ := range line {
			if slices.Contains(positions, geometry.Vector{X: j, Y: i}) {
				fmt.Print("O")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}

	return bestScore, posCount
}

func filter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
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

type Path struct {
	Path  []geometry.Vector
	Score int
}

func traverse(maze [][]string, progress *[][]int, path []geometry.Vector, currentDirection geometry.Vector, score int, completePaths *[]Path, end geometry.Vector) {
	newPos := path[len(path)-1].Add(currentDirection)
	newScore := score + 1
	if score > 65436 {
		return
	}
	if isBoundary(maze, newPos) {
		return
	}
	if slices.Contains(path, newPos) {
		return
	}

	if !fastestToHere(progress, newPos, newScore) {
		if isBoundary(maze, newPos.Add(currentDirection)) || !fastestToHere(progress, newPos.Add(currentDirection), newScore+1) {
			return
		}
	}
	setScore(progress, newPos, newScore)
	newPath := make([]geometry.Vector, len(path))
	_ = copy(newPath, path)
	newPath = append(newPath, newPos)

	if isEnd(maze, newPos) {
		(*completePaths) = append((*completePaths), Path{
			Path:  newPath,
			Score: newScore,
		})
		return
	}
	traverse(maze, progress, newPath, currentDirection, newScore, completePaths, end)
	if currentDirection.X == 0 {
		traverse(maze, progress, newPath, East, newScore+1000, completePaths, end)
		traverse(maze, progress, newPath, West, newScore+1000, completePaths, end)
	} else {
		traverse(maze, progress, newPath, North, newScore+1000, completePaths, end)
		traverse(maze, progress, newPath, South, newScore+1000, completePaths, end)
	}
}

func displayProgress(progress *[][]int) {
	for _, v := range *progress {
		fmt.Println(v)
	}
}

func setScore(progress *[][]int, pos geometry.Vector, score int) {
	if (*progress)[pos.Y][pos.X] == 0 {
		(*progress)[pos.Y][pos.X] = score
	} else if (*progress)[pos.Y][pos.X] > score {
		(*progress)[pos.Y][pos.X] = score
	}
}

func isEnd(maze [][]string, pos geometry.Vector) bool {
	return maze[pos.Y][pos.X] == "E"
}

func isBoundary(maze [][]string, pos geometry.Vector) bool {
	return maze[pos.Y][pos.X] == "#"
}

func fastestToHere(progress *[][]int, pos geometry.Vector, score int) bool {
	val := (*progress)[pos.Y][pos.X]
	if val == 0 || val >= score {
		return true
	}
	return false
}

func findStart(maze [][]string) geometry.Vector {
	for i, row := range maze {
		for j, cell := range row {
			if cell == "S" {
				return geometry.Vector{
					X: j,
					Y: i,
				}
			}
		}
	}
	panic("no start in maze")
}

func findEnd(maze [][]string) geometry.Vector {
	for i, row := range maze {
		for j, cell := range row {
			if cell == "E" {
				return geometry.Vector{
					X: j,
					Y: i,
				}
			}
		}
	}
	panic("no end in maze")
}
