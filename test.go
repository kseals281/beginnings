package main

import "fmt"

func main() {
	once := true
	for true {
		if fmt.Println(once == true); once == true {
			defer fmt.Println("world")
			once = false
		}

		fmt.Println("hello")
		var pause string
		fmt.Scan(&pause)
	}
}
