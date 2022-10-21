/*
河內塔
*/
package main

import "fmt"

func main() {

	plateNum := 8
	fmt.Printf("hanoiFunc 盤子數=%d,step=%d", plateNum, hanoiFunc(plateNum))
}

func hanoiFunc(i int) int {
	sum := 0

	if i == 1 {
		sum += i
	} else {
		sum = 2*hanoiFunc(i-1) + 1
	}

	return sum
}
