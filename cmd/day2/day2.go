package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func LoadFile(fileName string) (records [][]int) {
	readFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

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
	return
}

func CountSafeRecords(records [][]int) (safe int, safeWithDampener int) {
	for i := 0; i < len(records); i++ {
		record := records[i]
		if isSafe(record) {
			safe++
		} else {
			for j := 0; j < len(record); j++ {
				if isSafe(remove(record, j)) {
					safeWithDampener++
					break
				}
			}
		}
	}
	return
}

func remove(slice []int, s int) []int {
	removed := append([]int{}, slice...)
	removed = append(removed[:s], removed[s+1:]...)
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
