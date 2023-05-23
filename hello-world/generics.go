package main

import "fmt"

type Number interface {
	int64 | float64
}

func SumInts(m map[string]int64) int64 {
	var sum int64
	for _, v := range m {
		sum += v
	}
	return sum
}

func SumFloats(m map[string]float64) float64 {
	var sum float64
	for _, v := range m {
		sum += v
	}
	return sum
}

func Sum[V Number](m map[string]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum
}

func LearnGenerics() {
	//intSum := SumInts(map[string]int64{"a": 4, "b": 8})
	//floatSum := SumFloats(map[string]float64{"a": 5.6, "b": 7.88})

	intSum := Sum(map[string]int64{"a": 4, "b": 88})
	floatSum := Sum(map[string]float64{"a": 8.44, "b": 82.5})
	fmt.Println(intSum, floatSum)
}
