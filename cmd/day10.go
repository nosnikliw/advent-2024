/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"advent2024/cmd/day10"
	"fmt"

	"github.com/spf13/cobra"
)

// day10Cmd represents the day10 command
var day10Cmd = &cobra.Command{
	Use:   "day10",
	Short: "Solve day 10",
	Long: `Solve day 10

Hoof it`,
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		sourceDataFile := "./data/day10/input.txt"
		if len(args) > 0 {
			sourceDataFile = args[0]
		}

		area := day10.LoadFile(sourceDataFile)

		trailCount, rating := day10.GetTotalTrailCount(area, 9)

		fmt.Printf("Trail count: %d\n", trailCount)
		fmt.Printf("Rating     : %d\n", rating)
	},
}

func init() {
	rootCmd.AddCommand(day10Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day10Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day10Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
