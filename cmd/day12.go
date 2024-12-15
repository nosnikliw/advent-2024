/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"advent2024/cmd/day12"
	"fmt"

	"github.com/spf13/cobra"
)

// day12Cmd represents the day12 command
var day12Cmd = &cobra.Command{
	Use:   "day12",
	Short: "Solve day 12",
	Long: `Solve day 12

Garden Groups`,
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		sourceDataFile := "./data/day12/input.txt"
		if len(args) > 0 {
			sourceDataFile = args[0]
		}
		plots := day12.LoadFile(sourceDataFile)

		regions := day12.DetermineRegions(&plots)

		cost := 0
		bulk := 0
		for _, region := range regions {
			//fmt.Printf("A region of %s plants with price %d * %d = %d\n", region.Crop, region.Area, region.Perimeter, region.Area*region.Perimeter)
			cost += region.Area * region.Perimeter
			bulk += region.Area * region.Sides
		}

		fmt.Printf("Fence cost: %d\n", cost)
		fmt.Printf("Bulk rate: %d\n", bulk)
	},
}

func init() {
	rootCmd.AddCommand(day12Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day12Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day12Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
