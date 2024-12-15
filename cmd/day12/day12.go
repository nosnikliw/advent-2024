package day12

import (
	"advent2024/cmd/geometry"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Plot struct {
	Crop        string
	AreaID      int
	FencedSides int
}

func LoadFile(fileName string) (plots [][]Plot) {
	readFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		vals := strings.Split(fileScanner.Text(), "")
		line := []Plot{}
		for _, v := range vals {
			line = append(line, Plot{
				Crop:        v,
				AreaID:      0,
				FencedSides: 0,
			})
		}
		plots = append(plots, line)
	}
	return
}

type RegionStats struct {
	Crop      string
	Area      int
	Perimeter int
	Sides     int
}

func DetermineRegions(plots *[][]Plot) (regions map[int]RegionStats) {
	currentID := 1
	for x, row := range *plots {
		for y, plot := range row {
			if plot.AreaID > 0 {
				continue
			}
			if assignRegion(geometry.Vector{X: x, Y: y}, currentID, plot.Crop, plots) {
				currentID++
			}
		}
	}

	regions = map[int]RegionStats{}
	for _, row := range *plots {
		for _, plot := range row {
			region := regions[plot.AreaID]
			region.Crop = plot.Crop
			region.Area += 1
			region.Perimeter += plot.FencedSides
			regions[plot.AreaID] = region
		}
	}

	for id, region := range regions {
		region.Sides = countSides(id, *plots)
		regions[id] = region
	}

	return
}

func inRegion(x int, y int, region int, plots [][]Plot) bool {
	return inArea(x, y, plots) && plots[x][y].AreaID == region
}

func countSides(region int, plots [][]Plot) int {
	sides := 0
	for x := 0; x < len(plots); x++ {
		topBoundary := false
		bottomBoundary := false
		for y := 0; y < len(plots[x]); y++ {
			plot := plots[x][y]
			if topBoundary && (plot.AreaID != region || inRegion(x-1, y, region, plots)) {
				topBoundary = false
			} else if !topBoundary && plot.AreaID == region && !inRegion(x-1, y, region, plots) {
				topBoundary = true
				sides++
			}
			if bottomBoundary && (plot.AreaID != region || inRegion(x+1, y, region, plots)) {
				bottomBoundary = false
			} else if !bottomBoundary && plot.AreaID == region && !inRegion(x+1, y, region, plots) {
				bottomBoundary = true
				sides++
			}
		}
	}
	for y := 0; y < len(plots[0]); y++ {
		leftBoundary := false
		rightBoundary := false
		for x := 0; x < len(plots); x++ {
			plot := plots[x][y]
			if leftBoundary && (plot.AreaID != region || inRegion(x, y-1, region, plots)) {
				leftBoundary = false
			} else if !leftBoundary && plot.AreaID == region && !inRegion(x, y-1, region, plots) {
				leftBoundary = true
				sides++
			}
			if rightBoundary && (plot.AreaID != region || inRegion(x, y+1, region, plots)) {
				rightBoundary = false
			} else if !rightBoundary && plot.AreaID == region && !inRegion(x, y+1, region, plots) {
				rightBoundary = true
				sides++
			}
		}
	}
	return sides
}

func inArea[T any](x int, y int, area [][]T) bool {
	return x >= 0 && x < len(area) && y >= 0 && y < len(area[x])
}

var directions []geometry.Vector = []geometry.Vector{
	{X: 0, Y: -1},
	{X: 1, Y: 0},
	{X: 0, Y: 1},
	{X: -1, Y: 0},
}

func assignRegion(pos geometry.Vector, id int, crop string, plots *[][]Plot) bool {
	if !inArea(pos.X, pos.Y, *plots) || (*plots)[pos.X][pos.Y].Crop != crop {
		return false
	}
	if (*plots)[pos.X][pos.Y].AreaID == id {
		return true
	}
	(*plots)[pos.X][pos.Y].AreaID = id
	borderCount := 0
	for _, d := range directions {
		if !assignRegion(pos.Add(d), id, crop, plots) {
			borderCount++
		}
	}
	(*plots)[pos.X][pos.Y].FencedSides = borderCount
	return true
}
