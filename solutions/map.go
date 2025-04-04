package solutions

import solutions_2015 "github.com/BoburF/advent/solutions/2015"

var SolutionsMap = map[string]func([]byte){}

func Init() {
	SolutionsMap["2015-1"] = solutions_2015.Solution20151
	SolutionsMap["2015-2"] = solutions_2015.Solution20152
}
