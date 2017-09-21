package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	if n < 2 {
		fmt.Printf("n<2")
		return
	} else if n > 50 {
		n = 50
	}
	fmt.Printf("n is %d\n", n)
	fmt.Printf("1 1 ")
	for i, j, count := 1, 1, 1; count < n; count++ {
		tmp := j
		j = i + j
		i = tmp
		fmt.Printf("%d ", j)
	}
	//fmt.Printf("n is %d", n)
}
