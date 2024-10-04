package main

import (
	"fmt"
	"math"
)

func main() {
	sl := []int{1, 2, 3, 4, 5}
	powerSet(sl)
}

//給定一個集合，如何輸出它的冪集？
func powerSet(set []int) {

	sum := 0
	s := len(set) //集合內有幾個物件

	var x uint64 = 1
	x = x << (s - 1) //建立 bit 100000 (假設６個物件)

	for i := 0; i < s-1; i++ {
		x += uint64(math.Pow(2, float64(i))) //bit全部設1
	}

	fmt.Printf("x=%b\n", x)

	hasIndex := []int{}
	for ii := int(x); ii >= 0; ii-- {

		hasIndex = []int{}
		//取 x 每一bit的值，0 or 1
		for i := s - 1; i >= 0; i-- {
			n := x
			//fmt.Printf("n =%08b\n", n)
			//右移後,跟1 做 &, 取最後一位
			if (n>>i)&1 == 1 {
				hasIndex = append(hasIndex, i)
			}
		}
		//fmt.Printf("集合\n")
		for _, v := range hasIndex {
			fmt.Printf("%d,", set[v])
		}
		fmt.Printf("\n")
		sum++
		x -= 1
	}

	fmt.Printf("冪集數 %d\n", sum)
}
