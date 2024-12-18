package day18

import (
	"advent2024/cmd/geometry"
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parseInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func LoadFile(fileName string, space *[][]bool, limit int) (positions []geometry.Vector) {
	readFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for i := 0; fileScanner.Scan(); i++ {
		vals := strings.Split(fileScanner.Text(), ",")
		positions = append(positions, geometry.Vector{
			X: parseInt(vals[0]),
			Y: parseInt(vals[1]),
		})
		if i < limit {
			(*space)[parseInt(vals[0])][parseInt(vals[1])] = true
		}
	}

	for _, row := range *space {
		for _, v := range row {
			if v {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
	return
}

func Escape(space [][]bool) int {
	start := []geometry.Vector{{X: 0, Y: 0}}
	end := geometry.Vector{X: len(space[0]) - 1, Y: len(space) - 1}
	progress := make([][]int, len(space))
	for i := 0; i < len(progress); i++ {
		progress[i] = make([]int, len(space[i]))
	}

	paths := []Path{}

	progress[end.Y][end.X] = math.MaxInt

	traverse(space, &progress, start, East, 0, &paths, end)
	//	traverse(space, &progress, start, North, 0, &paths, end)
	traverse(space, &progress, start, South, 0, &paths, end)
	//	traverse(space, &progress, start, West, 0, &paths, end)

	bestScore := progress[end.Y][end.X]

	// bestPaths := filter(paths, func(p Path) bool { return p.Score == bestScore })

	// positions := []geometry.Vector{}
	// for _, p := range bestPaths {
	// 	positions = append(positions, p.Path...)
	// }
	// posCount := len(distinct(positions))

	// for i, line := range space {
	// 	for j, _ := range line {
	// 		if slices.Contains(positions, geometry.Vector{X: j, Y: i}) {
	// 			fmt.Print("O")
	// 		} else {
	// 			fmt.Print(".")
	// 		}
	// 	}
	// 	fmt.Println("")
	// }

	return bestScore
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

var North geometry.Vector = geometry.Vector{X: 0, Y: -1}
var East geometry.Vector = geometry.Vector{X: 1, Y: 0}
var South geometry.Vector = geometry.Vector{X: 0, Y: 1}
var West geometry.Vector = geometry.Vector{X: -1, Y: 0}

func fastestToHere(progress *[][]int, pos geometry.Vector, score int) bool {
	val := (*progress)[pos.Y][pos.X]
	if val == 0 || val > score {
		return true
	}
	return false
}

func setScore(progress *[][]int, pos geometry.Vector, score int) {
	if (*progress)[pos.Y][pos.X] == 0 {
		(*progress)[pos.Y][pos.X] = score
	} else if (*progress)[pos.Y][pos.X] > score {
		(*progress)[pos.Y][pos.X] = score
	}
}

func traverse(space [][]bool, progress *[][]int, path []geometry.Vector, currentDirection geometry.Vector, score int, completePaths *[]Path, end geometry.Vector) {
	newPos := path[len(path)-1].Add(currentDirection)
	newScore := score + 1
	if !geometry.Box(space).Contains(newPos) {
		return
	}
	if space[newPos.Y][newPos.X] {
		return
	}
	if slices.Contains(path, newPos) {
		return
	}
	if newScore >= (*progress)[end.Y][end.X] {
		return
	}

	if !fastestToHere(progress, newPos, newScore) {
		return
	}
	setScore(progress, newPos, newScore)
	newPath := make([]geometry.Vector, len(path))
	_ = copy(newPath, path)
	newPath = append(newPath, newPos)

	if newPos == end {
		// fmt.Println("Score:", newScore)
		(*completePaths) = append((*completePaths), Path{
			Path:  newPath,
			Score: newScore,
		})
		return
	}
	traverse(space, progress, newPath, currentDirection, newScore, completePaths, end)
	if currentDirection.X == 0 {
		traverse(space, progress, newPath, East, newScore, completePaths, end)
		traverse(space, progress, newPath, West, newScore, completePaths, end)
	} else {
		traverse(space, progress, newPath, North, newScore, completePaths, end)
		traverse(space, progress, newPath, South, newScore, completePaths, end)
	}
}
