package day11

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func LoadFile(fileName string) (stones []string) {
	readFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		vals := strings.Fields(fileScanner.Text())
		stones = append(stones, vals...)
	}
	return
}

func Blink(before []string) (after []string) {
	for _, stone := range before {
		if stone == "0" {
			after = append(after, "1")
		} else if len(stone)%2 == 0 {
			length := len(stone) / 2
			after = append(after, stone[:length])
			m, err := strconv.Atoi(stone[length:])
			if err != nil {
				panic(err)
			}
			after = append(after, strconv.Itoa(m))
		} else {
			m, err := strconv.Atoi(stone)
			if err != nil {
				panic(err)
			}
			after = append(after, strconv.Itoa(m*2024))
		}
	}
	return
}

type ValLevel struct {
	Val   int
	Level int
}

func CountStones(stone int, level int, counted *map[ValLevel]int) (count int) {
	defer func() { (*counted)[ValLevel{stone, level}] = count }()
	if level <= 0 {
		count = 1
		return
	}
	if (*counted)[ValLevel{stone, level}] > 0 {
		count = (*counted)[ValLevel{stone, level}]
		return
	}
	if stone == 0 {
		count = CountStones(1, level-1, counted)
		return
	} else if len(strconv.Itoa(stone))%2 == 0 {
		str := strconv.Itoa(stone)
		length := len(str) / 2
		n, _ := strconv.Atoi(str[:length])
		c1 := CountStones(n, level-1, counted)
		m, _ := strconv.Atoi(str[length:])
		c2 := CountStones(m, level-1, counted)
		count = c1 + c2
		return
	} else {
		count = CountStones(stone*2024, level-1, counted)
		return
	}
}

func filter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func Count(stones []string, times int) int {
	count := len(stones)
	current := stones
	for i := 0; i < times; i++ {
		toSplit := filter(current, func(s string) bool { return len(s)%2 == 0 })
		fmt.Println(toSplit)
		count += len(toSplit)
		current = []string{}
		for _, s := range toSplit {
			half := len(s) / 2
			n, _ := strconv.Atoi(s[half:])
			current = append(current, s[:half], strconv.Itoa(n))
		}
		// count += len(current)
	}
	return count
}
