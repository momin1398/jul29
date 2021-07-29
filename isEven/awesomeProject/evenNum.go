package main

import "fmt"

func main()  {
	var n int
	fmt.Scan(&n)
	if n>=1 {
		for i := 2; i <= n; i++ {
			if i%2 == 0 {
				fmt.Println(i)
			} else {
				fmt.Println("not even")
			}
		}
	}else {
		fmt.Println("not valid num")
	}

}
