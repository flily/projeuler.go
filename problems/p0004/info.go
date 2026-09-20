package p0004

import (
	"github.com/flily/projeuler.go/framework"
)

var Problem = framework.InitProblem(4, "Largest Palindrome Product").
	WithAnswer(906609).
	Solution("naive-string", SolveNaiveString).
	Solution("naive-int", SolveNaiveInt).
	Solution("naive-int6", SolveNaiveInt6).
	Solution("naive-int6-and", SolveNaiveInt6And).
	Solution("generator-channel", SolveGeneratorChannel).
	Solution("generator6-channel", SolveGenerator6Channel).
	WithDescription(
		`A palindromic number reads the same both ways. The largest palindrome made from the`,
		`product of two`,
		`2-digit numbers is 9009 = 91 × 99.`,
		``,
		`Find the largest palindrome made from the product of two 3-digit numbers.`,
	)
