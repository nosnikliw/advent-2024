/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"advent2024/cmd/day14"
	"advent2024/cmd/geometry"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// day14Cmd represents the day14 command
var day14Cmd = &cobra.Command{
	Use:   "day14",
	Short: "Solve day 14",
	Long: `Solve day 14

Restroom Redoubt`,
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		sourceDataFile := "./data/day14/input.txt"
		if len(args) > 0 {
			sourceDataFile = args[0]
		}

		width, _ := strconv.Atoi(cmd.Flag("width").Value.String())
		height, _ := strconv.Atoi(cmd.Flag("height").Value.String())

		quadrants := []geometry.VectorBox{
			{
				V1: geometry.Vector{X: 0, Y: 0},
				V2: geometry.Vector{X: (width / 2) - 1, Y: (height / 2) - 1},
			},
			{
				V1: geometry.Vector{X: 0, Y: (height / 2) + 1},
				V2: geometry.Vector{X: (width / 2) - 1, Y: height - 1},
			},
			{
				V1: geometry.Vector{X: (width / 2) + 1, Y: 0},
				V2: geometry.Vector{X: width - 1, Y: (height / 2) - 1},
			},
			{
				V1: geometry.Vector{X: (width / 2) + 1, Y: (height / 2) + 1},
				V2: geometry.Vector{X: width - 1, Y: height - 1},
			},
		}

		counts := []int{0, 0, 0, 0}

		robots := day14.LoadFile(sourceDataFile)

		for _, r := range robots {
			for i, q := range quadrants {
				if q.Contains(r.NewPosition(100, width, height)) {
					counts[i] += 1
				}
			}
		}
		f, err := os.Create("output.txt")
		if err != nil {
			return
		}

		seconds := 1
		initial := day14.GetGrid(robots, width, height, 0)
		current := day14.GetGrid(robots, width, height, 1)
		for gridsDiffer(initial, current) {
			if !looksLikeAChristmasTree(current) {
				seconds++
				current = day14.GetGrid(robots, width, height, seconds)
				continue
			}
			for _, line := range current {
				for _, v := range line {
					if v == 0 {
						f.WriteString(" ")
					} else {
						f.WriteString("X")
					}
				}
				f.WriteString("\n")
			}
			f.WriteString("Time: ")
			f.WriteString(strconv.Itoa(seconds))
			f.WriteString("\n__________________________________________________________________________________________________________________\n")
			seconds++
			current = day14.GetGrid(robots, width, height, seconds)
		}

		err = f.Close()
		if err != nil {
			fmt.Println(err)
			return
		}

		safetyFactor := 1
		for _, c := range counts {
			safetyFactor *= c
		}

		fmt.Printf("Safety Factor: %d\n", safetyFactor)
		fmt.Println("Check output.txt for candidate Easter eggs")
	},
}

func looksLikeAChristmasTree(a [][]int) bool {

	for _, line := range a {
		t := 0
		for _, v := range line {
			t += v
		}
		if t > 30 {
			return true
		}
	}
	return false
}

func gridsDiffer(a [][]int, b [][]int) bool {
	for i, line := range a {
		for j, cell := range line {
			if b[i][j] != cell {
				return true
			}
		}
	}
	return false
}

func init() {
	rootCmd.AddCommand(day14Cmd)

	day14Cmd.Flags().Int("width", 101, "The width of the patrol area")
	day14Cmd.Flags().Int("height", 103, "The height of the patrol area")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day14Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day14Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
