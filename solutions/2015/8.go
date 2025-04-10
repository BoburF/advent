package solutions_2015

import (
	"bytes"
	"fmt"
	"strconv"
)

func Solution20158(data []byte) {
	linesParsed := bytes.Split(data, []byte("\n"))
	linesParsed = linesParsed[:len(linesParsed)-1]

	result := 0
	for _, line := range linesParsed {
		lenStr, lenStrInMem := calcTotalInLine(line)
		result += (lenStr - lenStrInMem)
	}
	fmt.Println("Result:", result)
}

func calcTotalInLine(line []byte) (int, int) {
	lineStr := string(line)

	str := strconv.Quote(lineStr)
	return len(str), len(lineStr)
}

// func calcTotalInLine(line []byte) (int, int) {
// 	lineStr := string(line)
//
// 	str, _ := strconv.Unquote(lineStr)
// 	return len(lineStr), len(str)
// }
