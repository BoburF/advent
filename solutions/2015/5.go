package solutions_2015

import (
	"bytes"
	"fmt"
)

func Solution20155(data []byte) {
	linesParsedData := bytes.Split(data, []byte("\n"))
	linesParsedData = linesParsedData[:len(linesParsedData)-1]

	counter := 0
	for _, word := range linesParsedData {
		if !isWordNaughty(word) {
			counter++
		}
	}

	fmt.Println("Result:", counter)
}

func isWordNaughty(word []byte) bool {
	if !containsTwoDup(word) {
		return true
	}

	if !containsRepChar(word) {
		return true
	}

	return false
}

func containsTwoDup(word []byte) bool {
	seen := make(map[string]int)

	for i := 0; i < len(word)-1; i++ {
		pair := string(word[i : i+2])

		if prevIndex, exists := seen[pair]; exists {
			if i-prevIndex >= 2 {
				return true
			}
		} else {
			seen[pair] = i
		}
	}

	return false
}

func containsRepChar(word []byte) bool {
	for i := range len(word) - 2 {
		char := word[i]

		if char == word[i+2] {
			return true
		}
	}

	return false
}

// func isWordNaughty(word []byte) bool {
// 	if containsBan(word) {
// 		return true
// 	}
//
// 	if !doesContainDoubleChar(word) {
// 		return true
// 	}
//
// 	if !doesContainThreeVowels(word) {
// 		return true
// 	}
//
// 	return false
// }
//
// func containsBan(word []byte) bool {
// 	banned := [][]byte{[]byte("ab"), []byte("cd"), []byte("pq"), []byte("xy")}
//
// 	for _, ban := range banned {
// 		if bytes.Contains(word, ban) {
// 			return true
// 		}
// 	}
//
// 	return false
// }
//
// func doesContainDoubleChar(word []byte) bool {
// 	tmp := word[0]
//
// 	for i := 1; i < len(word); i++ {
// 		if tmp == word[i] {
// 			return true
// 		}
// 		tmp = word[i]
// 	}
//
// 	return false
// }
//
// func doesContainThreeVowels(word []byte) bool {
// 	vowels := map[byte]bool{
// 		'a': true,
// 		'e': true,
// 		'i': true,
// 		'o': true,
// 		'u': true,
// 	}
//
// 	counter := 0
// 	for _, ch := range word {
// 		if vowels[ch] {
// 			counter++
// 			if counter == 3 {
// 				return true
// 			}
// 		}
// 	}
//
// 	return false
// }
