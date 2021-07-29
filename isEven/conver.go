package main

import (
	"fmt"
	"strconv"
)

func main()  {

	var num1,num2 int
	var str1,str2 string
	var f1,f2 float64
	fmt.Println("enter num")
	fmt.Scan(&num1)
	str1=strconv.Itoa(num1)
   fmt.Printf("The  %T: to %T: %v\n ",num1,str1,str1)
    fmt.Println("Enter string")
     fmt.Scan(&str2)
   num2,_=strconv.Atoi(str2)
	fmt.Printf("The %T to %T %v\n ",str2,num2,num2)
   f2,_=strconv.ParseFloat(str2,64)
	fmt.Printf("The %T to %T %v\n ",str2,f2,f2)
   fmt.Println("Enter float")
   fmt.Scan(&f1)
   str1=fmt.Sprintf("%f",f1)
	fmt.Printf("The %T to  %T %v\n ",f1,str1,str1)





}
