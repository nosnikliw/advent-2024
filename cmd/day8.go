/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"advent2024/cmd/geometry"
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

		boundary := geometry.VectorBox{
			V1: geometry.Vector{X: 0, Y: 0},
			V2: geometry.Vector{X: (len(grid) - 1), Y: (len(grid[0]) - 1)},
		}
		lines := []geometry.Line{}
		antinodes := make([][]int, len(grid))
		for i := range antinodes {
			antinodes[i] = make([]int, len(grid[i]))
		}
		for _, network := range nodeLists {
			for i := 0; (i - 1) < len(network); i++ {
				for j := i + 1; j < len(network); j++ {
					ans := getAntinodes(network[i], network[j])
					lines = append(lines, geometry.Line{
						P1: network[i],
						P2: network[j],
					})
					for _, v := range ans {
						if boundary.Contains(v) {
							antinodes[v.X][v.Y] = 1
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
					v := geometry.Vector{X: i, Y: j}
					if v.Equals(l.P1) || v.Equals(l.P2) {
						count2++
						break
					}
					v1 := l.P1.To(l.P2)
					v2 := l.P1.To(geometry.Vector{X: i, Y: j})
					if v1.IsParallelTo(v2) {
						count2++
						break
					}
				}
			}
		}

		fmt.Printf("Antinode count 2: %d\n", count2)
	},
}

func getAntinodes(a geometry.Vector, b geometry.Vector) (antinodes []geometry.Vector) {
	atob := a.To(b)
	antinodes = append(antinodes, a.Subtract(atob))
	antinodes = append(antinodes, b.Add(atob))
	return
}

func buildNodeLists(grid [][]string) (result map[string][]geometry.Vector) {
	result = map[string][]geometry.Vector{}
	for i := range grid {
		for j, v := range grid[i] {
			if v != "." {
				result[v] = append(result[v], geometry.Vector{X: i, Y: j})
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
