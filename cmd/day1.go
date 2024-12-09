/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"advent2024/cmd/day1"
	"fmt"
	"sort"

	"github.com/spf13/cobra"
)

// day1Cmd represents the day1 command
var day1Cmd = &cobra.Command{
	Use:   "day1",
	Short: "Solve day 1",
	Long: `Solve day 1.
	
Find the distance between two lists`,
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		sourceDataFile := "./data/day1/input.txt"
		if len(args) > 0 {
			sourceDataFile = args[0]
		}

		list1, list2 := day1.LoadFile(sourceDataFile)

		sort.Ints(list1)
		sort.Ints(list2)

		distance := day1.CalculateDistance(list1, list2)
		fmt.Printf("Distance: %d\n", distance)

		similarity := day1.CalculateSimilarity(list1, list2)
		fmt.Printf("Similarity: %d\n", similarity)
	},
}

func init() {
	rootCmd.AddCommand(day1Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day1Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day1Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
