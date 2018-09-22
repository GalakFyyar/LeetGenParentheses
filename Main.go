package main

import (
	"fmt"
	"time"
)

var mapOfResults map[int][]string

func main() {
	n := 17

	mapOfResults = map[int][]string{
		2: {"(())", "()()"},
	}

	start := time.Now()
	generateParentheses(n)
	//runWithPrint(n)
	elapsed := time.Since(start)

	fmt.Println("DONE ", elapsed)
}

func generateParentheses(n int) *[]string {
	elem, ok := mapOfResults[n]
	if ok {
		return &elem
	}

	baseSet := *generateParentheses(n - 1)

	var returnSet []string

	//append parentheses on the left
	for _, str := range baseSet {
		returnSet = append(returnSet, "[]"+str)
	}

	//surround with parentheses
	for _, str := range baseSet {
		returnSet = append(returnSet, "<"+str+">")
	}

	//append parentheses on the right
	for i := 0; i < len(baseSet)-1; i++ {
		returnSet = append(returnSet, baseSet[i]+"{}")
	}

	//add to results
	mapOfResults[n] = returnSet
	return &returnSet
}

//noinspection GoUnusedFunction
func runWithPrint(n int) {
	setOfParentheses := *generateParentheses(n)
	for _, paren := range setOfParentheses {
		fmt.Println(paren)
	}
}
