package main

import (
    //"flag"
    "fmt"
    //"strconv"
)

func f() (result int){

	defer func(result int){
		fmt.Println("defer")
		result += 5
	}(result)

	fmt.Println("fun finish")
	return 1
}

func main() {

	res := 0

	res = f()

	fmt.Println("finish", res)
}

