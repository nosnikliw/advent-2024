/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"advent2024/cmd/day2"
	"fmt"

	"github.com/spf13/cobra"
)

// day2Cmd represents the day2 command
var day2Cmd = &cobra.Command{
	Use:   "day2",
	Short: "Solve day 2",
	Long: `Solve day 2
	
How many reports are safe?`,
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		sourceDataFile := "./data/day2/input.txt"
		if len(args) > 0 {
			sourceDataFile = args[0]
		}
		records := day2.LoadFile(sourceDataFile)

		safeCount, safeWithDampenerCount := day2.CountSafeRecords(records)
		fmt.Printf("Safe count: %d\n", safeCount)
		fmt.Printf("Safe with dampener count: %d\n", safeCount+safeWithDampenerCount)
	},
}

func init() {
	rootCmd.AddCommand(day2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
