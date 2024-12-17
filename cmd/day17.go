/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"advent2024/cmd/day17"
	"fmt"

	"github.com/spf13/cobra"
)

// day17Cmd represents the day17 command
var day17Cmd = &cobra.Command{
	Use:   "day17",
	Short: "Solve day 17",
	Long: `Solve day 17

Chronospatial Computer`,
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		sourceDataFile := "./data/day17/input.txt"
		if len(args) > 0 {
			sourceDataFile = args[0]
		}

		machine := day17.LoadFile(sourceDataFile)

		machine.DisplayProgram()

		output := machine.Run()

		fmt.Print("Output: ")
		for _, v := range output {
			fmt.Print(v)
			fmt.Print(",")
		}
		fmt.Println("")
		last := 1
		i := 0
		for j := len(machine.Program) - 1; j >= 0; j-- {
			for ; true; i++ {
				machine.A = i
				machine.B = 0
				machine.C = 0
				machine.Output = []int{}
				if machine.Expect(machine.Program[j:]) {
					fmt.Println("Min:", i, i/last)
					last = i
					i = i * 8
					break
				}
				// if i%1000000 == 0 {
				// 	fmt.Println(i)
				// }
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(day17Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day17Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day17Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
