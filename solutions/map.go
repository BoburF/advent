package solutions

import solutions_2015 "github.com/BoburF/advent/solutions/2015"

var SolutionsMap = map[string]func([]byte){}

func Init() {
	SolutionsMap["2015-1"] = solutions_2015.Solution20151
	SolutionsMap["2015-2"] = solutions_2015.Solution20152
	SolutionsMap["2015-3"] = solutions_2015.Solution20153
	SolutionsMap["2015-4"] = solutions_2015.Solution20154
	SolutionsMap["2015-5"] = solutions_2015.Solution20155
	SolutionsMap["2015-6"] = solutions_2015.Solution20156
}
