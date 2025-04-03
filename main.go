package main

import (
	"fmt"

	common_util "github.com/BoburF/advent/common-util"
)

func main() {
	fmt.Println("Hello world!")
	data, err := common_util.ReadFile(2015, 1)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	fmt.Println("Result:", string(data))
}
