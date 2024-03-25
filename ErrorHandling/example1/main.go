package main

import (
	"fmt"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic cached in rocover func: ", r)
		}
	}()

	panic("random panic")

}
