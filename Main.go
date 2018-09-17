package main

import "fmt"

type Pair struct {
	children []*Pair
}

func main() {
	n := 4
	fmt.Println("n = ", n)

	container := make([]*Pair, n)
	for i := 0; i < n; i++ {
		container[i] = &Pair{make([]*Pair, 0)}
	}
	//container[0] = &Pair{make([]*Pair, 1)}
	//container[0].children[0] = &Pair{make([]*Pair, 1)}
	//container[0].children[0].children[0] = &Pair{make([]*Pair, 0)}


	fmt.Println(printChildren(&container))
}

func getNextCombination(container *[]*Pair) (*[]*Pair, error) {


	return nil, nil
}

func printChildren(children *[]*Pair) string {
	if len(*children) == 0 {
		return ""
	}

	buf := ""

	for _, v := range *children {
		buf += "(" + printChildren(&v.children) + ")"
	}

	return buf
}
