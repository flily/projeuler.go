package p0022

import (
	"github.com/flily/projeuler.go/framework"
)

var Problem = framework.Problem{
	Id:    22,
	Title: "Names scores",
	Description: []string{
		`Using names.txt (right click and 'Save Link/Target As...'), a 46K text file containing`,
		`over five-thousand first names, begin by sorting it into alphabetical order. Then`,
		`working out the alphabetical value for each name, multiply this value by its`,
		`alphabetical position in the list to obtain a name score.`,
		``,
		`For example, when the list is sorted into alphabetical order, COLIN, which is`,
		`worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN would`,
		`obtain a score of 938 × 53 = 49714.`,
		`What is the total of all the name scores in the file?`,
	},
	Answer: 871198282,
	Methods: map[string]framework.Solution{
		"naive": SolveNaive,
	},
}

func Load() []string {
	raw, err := framework.Import()
	if err != nil {
		panic(err)
	}

	result := make([]string, 0, 5000)
	startIndex := 0
	started := false
	for i, c := range raw {
		switch {
		case !started && c == '"':
			started = true
			startIndex = i + 1

		case started && c == '"':
			started = false
			result = append(result, string(raw[startIndex:i]))
		}
	}

	return result
}
