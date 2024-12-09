/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// day9Cmd represents the day9 command
var day9Cmd = &cobra.Command{
	Use:   "day9",
	Short: "Solve day 9",
	Long: `Solve day 9

Disk fragmenter`,
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		sourceDataFile := "./data/day9/input.txt"
		if len(args) > 0 {
			sourceDataFile = args[0]
		}

		disk, fileIndex, spaceIndex := buildDiskImage(sourceDataFile)
		newDisk := []int64{}
		for i, j := 0, len(disk); i < j; i++ {
			if disk[i] >= 0 {
				newDisk = append(newDisk, disk[i])
			} else {
				j--
				for disk[j] == -1 {
					j--
				}
				newDisk = append(newDisk, disk[j])
			}
		}
		checksum := int64(0)
		for i, v := range newDisk {
			checksum += int64(i) * v
		}
		fmt.Printf("Checksum: %d\n", checksum)

		for i := len(fileIndex) - 1; i > 0; i-- {
			f := fileIndex[i]
			for j := 0; j < len(spaceIndex); j++ {
				s := spaceIndex[j]
				if s.pos >= f.pos {
					break
				}
				if s.size >= f.size {
					fileIndex[i].pos = s.pos
					spaceIndex[j].size = s.size - f.size
					spaceIndex[j].pos = s.pos + f.size
					break
				}
			}
		}
		checksum2 := 0
		for _, v := range fileIndex {
			for i := 0; i < v.size; i++ {
				checksum2 += v.val * (i + v.pos)
			}
		}
		fmt.Printf("Checksum 2: %d\n", checksum2)
	},
}

type block struct {
	pos  int
	size int
	val  int
}

func buildDiskImage(fileName string) (disk []int64, fileIndex []block, spaceIndex []block) {
	readFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanRunes)

	isFile := true
	fileNo := int64(0)
	for fileScanner.Scan() {
		size, err := strconv.Atoi(fileScanner.Text())
		if err != nil {
			panic(err)
		}
		thisVal := int64(-1)
		if isFile {
			thisVal = fileNo
			fileIndex = append(fileIndex, block{pos: len(disk), size: size, val: int(fileNo)})
		} else {
			spaceIndex = append(spaceIndex, block{pos: len(disk), size: size, val: -1})
		}

		for i := 0; i < size; i++ {
			disk = append(disk, thisVal)
		}
		if isFile {
			fileNo++
		}
		isFile = !isFile
	}
	return
}

func init() {
	rootCmd.AddCommand(day9Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day9Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day9Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
