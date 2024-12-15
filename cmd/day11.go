/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"advent2024/cmd/day11"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// day11Cmd represents the day11 command
var day11Cmd = &cobra.Command{
	Use:   "day11",
	Short: "Solve day 11",
	Long: `Solve day 11

Plutonian Pebbles`,
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		sourceDataFile := "./data/day11/input.txt"
		if len(args) > 0 {
			sourceDataFile = args[0]
		}

		stones := day11.LoadFile(sourceDataFile)
		fmt.Println(stones)
		count, err := cmd.Flags().GetInt("blink-count")
		if err != nil {
			fmt.Println(err)
		}

		stoneCount := 0
		counted := map[day11.ValLevel]int{}
		for _, stone := range stones {
			s, _ := strconv.Atoi(stone)
			stoneCount += day11.CountStones(s, count, &counted)
		}
		fmt.Printf("Number of stones after blinking %d times: %d\n", count, stoneCount)
	},
}

func init() {
	rootCmd.AddCommand(day11Cmd)

	day11Cmd.Flags().Int("blink-count", 5, "How many times to blink")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day11Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day11Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
