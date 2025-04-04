package solutions_2015

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func Solution20154(_ []byte) {
	n := 0

	for {
		seed := append([]byte("ckczppom"), []byte(strconv.Itoa(n))...)

		generatedData := md5.Sum(seed)

		hexStr := fmt.Sprintf("%x", generatedData)

		if strings.HasPrefix(hexStr, "000000") {
			fmt.Println("Result:", n)
			break
		}
		n++
	}

}
