/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// day8Cmd represents the day8 command
var day8Cmd = &cobra.Command{
	Use:   "day8",
	Short: "Solve day 8",
	Long: `Solve day 8

Resonant Collinearity`,
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		sourceDataFile := "./data/day8/input.txt"
		if len(args) > 0 {
			sourceDataFile = args[0]
		}

		grid := loadDay8Input(sourceDataFile)

		nodeLists := buildNodeLists(grid)

		boundary := VectorSquare{
			v1: Vector{0, 0},
			v2: Vector{x: (len(grid) - 1), y: (len(grid[0]) - 1)},
		}
		lines := []Line{}
		antinodes := make([][]int, len(grid))
		for i := range antinodes {
			antinodes[i] = make([]int, len(grid[i]))
		}
		for _, network := range nodeLists {
			for i := 0; (i - 1) < len(network); i++ {
				for j := i + 1; j < len(network); j++ {
					ans := getAntinodes(network[i], network[j])
					lines = append(lines, Line{
						p1: network[i],
						p2: network[j],
					})
					for _, v := range ans {
						if boundary.contains(v) {
							antinodes[v.x][v.y] = 1
						}
					}
				}
			}
		}

		count := 0
		for _, line := range antinodes {
			for _, v := range line {
				count += v
			}
		}

		fmt.Printf("Antinode count: %d\n", count)

		count2 := 0
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				for _, l := range lines {
					v := Vector{i, j}
					if v.equals(l.p1) || v.equals(l.p2) {
						count2++
						break
					}
					v1 := l.p1.vectorTo(l.p2)
					v2 := l.p1.vectorTo(Vector{i, j})
					if v1.isParallelTo(v2) {
						count2++
						break
					}
				}
			}
		}

		fmt.Printf("Antinode count 2: %d\n", count2)
	},
}

func getAntinodes(a Vector, b Vector) (antinodes []Vector) {
	atob := a.vectorTo(b)
	antinodes = append(antinodes, a.subtract(atob))
	antinodes = append(antinodes, b.add(atob))
	return
}

type Line struct {
	p1 Vector
	p2 Vector
}

type VectorSquare struct {
	v1 Vector
	v2 Vector
}

func (s VectorSquare) contains(v Vector) bool {
	return ((s.v1.x-v.x)*(s.v2.x-v.x) <= 0) && ((s.v1.y-v.y)*(s.v2.y-v.y) <= 0)
}

type Vector struct {
	x int
	y int
}

func (dv1 Vector) isParallelTo(dv2 Vector) bool {
	return float64(dv1.x)/float64(dv1.y) == float64(dv2.x)/float64(dv2.y)
}

func (from Vector) vectorTo(to Vector) Vector {
	return to.subtract(from)
}

func (a Vector) add(b Vector) Vector {
	return Vector{
		x: a.x + b.x,
		y: a.y + b.y,
	}
}

func (a Vector) subtract(b Vector) Vector {
	return Vector{
		x: a.x - b.x,
		y: a.y - b.y,
	}
}

func (a Vector) equals(b Vector) bool {
	return a.x == b.x && a.y == b.y
}

func buildNodeLists(grid [][]string) (result map[string][]Vector) {
	result = map[string][]Vector{}
	for i, _ := range grid {
		for j, v := range grid[i] {
			if v != "." {
				result[v] = append(result[v], Vector{x: i, y: j})
			}
		}
	}
	return
}

func loadDay8Input(fileName string) [][]string {
	readFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	grid := [][]string{}
	for fileScanner.Scan() {
		grid = append(grid, strings.Split(fileScanner.Text(), ""))
	}
	return grid
}

func init() {
	rootCmd.AddCommand(day8Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day8Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day8Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
