package main

import "fmt"

func main() {
	var num int
	num, _= fmt.Scan(&num)
	if num > 5 && num < 8 {
		fmt.Println("true")
	}
}
