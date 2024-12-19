/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"advent2024/cmd/day19"
	"fmt"

	"github.com/spf13/cobra"
)

// day19Cmd represents the day19 command
var day19Cmd = &cobra.Command{
	Use:   "day19",
	Short: "Solve day 19",
	Long: `Solve day 19

RAM Run`,
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		sourceDataFile := "./data/day19/input.txt"
		if len(args) > 0 {
			sourceDataFile = args[0]
		}

		towels, designs := day19.LoadFile(sourceDataFile)

		fmt.Println(towels)
		fmt.Println(designs)

		possible, ways := day19.FindPossible(towels, designs)

		fmt.Println("Possible count", len(possible))
		fmt.Println("Possible ways", ways)
	},
}

func init() {
	rootCmd.AddCommand(day19Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day19Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day19Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
