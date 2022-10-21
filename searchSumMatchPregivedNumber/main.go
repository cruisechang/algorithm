/*
給定一個長度大於2的陣列，在陣列中找出一個區間，使得區間和，等於給定的數字
*/
package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 5, 6, 7}
	want := 13
	idx0, idx1 := method0(sl, want)
	fmt.Printf("method0 O(n^3) %d %d\n", idx0, idx1)

	idx0, idx1 = method1(sl, want)
	fmt.Printf("method1 O(n^2) %d %d\n", idx0, idx1)
}

//直覺算法
//O(n^3)
func method0(sl []int, want int) (int, int) {

	llen := len(sl)

	var sum int
	for i := 0; i < llen; i++ {
		if sl[i] == want {
			return i, i
		}
		for j := i; j < llen; j++ {
			sum = 0
			for k := i; k <= j; k++ {
				sum += sl[k]
				if sum == want {
					return i, j
				}
			}
		}
	}
	return -1, -1
}

//改良算法
//O(n^2)
func method1(sl []int, want int) (int, int) {

	llen := len(sl)

	var sum int
	for i := 0; i < llen; i++ {
		if sl[i] == want {
			return i, i
		}

		j := i + 1
		sum = sl[i]

		for sum < want && j < llen {
			sum += sl[j]
			if sum == want {
				return i, j
			}
			j++
		}

	}
	return -1, -1
}

//O(n^2)
func method2(sl []int, want int) (int, int) {

	llen := len(sl)

	var sum int
	for i := 0; i < llen; i++ {
		if sl[i] == want {
			return i, i
		}

		j := i + 1
		sum = sl[i]

		for sum < want && j < llen {
			sum += sl[j]
			if sum == want {
				return i, j
			}
			j++
		}

	}
	return -1, -1
}
