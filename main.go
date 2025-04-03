package main

import (
	"flag"
	"fmt"

	common_util "github.com/BoburF/advent/common-util"
)

func main() {
  year := flag.Int("y", 2015, "year of advent")
  day := flag.Int("d", 1, "day of advent year")

  flag.Parse()

	data, err := common_util.ReadFile(*year, *day)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	fmt.Println("Result:", string(data))
}
