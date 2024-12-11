package day10

import (
	"advent2024/cmd/geometry"
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
		vals := strings.Split(fileScanner.Text(), "")
		record := []int{}
		for _, v := range vals {
			val, err := strconv.Atoi(v)
			if err != nil {
				record = append(record, -1)
			} else {
				record = append(record, int(val))
			}
		}
		records = append(records, record)
	}
	return
}

type WayPoint struct {
	Pos geometry.Vector
	Z   int
}

func (a WayPoint) Move(area [][]int, direction geometry.Vector) (WayPoint, error) {
	newPos := a.Pos.Add(direction)
	if geometry.Box(area).Contains(newPos) {
		return WayPoint{
			Pos: newPos,
			Z:   area[newPos.X][newPos.Y],
		}, nil
	}
	return a, fmt.Errorf("out of bounds")
}

var directions []geometry.Vector = []geometry.Vector{
	{X: 0, Y: -1},
	{X: 1, Y: 0},
	{X: 0, Y: 1},
	{X: -1, Y: 0},
}

func getTrailEnds(area [][]int, start WayPoint) []WayPoint {
	ends := []WayPoint{}
	for _, d := range directions {
		next, err := start.Move(area, d)
		if err != nil {
			continue
		}
		if next.Z == start.Z+1 {
			ends = append(ends, getTrailEnds(area, next)...)
		}
	}
	if len(ends) == 0 {
		ends = append(ends, start)
	}
	return ends
}

func filter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func distinct[T comparable](list []T) []T {
	dist := map[T]bool{}
	for _, v := range list {
		dist[v] = true
	}
	keys := make([]T, 0, len(dist))
	for k := range dist {
		keys = append(keys, k)
	}
	return keys
}

func getTrailCount(area [][]int, start geometry.Vector, maxHeight int) (int, int) {
	if area[start.X][start.Y] != 0 {
		return 0, 0
	}
	trails := filter(getTrailEnds(area, WayPoint{Pos: start, Z: 0}), func(w WayPoint) bool { return w.Z == maxHeight })
	endpoints := distinct(trails)
	return len(endpoints), len(trails)
}

func GetTotalTrailCount(area [][]int, maxHeight int) (total int, rating int) {
	for i := 0; i < len(area); i++ {
		for j := 0; j < len(area[0]); j++ {
			t, r := getTrailCount(area, geometry.Vector{X: j, Y: i}, maxHeight)
			total += t
			rating += r
		}
	}
	return
}
