/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"advent2024/cmd/day16"
	"fmt"

	"github.com/spf13/cobra"
)

// day16Cmd represents the day16 command
var day16Cmd = &cobra.Command{
	Use:   "day16",
	Short: "Solve day 16",
	Long: `Solve day 16

Reindeer Maze`,
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		sourceDataFile := "./data/day16/input.txt"
		if len(args) > 0 {
			sourceDataFile = args[0]
		}

		maze := day16.LoadFile(sourceDataFile)

		best, count := day16.RunMaze(maze)

		fmt.Println("Best:", best)
		fmt.Println("Count:", count)
	},
}

func init() {
	rootCmd.AddCommand(day16Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day16Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day16Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
