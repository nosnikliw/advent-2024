package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func LoadFile(fileName string) (list1 []int, list2 []int) {
	readFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
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
	return
}

func CalculateDistance(list1 []int, list2 []int) (distance int) {
	for i := 0; i < len(list1); i++ {
		d := list1[i] - list2[i]
		if d < 0 {
			distance = distance - d
		} else {
			distance = distance + d
		}
	}
	return
}

func CalculateSimilarity(list1 []int, list2 []int) (similarity int) {
	cursor := 0
	for i := 0; i < len(list1); i++ {
		val := list1[i]
		for ; len(list2) > cursor && val > list2[cursor]; cursor++ {
		}
		for ; len(list2) > cursor && val == list2[cursor]; cursor++ {
			similarity = similarity + val
		}
	}
	return
}
