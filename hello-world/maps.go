package main

import "fmt"

func LearnMaps() {
	//fmt.Println("Learning Maps")
	lookup := make(map[string]int)
	lookup["Score"] = 5

	score, ex := lookup["score"]

	fmt.Println(score, ex)
	fmt.Println("Length", len(lookup))
	fmt.Println(lookup)

	delete(lookup, "Score")

	fmt.Println(lookup)
}
