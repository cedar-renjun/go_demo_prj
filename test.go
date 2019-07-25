package main

import (
    //"flag"
    "fmt"
    "strconv"
)

var server_addrs = [] string {":8890", ":8891",":8892",":8893"}

func main() {

	str := strconv.Itoa(134)
	str1 := "hello,world"
	str2 := str1+str
	fmt.Println(str2)

	for i:= 0; i < 4; i++{
		fmt.Println(server_addrs[i])
	}

	fmt.Println("finish")
}

