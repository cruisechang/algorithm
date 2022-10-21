/*
給定一個序列，設計最有效的算法，找出總合最大的區間
*/
package main

import "fmt"

var (
	sl []float32
)

func main() {
	//sl = []float32{1.5, -12.3, 3.2, -5.5, 23.2, 3.2, -1.4, -12.2, 34.2, 5.4, -7.8, 1.1, -4.9}
	//ans:4,9

	sl = []float32{1.5, -12.3, 3.2, -5.5, 23.2, 3.2, -1.4, -62.2, 44.2, 5.4, -7.8, 1.1, -4.9}
	//ans:8, 9

	//sl = []float32{1.5, -1.4, -62.2, -44.2, -5.4, -7.8, -1.1, -4.9}
	//ans:0,0

	//sl = []float32{-1.4, -62.2, -44.2, -5.4, -7.8, -1.1, -4.9, 1}
	//ans:7,7

	idx0, idx1 := method0(sl)
	fmt.Printf("method0 O(n^3) %d %d\n", idx0, idx1)

	idx0, idx1 = method1(sl)
	fmt.Printf("method1 O(n^2) %d %d\n", idx0, idx1)

	idx0, idx1 = method1(sl)
	fmt.Printf("method2 O(n) %d %d\n", idx0, idx1)
}

//O(n^3) 直覺算法
func method0(s []float32) (int, int) {
	var sumMax float32 = 0
	idx0 := 0
	idx1 := 0
	ls := len(s)
	var sum float32
	//第1循環
	for i := 0; i < ls; i++ {
		//第2循環, i-j
		for j := i; j < ls; j++ {
			sum = 0
			//第3循環, i-j加總
			for k := i; k <= j; k++ {
				sum += s[k] //這個區間的總和
			}
			//區間總和 > sum max
			if sum > sumMax {
				sumMax = sum
				idx0 = i
				idx1 = j
			}
		}

	}
	return idx0, idx1
}

//O(n^2) 直覺算法
func method1(s []float32) (int, int) {
	var sumMax float32 = 0
	idx0 := 0
	idx1 := 0
	ls := len(s)
	var sum, previousSum float32
	//第1循環
	for i := 0; i < ls; i++ {
		//第2循環, i-j
		previousSum = 0
		for j := i; j < ls; j++ {
			sum = 0
			sum = previousSum + s[j] //保留之前的sum, 每次只加最新的那個
			previousSum = sum

			//區間總和 > sum max
			if sum > sumMax {
				sumMax = sum
				idx0 = i
				idx1 = j
			}
		}
	}
	return idx0, idx1
}

/*
算右邊界、算左邊界，中間的區間就是
//狀況1 rightIndex==leftIndex
//狀況2 rightIndex 在 leftIndex 右邊
//狀況3 rightIndex 在 leftIndex 左邊
*/
func method2(s []float32) (int, int) {

	var max float32
	var rightIndex, leftIndex int
	maxIndex := len(s) - 1
	var sum float32 = 0

	//算右邊界，從左到右累加
	//累加的最大值index 就是右邊界
	//因為繼續往右加，數字變小，表示往右的是負的
	for i := 0; i <= maxIndex; i++ {
		sum = sum + s[i]
		if sum > max {
			max = sum
			rightIndex = i
		} else if sum <= 0 {
			//加到變成負的，這個負的跳過,歸0,從下一個開始
			sum = 0
		}
	}

	//算左邊界，右到左相加
	leftIndex = maxIndex
	for i := maxIndex; i >= 0; i-- {
		sum = sum + s[i]
		if sum > max {
			max = sum
			leftIndex = i
		} else if sum <= 0 {
			sum = 0
		}
	}

	return leftIndex, rightIndex
}
