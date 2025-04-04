package solutions_2015

import (
	"fmt"
)

// func Solution20153(data []byte) {
// 	x, y := 0, 0
// 	visited := make(map[string]bool)
// 	visited["0,0"] = true
//
// 	for idx, move := range data {
// 		switch move {
// 		case '^':
// 			y++
// 		case 'v':
// 			y--
// 		case '>':
// 			x++
// 		case '<':
// 			x--
// 		}
// 		pos := fmt.Sprintf("%d,%d", x, y)
//
// 		visited[pos] = true
// 	}
//
// 	fmt.Println("Result:", len(visited))
// }

func Solution20153(data []byte) {
	x, y := 0, 0
	z, l := 0, 0
	visited := make(map[string]bool)
	visited["0,0"] = true

	for idx, move := range data {
		switch move {
		case '^':
			if idx&1 == 0 {
				y++
			} else {
				l++
			}
		case 'v':
			if idx&1 == 0 {
				y--
			} else {
				l--
			}
		case '>':
			if idx&1 == 0 {
				x++
			} else {
				z++
			}
		case '<':
			if idx&1 == 0 {
				x--
			} else {
				z--
			}
		}
		pos := ""

		if idx&1 == 0 {
			pos = fmt.Sprintf("%d,%d", x, y)
		} else {
			pos = fmt.Sprintf("%d,%d", z, l)
		}

		visited[pos] = true
	}

	fmt.Println("Result:", len(visited))
}
