package solutions_2015

import (
	"bytes"
	"fmt"
	"strconv"
)

func Solution20152(data []byte) {
	giftDimensions := make(map[int][3]int)

	lineParsedData := bytes.Split(data, []byte("\n"))
	lineParsedData = lineParsedData[:len(lineParsedData)-1]

	for idx, line := range lineParsedData {
		giftDimension := parseDimensions(line)
		giftDimensions[idx] = giftDimension
	}

	fmt.Println("Result:", calculate(giftDimensions))
}

// func calculate(giftDimensions map[int][3]int) int {
// 	totalAreaNeeded := 0
//
// 	for _, gift := range giftDimensions {
// 		x := 2 * gift[0] * gift[1]
// 		y := 2 * gift[1] * gift[2]
// 		z := 2 * gift[2] * gift[0]
//
// 		smallest := min(x, y, z)
//
// 		area := x + y + z + (smallest / 2)
//
// 		totalAreaNeeded += area
// 	}
// 	return totalAreaNeeded
// }

func calculate(giftDimensions map[int][3]int) int {
	totalNeeded := 0

	for _, gift := range giftDimensions {
		x := (2 * gift[0]) + (2 * gift[1]) + (2 * gift[2]) - (max(gift[0], gift[1], gift[2]) * 2)
		y := gift[0] * gift[1] * gift[2]

		totalNeeded += x + y
	}
	return totalNeeded
}

func parseDimensions(data []byte) [3]int {
	dimensions := make([]int, 0, 3)

	dimensionsStr := bytes.Split(data, []byte("x"))

	for _, dimension := range dimensionsStr {
		number, _ := strconv.Atoi(string(dimension))
		dimensions = append(dimensions, number)
	}

	return [3]int(dimensions)
}
