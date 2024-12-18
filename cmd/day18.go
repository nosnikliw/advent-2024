/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"advent2024/cmd/day18"
	"fmt"
	"math"
	"strconv"

	"github.com/spf13/cobra"
)

// day18Cmd represents the day18 command
var day18Cmd = &cobra.Command{
	Use:   "day18",
	Short: "Solve day 17",
	Long: `Solve day 17

RAM Run`,
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		sourceDataFile := "./data/day18/input.txt"
		if len(args) > 0 {
			sourceDataFile = args[0]
		}

		width, _ := strconv.Atoi(cmd.Flag("width").Value.String())
		height, _ := strconv.Atoi(cmd.Flag("height").Value.String())
		limit, _ := strconv.Atoi(cmd.Flag("limit").Value.String())

		space := make([][]bool, height)
		for i, _ := range space {
			space[i] = make([]bool, width)
		}

		locations := day18.LoadFile(sourceDataFile, &space, limit)

		shortest := day18.Escape(space)

		fmt.Println("Shortest:", shortest)

		bottom := limit
		top := len(locations)
		for bottom < top {
			// fmt.Print("Range:", bottom, top)
			testPos := (bottom + top) / 2
			s := make([][]bool, height)
			for i, _ := range s {
				s[i] = make([]bool, width)
			}
			for i := 0; i < testPos; i++ {
				l := locations[i]
				s[l.X][l.Y] = true
			}
			// fmt.Print(" Added to:", locations[testPos-1])
			len := day18.Escape(s)
			// fmt.Println(" Shortest:", len)
			if testPos == bottom {
				break
			}
			if len < math.MaxInt {
				bottom = testPos
			} else {
				top = testPos
			}
		}
		fmt.Println(locations[top-1], "prevents escape")
	},
}

func init() {
	rootCmd.AddCommand(day18Cmd)

	day18Cmd.Flags().Int("width", 71, "The width of the memory area")
	day18Cmd.Flags().Int("height", 71, "The height of the memory area")
	day18Cmd.Flags().Int("limit", 1024, "The number of bytes to let fall")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day18Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day18Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
