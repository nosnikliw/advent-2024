/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// day4Cmd represents the day4 command
var day4Cmd = &cobra.Command{
	Use:   "day4",
	Short: "Solve day 4",
	Long: `Solve day 4

Word search`,
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		sourceDataFile := "./data/day4/input.txt"
		if len(args) > 0 {
			sourceDataFile = args[0]
		}
		readFile, err := os.Open(sourceDataFile)
		if err != nil {
			fmt.Println(err)
		}
		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		grid := [][]byte{}
		for fileScanner.Scan() {
			grid = append(grid, []byte(fileScanner.Text()))
		}
		xmas := []byte{'X', 'M', 'A', 'S'}
		count := countOccurences(grid, xmas)

		fmt.Printf("Count in grid: %d\n", count)

		xCount := countCrossedMas(grid)

		fmt.Printf("Count of crosses: %d\n", xCount)
	},
}

func countCrossedMas(grid [][]byte) int {
	count := 0
	for x := 1; x < (len(grid) - 1); x++ {
		for y := 1; y < (len(grid[x]) - 1); y++ {
			if checkLetterMatch(x, y, grid, 'A') {
				count += countCrossesAtPos(x, y, grid)
			}
		}
	}
	return count
}

func countCrossesAtPos(x, y int, grid [][]byte) int {
	count := 0
	word := []byte{'M', 'A', 'S'}

	d1Match := checkWordMatch(x-1, y-1, 1, 1, grid, word) + checkWordMatch(x+1, y+1, -1, -1, grid, word)
	d2Match := checkWordMatch(x+1, y-1, -1, 1, grid, word) + checkWordMatch(x-1, y+1, 1, -1, grid, word)

	if (d1Match == 1) && (d2Match == 1) {
		count += 1
	}
	return count
}

func countOccurences(grid [][]byte, word []byte) int {
	count := 0
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			count += countStartingFromPosition(x, y, grid, word)
		}
	}
	return count
}

func countStartingFromPosition(x, y int, grid [][]byte, word []byte) int {
	count := checkWordMatch(x, y, 0, 1, grid, word)
	count += checkWordMatch(x, y, 1, 1, grid, word)
	count += checkWordMatch(x, y, 1, 0, grid, word)
	count += checkWordMatch(x, y, 1, -1, grid, word)
	count += checkWordMatch(x, y, 0, -1, grid, word)
	count += checkWordMatch(x, y, -1, -1, grid, word)
	count += checkWordMatch(x, y, -1, 0, grid, word)
	count += checkWordMatch(x, y, -1, 1, grid, word)
	return count
}

func checkWordMatch(x, y, dx, dy int, grid [][]byte, word []byte) int {
	if x < 0 || x >= len(grid) {
		return 0
	}
	if y < 0 || y >= len(grid[x]) {
		return 0
	}
	for i := 0; i < len(word); i++ {
		xpos := x + i*dx
		ypos := y + i*dy
		if !checkLetterMatch(xpos, ypos, grid, word[i]) {
			return 0
		}
	}
	return 1
}

func checkLetterMatch(x, y int, grid [][]byte, letter byte) bool {
	if x < 0 || x >= len(grid) {
		return false
	}
	if y < 0 || y >= len(grid[x]) {
		return false
	}
	if grid[x][y] == letter {
		return true
	}
	return false
}

func init() {
	rootCmd.AddCommand(day4Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day4Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day4Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
