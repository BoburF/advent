package solutions_2015

import (
	"bytes"
	"fmt"
	"strconv"
)

func Solution20156(data []byte) {
	linesParsed := bytes.Split(data, []byte("\n"))
	linesParsed = linesParsed[:len(linesParsed)-1]

	grid := makeGrid(1000, 1000)

	fmt.Println("executing command...")
	for _, line := range linesParsed {
		command, x, y, z, l := parseConfigurationLine(line)
		executeConfig(&grid, command, x, y, z, l)
	}

	fmt.Println("calculating...")
	counter := 0
	// for _, line := range grid {
	// 	for _, switcher := range line {
	// 		if switcher {
	// 			counter++
	// 		}
	// 	}
	// }
	for _, line := range grid {
		for _, switcher := range line {
            counter+=switcher
		}
	}


	fmt.Println("Result:", counter)
}

//	func executeConfig(grid *[][]bool, command, x, y, z, l int) {
//		for i := x; i < z+1; i++ {
//			for j := y; j < l+1; j++ {
//				target := &(*grid)[i][j]
//
//				switch command {
//				case 1:
//					*target = true
//				case 2:
//					*target = false
//				case 3:
//					*target = !*target
//				}
//			}
//		}
//	}
func executeConfig(grid *[][]int, command, x, y, z, l int) {
	for i := x; i < z+1; i++ {
		for j := y; j < l+1; j++ {
			target := &(*grid)[i][j]

			switch command {
			case 1:
				*target += 1
			case 2:
				*target -= 1
				if *target < 0 {
					*target = 0
				}
			case 3:
				*target += 2
			}
		}
	}
}

func parseConfigurationLine(line []byte) (int, int, int, int, int) {
	command := 1
	lineIdx := 0
	columnIdx := 0
	endLineIdx := 0
	endColumnIdx := 0

	initialParsedLine := bytes.Split(line, []byte(" "))

	commandEndIdx := 0
	// command parsing
	if bytes.Equal(initialParsedLine[0], []byte("turn")) {
		if bytes.Equal(initialParsedLine[1], []byte("off")) {
			command = 2
		}
		commandEndIdx = 2
	} else {
		command = 3
		commandEndIdx = 1
	}

	// coordinates parsing
	lineIdx, columnIdx = parseCoordinates(initialParsedLine[commandEndIdx])
	endLineIdx, endColumnIdx = parseCoordinates(initialParsedLine[commandEndIdx+2])

	return command, lineIdx, columnIdx, endLineIdx, endColumnIdx
}

func parseCoordinates(coordinate []byte) (int, int) {
	x, y := 0, 0
	firstParsedCoordinates := bytes.Split(coordinate, []byte(","))

	x, _ = strconv.Atoi(string(firstParsedCoordinates[0]))
	y, _ = strconv.Atoi(string(firstParsedCoordinates[1]))

	return x, y
}

//	func makeGrid(width, height int) [][]bool {
//		grid := make([][]bool, height)
//
//		for i := range height {
//			line := make([]bool, width, width)
//			grid[i] = line
//		}
//
//		return grid
//	}
func makeGrid(width, height int) [][]int {
	grid := make([][]int, height)

	for i := range height {
		line := make([]int, width, width)
		grid[i] = line
	}

	return grid
}
