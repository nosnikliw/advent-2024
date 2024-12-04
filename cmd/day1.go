/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

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
		readFile, err := os.Open(sourceDataFile)
		if err != nil {
			fmt.Println(err)
		}
		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		list1 := []int{}
		list2 := []int{}
		for fileScanner.Scan() {
			vals := strings.Fields(fileScanner.Text())
			val1, err1 := strconv.ParseInt(vals[0], 10, 64)
			if err1 != nil {
				fmt.Println("Error parsing input")
				os.Exit(1)
			}
			list1 = append(list1, int(val1))
			val2, err2 := strconv.ParseInt(vals[1], 10, 64)
			if err2 != nil {
				fmt.Println("Error parsing input")
				os.Exit(1)
			}
			list2 = append(list2, int(val2))
		}
		readFile.Close()

		sort.Ints(list1)
		sort.Ints(list2)
		distance := 0
		for i := 0; i < len(list1); i++ {
			d := list1[i] - list2[i]
			if d < 0 {
				distance = distance - d
			} else {
				distance = distance + d
			}
		}
		fmt.Printf("Distance: %d\n", distance)
		similarity := 0
		cursor := 0
		for i := 0; i < len(list1); i++ {
			val := list1[i]
			for ; len(list2) > cursor && val > list2[cursor]; cursor++ {
			}
			for ; len(list2) > cursor && val == list2[cursor]; cursor++ {
				similarity = similarity + val
			}
		}
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
