package main

import (
	"fmt"
	"piscine"
)

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}
	a3 := []int{62836, -174282, -125435, -402558, 326003, 190315, -312407, -465953}
	result1 := piscine.IsSorted(piscine.TwoNumberSort, a1)
	result2 := piscine.IsSorted(piscine.TwoNumberSort, a2)
	result3 := piscine.IsSorted(piscine.TwoNumberSort, a3)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
}
