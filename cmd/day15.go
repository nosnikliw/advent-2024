/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"advent2024/cmd/day15"
	"fmt"

	"github.com/spf13/cobra"
)

// day15Cmd represents the day15 command
var day15Cmd = &cobra.Command{
	Use:   "day15",
	Short: "Solve day 15",
	Long: `Solve day 15

Warehouse Woes`,
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		sourceDataFile := "./data/day15/input.txt"
		if len(args) > 0 {
			sourceDataFile = args[0]
		}

		warehouse, moves, position := day15.LoadFile(sourceDataFile)
		wide, position2 := day15.MakeWide(warehouse)
		for _, move := range moves {
			position = day15.Move(&warehouse, move, position)
		}

		coordinates := 0
		for i, line := range warehouse {
			for j, v := range line {
				fmt.Print(v)
				if v == "O" {
					coordinates += (100 * i) + j
				}
			}
			fmt.Println("")
		}

		fmt.Println("GPS coordinates:", coordinates)

		for _, move := range moves {
			position2 = day15.MoveWide(&wide, move, position2)
		}

		wideCoordinates := 0
		for i, line := range wide {
			for j, v := range line {
				fmt.Print(v)
				if v == "[" {
					wideCoordinates += (100 * i) + j
				}
			}
			fmt.Println("")
		}

		fmt.Println("Wide GPS coordinates:", wideCoordinates)
	},
}

func init() {
	rootCmd.AddCommand(day15Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day15Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day15Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
