package main

import (
	"fmt"
	"sort"
	"strings"
)

func LearnArrays() {

	//fmt.Println("\nUsing Array and slices")
	//books := []string{"abc", "def", "xyz"}
	//books := make([]string, 5)
	//var books []string

	books := make([]string, 0, 10)
	books = append(books, "Abc")
	books = books[0:8]
	books = append(books, "Zyz")
	fmt.Println(books)

	haystack := "the spice mush flow"

	fmt.Println(strings.Index(haystack[5:], " "))

	scores := make([]int, 100)
	for i := 0; i < 100; i++ {
		scores[i] = i + 1 //int(rand.Int31n(1000))
	}
	sort.Ints(scores)

	worst := make([]int, 5)
	copy(worst[3:], scores[1:3])
	fmt.Println(worst)

}
