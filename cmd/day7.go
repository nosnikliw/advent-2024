/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// day7Cmd represents the day7 command
var day7Cmd = &cobra.Command{
	Use:   "day7",
	Short: "Solve day 7",
	Long: `Solve day 7

Bridge Repair`,
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		sourceDataFile := "./data/day7/input.txt"
		if len(args) > 0 {
			sourceDataFile = args[0]
		}

		calibrations := loadDay7Input(sourceDataFile)

		total := 0
		total2 := 0
		for _, c := range calibrations {
			if possiblyCorrect(c, []Operator{add, mul}) {
				total += c.result
			}
			if possiblyCorrect(c, []Operator{add, mul, app}) {
				total2 += c.result
			}
		}

		fmt.Printf("Calibration result: %d\n", total)
		fmt.Printf("Calibration result 2: %d\n", total2)
	},
}

type Operator func(int, int) int

func add(a int, b int) int { return a + b }
func mul(a int, b int) int { return a * b }
func app(a int, b int) int {
	val, _ := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
	return val
}

func getOperatorPermutations(operands int, operators []Operator) [][]Operator {
	count := math.Pow(float64(len(operators)), float64(operands-1))
	permutations := [][]Operator{}
	for i := 0; float64(i) < count; i++ {
		ops := strings.Split(strconv.FormatInt(int64(i), len(operators)), "")
		perm := []Operator{}
		if len(ops) < (operands - 1) {
			missingOps := make([]string, (operands - 1 - len(ops)))
			for i := 0; i < len(missingOps); i++ {
				missingOps[i] = "0"
			}
			ops = append(missingOps, ops...)
		}
		for _, o := range ops {
			idx, _ := strconv.Atoi(o)
			perm = append(perm, operators[idx])
		}
		permutations = append(permutations, perm)
	}
	return permutations
}

func possiblyCorrect(cal Calibration, operators []Operator) bool {
	permutations := getOperatorPermutations(len(cal.operands), operators)
	for _, p := range permutations {
		result := cal.operands[0]
		for j := 0; j < len(p); j++ {
			result = p[j](result, cal.operands[j+1])
		}
		if result == cal.result {
			return true
		}
	}
	return false
}

type Calibration struct {
	result   int
	operands []int
}

func loadDay7Input(fileName string) []Calibration {
	readFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	calibrations := []Calibration{}
	for fileScanner.Scan() {
		parts := strings.Split(fileScanner.Text(), ":")
		result, _ := strconv.Atoi(parts[0])
		operands := []int{}
		for _, s := range strings.Fields(parts[1]) {
			op, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			operands = append(operands, op)
		}
		calibrations = append(calibrations, Calibration{result: result, operands: operands})
	}
	return calibrations
}

func init() {
	rootCmd.AddCommand(day7Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day7Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day7Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
