package main

import (
	"os"
	"strconv"
)

//import "strconv"

func fibRec(a int) int {
	if a <= 1 {
		return a
	}
	return fibRec(a-1) + fibRec(a-2)
}

func fibIteration(a int) int {
	var fib, prev1, prev2 int
	prev1 = 0
	prev2 = 1
	if a <= 1 {
		return a
	}
	fib = prev1 + prev2
	for i := 3; i <= a; i++ {
		prev1 = prev2
		prev2 = fib
		fib  = prev1 + prev2
	}
	return fib
}

func main() {
	param, err := strconv.Atoi(os.Args[1])

	if err != nil {
		panic(err)
	}
	println("fibbo recursive for: ", param, " = ", fibRec(param))
	println("fibbo by itertion for: ", param, " = ", fibIteration(param)) 
}
