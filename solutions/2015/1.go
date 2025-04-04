package solutions_2015

import "fmt"

// func Solution20151(data []byte) {
// 	count := 0
// 	for _, ch := range data {
// 		switch ch {
// 		case '(':
// 			count++
// 		case ')':
// 			count--
// 		}
// 	}
//
// 	fmt.Println("Result:", count)
// }

func Solution20151(data []byte) {
	count := 0
	for idx, ch := range data {
		switch ch {
		case '(':
			count++
		case ')':
			count--
		}

		if count == -1 {
			fmt.Println("Result:", idx+1)
			break
		}
	}
}
