package day13

import (
	"advent2024/cmd/geometry"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Machine struct {
	A     geometry.Vector
	B     geometry.Vector
	Prize geometry.Vector
}

func (m Machine) CountA() (int, error) {
	num := m.Prize.X*m.B.Y - m.Prize.Y*m.B.X
	den := m.B.Y*m.A.X - m.A.Y*m.B.X

	integer := num / den
	fractional := float64(num) / float64(den)

	if float64(integer) == fractional {
		return integer, nil
	}
	// fmt.Println(integer)
	// fmt.Println(fractional)
	return -1, errors.New("not possible to reach the prize")
}

func (m Machine) CountB() (int, error) {
	num := m.Prize.X*m.A.Y - m.Prize.Y*m.A.X
	den := m.B.X*m.A.Y - m.A.X*m.B.Y

	integer := num / den
	fractional := float64(num) / float64(den)

	if float64(integer) == fractional {
		return integer, nil
	}
	// fmt.Println(integer)
	// fmt.Println(fractional)
	return -1, errors.New("not possible to reach the prize")
}

func LoadFile(fileName string) (machines []Machine, corrected []Machine) {
	readFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line1 := strings.Fields(fileScanner.Text())
		if len(line1) < 4 {
			fileScanner.Scan()
			line1 = strings.Fields(fileScanner.Text())
		}
		fileScanner.Scan()
		line2 := strings.Fields(fileScanner.Text())
		fileScanner.Scan()
		line3 := strings.Fields(fileScanner.Text())
		machines = append(machines, Machine{
			A:     readButton(line1),
			B:     readButton(line2),
			Prize: readPrize(line3, 0),
		})
		corrected = append(corrected, Machine{
			A:     readButton(line1),
			B:     readButton(line2),
			Prize: readPrize(line3, 10000000000000),
		})
	}
	return
}

func readButton(line []string) geometry.Vector {
	x, _ := strconv.Atoi(strings.TrimRight(strings.Split(line[2], "+")[1], ","))
	y, _ := strconv.Atoi(strings.Split(line[3], "+")[1])
	return geometry.Vector{
		X: x,
		Y: y,
	}
}

func readPrize(line []string, offset int) geometry.Vector {
	x, _ := strconv.Atoi(strings.TrimRight(strings.Split(line[1], "=")[1], ","))
	y, _ := strconv.Atoi(strings.Split(line[2], "=")[1])
	return geometry.Vector{
		X: x + offset,
		Y: y + offset,
	}
}
