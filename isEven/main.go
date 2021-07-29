package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	slice := make([][]int, 0, n)
	for i := 0; i < n; i++ {
		row := make([]int, n)
		for j, _ := range row {
			fmt.Scan(&row[j])
		}
		slice = append(slice, row)
	}
	fmt.Println(slice)
}
