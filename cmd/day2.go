/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

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
		readFile, err := os.Open(sourceDataFile)
		if err != nil {
			fmt.Println(err)
		}
		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		records := [][]int{}
		for fileScanner.Scan() {
			vals := strings.Fields(fileScanner.Text())
			record := []int{}
			for i := 0; i < len(vals); i++ {
				val, err := strconv.ParseInt(vals[i], 10, 64)
				if err != nil {
					fmt.Println("Error parsing input")
					os.Exit(1)
				}
				record = append(record, int(val))
			}
			records = append(records, record)
		}

		safeCount := 0
		safeWithDampenerCount := 0
		for i := 0; i < len(records); i++ {
			record := records[i]
			if isSafe(record) {
				// fmt.Println(record)
				safeCount++
			} else {
				for j := 0; j < len(record); j++ {
					if isSafe(remove(record, j)) {
						safeWithDampenerCount++
						break
					}
				}
			}

		}
		fmt.Printf("Safe count: %d\n", safeCount)
		fmt.Printf("Safe with dampener count: %d\n", safeCount+safeWithDampenerCount)
	},
}

func remove(slice []int, s int) []int {
	removed := append([]int{}, slice...)
	removed = append(removed[:s], removed[s+1:]...)
	// fmt.Println(s)
	// fmt.Println(slice)
	// fmt.Println(removed)
	return removed
}

func isSafe(record []int) bool {
	lastVal := record[0]
	lastDiff := 0
	for j := 1; j < len(record); j++ {
		thisVal := record[j]
		thisDiff := lastVal - thisVal
		if thisDiff*lastDiff < 0 {
			return false
		}
		if thisDiff > 3 || thisDiff < -3 || thisDiff == 0 {
			return false
		}
		lastDiff = thisDiff
		lastVal = thisVal
	}
	return true
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
