/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// day5Cmd represents the day5 command
var day5Cmd = &cobra.Command{
	Use:   "day5",
	Short: "Solve day 5",
	Long: `Solve day 5

Printing order`,
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		sourceDataFile := "./data/day5/input.txt"
		if len(args) > 0 {
			sourceDataFile = args[0]
		}

		before, after, updates := loadInput(sourceDataFile)

		result := 0
		invalidUpdates := [][]string{}
		for i := 0; i < len(updates); i++ {
			update := updates[i]
			if updateIsValid(update, before, after) {
				middleValue, _ := strconv.ParseInt(update[len(update)/2], 10, 64)
				result += int(middleValue)
			} else {
				invalidUpdates = append(invalidUpdates, update)
			}
		}
		fmt.Printf("Valid updates sum: %d\n", result)
		correctedUpdates := [][]string{}
		correctedResult := 0
		for i := 0; i < len(invalidUpdates); i++ {
			correctedUpdates = append(correctedUpdates, fixUpdate(invalidUpdates[i], before))
			middleValue, _ := strconv.ParseInt(correctedUpdates[i][len(correctedUpdates[i])/2], 10, 64)
			correctedResult += int(middleValue)
		}
		fmt.Printf("Corrected updates sum: %d\n", correctedResult)
	},
}

func fixUpdate(update []string, before map[string][]string) []string {

	result := []string{}
	toPlace := []string{}
	toPlace = append(toPlace, update...)
	for len(toPlace) > 0 {
		for i := 0; i < len(toPlace); i++ {
			neededBefore := filter(before[toPlace[i]], func(s string) bool { return slices.Contains(toPlace, s) })
			if len(neededBefore) == 0 {
				result = append(result, toPlace[i])
				toPlace = removeString(toPlace, i)
				break
			}
			if i+1 == len(toPlace) {
				panic("Impossible to correct")
			}
		}
	}

	return result
}

func filter(ss []string, test func(string) bool) (ret []string) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func removeString(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func updateIsValid(update []string, before map[string][]string, after map[string][]string) bool {
	for i := 0; i < len(update); i++ {
		beforeRule := before[update[i]]
		for j := 0; j < len(beforeRule); j++ {
			if slices.Contains(update[i+1:], beforeRule[j]) {
				return false
			}
		}
		afterRule := after[update[i]]
		for j := 0; j < len(afterRule); j++ {
			if slices.Contains(update[:i], afterRule[j]) {
				return false
			}
		}
	}
	return true
}

func loadInput(fileName string) (map[string][]string, map[string][]string, [][]string) {
	readFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	before := map[string][]string{}
	after := map[string][]string{}
	for fileScanner.Scan() && strings.TrimSpace(fileScanner.Text()) != "" {
		rule := strings.Split(strings.TrimSpace(fileScanner.Text()), "|")
		before[rule[1]] = append(before[rule[1]], rule[0])
		after[rule[0]] = append(after[rule[0]], rule[1])
	}

	updates := [][]string{}
	for fileScanner.Scan() {
		updates = append(updates, strings.Split(strings.TrimSpace(fileScanner.Text()), ","))
	}
	return before, after, updates
}

func init() {
	rootCmd.AddCommand(day5Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day5Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day5Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
