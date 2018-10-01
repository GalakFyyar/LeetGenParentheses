package main

import (
	"fmt"
	"time"
)

var mapOfResults map[int][]string

func main() {
	n := 16

	mapOfResults = map[int][]string{
		2: {"(())", "()()"},
	}

	for i := 0; i < 4; i++ {
		start := time.Now()
		x := generateParenthesesNoAppend(n)
		//generateParentheses(n)
		printSet(x)
		elapsed := time.Since(start)
		fmt.Println("Set", i, elapsed)
	}

	fmt.Println("DONE")
}

func generateParenthesesByteSliceNoAppend(n int) *[][]byte {
	if n == 2 {
		return &[][]byte{[]byte("(())"), []byte("()()")}
	}

	baseSet := *generateParenthesesByteSliceNoAppend(n - 1)
	baseLength := len(baseSet)

	returnSet := make([][]byte, baseLength*3-1)

	firstThird := returnSet[0:baseLength]
	secondThird := returnSet[baseLength : baseLength*2]
	lastThird := returnSet[baseLength*2:]

	//append parentheses on the left for the first third
	for i := range firstThird {
		firstThird[i][0] = '['
		firstThird[i][1] = ']'
		copy(firstThird[2:], baseSet)
	}

	//surround with parentheses for the second third
	for i := range secondThird {
		secondThird[i] = []byte{'<'} // + baseSet[i] + ">"
	}

	//append parentheses on the right for the last third
	for i := range lastThird {
		//lastThird[i] = baseSet[i] + "{}"
		j := 0
		for ; j < len(baseSet); j++ {
			lastThird[i][j] = baseSet[i][j]
		}
		lastThird[i][j] = '{'
		j++
		lastThird[i][j] = '}'
	}

	return &returnSet
}

func generateParenthesesNoAppend(n int) *[]string {
	if n == 2 {
		return &[]string{"(())", "()()"}
	}

	baseSet := *generateParenthesesNoAppend(n - 1)
	baseLength := len(baseSet)

	returnSet := make([]string, baseLength*3-1)

	firstThird := returnSet[0:baseLength]
	secondThird := returnSet[baseLength : baseLength*2]
	lastThird := returnSet[baseLength*2:]

	//append parentheses on the left for the first third
	for i := range firstThird {
		firstThird[i] = "[]" + baseSet[i]
	}

	//surround with parentheses for the second third
	for i := range secondThird {
		secondThird[i] = "<" + baseSet[i] + ">"
	}

	//append parentheses on the right for the last third
	for i := range lastThird {
		lastThird[i] = baseSet[i] + "{}"
	}

	return &returnSet
}

func generateParenthesesWithMap(n int) *[]string {
	elem, ok := mapOfResults[n]
	if ok {
		fmt.Println("hit", n)
		return &elem
	}

	baseSet := *generateParenthesesWithMap(n - 1)

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

func generateParentheses(n int) *[]string {
	if n == 2 {
		return &[]string{"(())", "()()"}
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

	return &returnSet
}

//noinspection GoUnusedFunction
func runWithPrint(n int) {
	setOfParentheses := generateParentheses(n)
	printSet(setOfParentheses)
}

func printSet(setOfParentheses *[]string) {
	for _, str := range *setOfParentheses {
		fmt.Println(str)
	}
}

func printSetByte(setOfParentheses *[][]byte) {
	for _, str := range *setOfParentheses {
		fmt.Println(str)
	}
}
