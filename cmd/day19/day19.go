package day19

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LoadFile(fileName string) (towels []string, designs []string) {
	readFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	fileScanner.Scan()
	towels = strings.Split(fileScanner.Text(), ", ")
	fileScanner.Scan()

	for fileScanner.Scan() {
		designs = append(designs, fileScanner.Text())
	}
	return
}

func FindPossible(towels []string, designs []string) (possible []string, count int) {
	calculated := map[string]int{}
	for _, design := range designs {
		ways := canMake(towels, design, &calculated)
		if ways > 0 {
			possible = append(possible, design)
			count += ways
		}
	}
	return
}

func canMake(towels []string, design string, calculated *map[string]int) int {
	if design == "" {
		return 1
	}
	if c, found := (*calculated)[design]; found {
		return c
	}
	ways := 0
	for _, t := range towels {
		if len(t) <= len(design) && design[:len(t)] == t {
			ways += canMake(towels, design[len(t):], calculated)
		}
	}
	(*calculated)[design] = ways
	return ways
}
