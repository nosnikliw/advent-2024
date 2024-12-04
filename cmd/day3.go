/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/spf13/cobra"
)

// day3Cmd represents the day3 command
var day3Cmd = &cobra.Command{
	Use:   "day3",
	Short: "Solve day 3",
	Long: `Solve day 3

Uncorrupt the program`,
	Run: func(cmd *cobra.Command, args []string) {
		sourceDataFile := "./data/day3/input.txt"
		if len(args) > 0 {
			sourceDataFile = args[0]
		}
		buf, err := os.ReadFile(sourceDataFile)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Sum of products: %d\n", sumProducts(buf))

		dont, _ := regexp.Compile(`don't\(\)`)
		do, _ := regexp.Compile(`do\(\)`)
		parts := dont.Split(string(buf), -1)
		sum := sumProducts([]byte(parts[0]))
		for i := 1; i < len(parts); i++ {
			subParts := do.Split(parts[i], 2)
			if len(subParts) > 1 {
				sum += sumProducts([]byte(subParts[1]))
			}
		}

		fmt.Printf("Sum of products respecting dos and don'ts: %d\n", sum)
	},
}

func sumProducts(buf []byte) int {
	sum := 0
	r, _ := regexp.Compile(`mul\(([0-9]+),([0-9]+)\)`)
	results := r.FindAllSubmatch(buf, -1)
	for i := 0; i < len(results); i++ {
		op1, _ := strconv.ParseInt(string(results[i][1]), 10, 64)
		op2, _ := strconv.ParseInt(string(results[i][2]), 10, 64)
		sum += int(op1 * op2)
	}
	return sum
}

func init() {
	rootCmd.AddCommand(day3Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day3Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day3Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
