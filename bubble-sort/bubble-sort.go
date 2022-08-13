package main

import "fmt"

func sort(a []int) {
	for i := len(a) - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	fmt.Println(a)
}

func main() {
	sort([]int{1})
	sort([]int{2, 1})
	sort([]int{5, 4, 2, 3, 1})
	sort([]int{2, 1, 1, 2})
}
