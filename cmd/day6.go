/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// day6Cmd represents the day6 command
var day6Cmd = &cobra.Command{
	Use:   "day6",
	Short: "Solve day 6",
	Long: `Solve day 6

Guard patrol route`,
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		sourceDataFile := "./data/day6/input.txt"
		if len(args) > 0 {
			sourceDataFile = args[0]
		}

		area := loadDay6Input(sourceDataFile)

		patrolled, _, err := patrol(area)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Patrolled locations: %d\n", countVisited(patrolled))

		possibleLoops := 0
		for i := range area {
			for j := range area[i] {
				if area[i][j] == "." {
					_, loop, _ := patrol(area, Position{x: j, y: i})
					if loop {
						possibleLoops++
					}
				}
			}
		}

		fmt.Printf("PossibleLoops: %d\n", possibleLoops)

	},
}

func countVisited(area [][]int) int {
	result := 0
	for _, line := range area {
		for _, pos := range line {
			result += pos
		}
	}
	return result
}

type Direction struct {
	dx int
	dy int
}

type Position struct {
	x int
	y int
}

var North Direction = Direction{dx: 0, dy: -1}
var East Direction = Direction{dx: 1, dy: 0}
var South Direction = Direction{dx: 0, dy: 1}
var West Direction = Direction{dx: -1, dy: 0}

func sameDirection(a Direction, b Direction) bool {
	if a.dx == b.dx && a.dy == b.dy {
		return true
	}
	return false
}

func turnRight(direction Direction) (Direction, error) {
	if sameDirection(direction, North) {
		return East, nil
	}
	if sameDirection(direction, East) {
		return South, nil
	}
	if sameDirection(direction, South) {
		return West, nil
	}
	if sameDirection(direction, West) {
		return North, nil
	}
	return direction, errors.New("input direction invalid")
}

func patrol(area [][]string, blocks ...Position) ([][]int, bool, error) {
	visited := make([][]int, len(area))
	for i := range visited {
		visited[i] = make([]int, len(area[i]))
	}
	directions := make([][][]Direction, len(area))
	for i := range directions {
		directions[i] = make([][]Direction, len(area[i]))
	}

	guardLoc, err := findGuard(area)
	if err != nil {
		return nil, false, err
	}

	direction := North

	for inArea(area, guardLoc) {
		visited[guardLoc.y][guardLoc.x] = 1
		nextLoc := move(guardLoc, direction)
		if inArea(area, nextLoc) && directionInList(directions[nextLoc.y][nextLoc.x], direction) {
			return visited, true, nil
		}
		if blocked(area, nextLoc) || inList(blocks, nextLoc) {
			direction, _ = turnRight(direction)
		} else {
			guardLoc = nextLoc
			if inArea(area, guardLoc) {
				directions[guardLoc.y][guardLoc.x] = append(directions[guardLoc.y][guardLoc.x], direction)
			}
		}
	}
	return visited, false, nil
}

func directionInList(list []Direction, val Direction) bool {
	for _, d := range list {
		if sameDirection(d, val) {
			return true
		}
	}
	return false
}

func inList(list []Position, val Position) bool {
	for _, p := range list {
		if p.x == val.x && p.y == val.y {
			return true
		}
	}
	return false
}
func blocked(area [][]string, pos Position) bool {
	if inArea(area, pos) {
		if area[pos.y][pos.x] == "#" {
			return true
		}
	}
	return false
}

func move(pos Position, dir Direction) Position {
	return Position{
		x: pos.x + dir.dx,
		y: pos.y + dir.dy,
	}
}

func inArea(area [][]string, pos Position) bool {
	if pos.x < 0 || pos.y < 0 || pos.x >= len(area[0]) || pos.y >= len(area) {
		return false
	}
	return true
}

func findGuard(area [][]string) (Position, error) {
	for i := 0; i < len(area); i++ {
		for j := 0; j < len(area[i]); j++ {
			if area[i][j] == "^" {
				return Position{y: i, x: j}, nil
			}
		}
	}
	return Position{}, errors.New("no guard in area")
}

func loadDay6Input(fileName string) [][]string {
	readFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	area := [][]string{}
	for fileScanner.Scan() {
		area = append(area, strings.Split(strings.TrimSpace(fileScanner.Text()), ""))
	}
	return area
}

func init() {
	rootCmd.AddCommand(day6Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day6Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day6Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
