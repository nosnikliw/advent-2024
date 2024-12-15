/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"advent2024/cmd/day13"
	"fmt"

	"github.com/spf13/cobra"
)

// day13Cmd represents the day13 command
var day13Cmd = &cobra.Command{
	Use:   "day13",
	Short: "Solve day 13",
	Long: `Solve day 13

Claw Contraption`,
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		sourceDataFile := "./data/day13/input.txt"
		if len(args) > 0 {
			sourceDataFile = args[0]
		}
		machines, corrected := day13.LoadFile(sourceDataFile)
		fmt.Printf("Tokens: %d\n", countTokens(machines))
		fmt.Printf("Corrected tokens: %d\n", countTokens(corrected))
	},
}

func countTokens(machines []day13.Machine) (count int) {
	for _, machine := range machines {
		a, err := machine.CountA()
		if err != nil {
			continue
		}
		b, err := machine.CountB()
		if err != nil {
			continue
		}
		count += 3*a + b
	}
	return
}

func init() {
	rootCmd.AddCommand(day13Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day13Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day13Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
