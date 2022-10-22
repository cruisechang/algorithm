/*
遞迴應用
遞迴在執行時是用堆疊(stack)格式存在記憶體，
堆疊的意義是先進後出。

function的遞迴是function嵌套的概念。
func0(func1(func2()))
func0最後執行。
func0最先進入stack,func1次之...
子func會return值父一個func。

只要是重複執行可以解決的，都可以用recursion，
1 資料定義是按照遞迴定義的
2 問題解法事按照遞迴演算法實現的
3 資料的結構是按照遞迴定義的
*/
package main

import "fmt"

func main() {

	fmt.Printf("countFibonacci() digit amount=%d, ans=%d\n", 10, countFibonacci(10))
	fmt.Printf("countFibonacciByLoop() digit amount=%d, ans=%d\n", 6, countFibonacciByLoop(6))
	printFibonacci(0, 10, 0, 0)

	plateNum := 4
	fmt.Printf("\ntowerOfHanoiByTrick()========\n")
	fmt.Printf("towerOfHanoiByTrick()======== 盤子數=%d,step=%d\n", plateNum, towerOfHanoiByTrick(plateNum, 'A', 'B', 'C'))
	fmt.Printf("countStepOfTowerOfHanoiByFormula() 盤子數=%d,step=%d\n", plateNum, countStepOfTowerOfHanoiByFormula(plateNum))

	a, b := 546, 429
	fmt.Printf("\ngreatestCommonFactor() a=%d,b=%d, ans=%d\n", a, b, greatestCommonFactor(a, b))

}

//這算法跟merge sort(歸併排序)一樣
//概念是預設前一階段有解答，
//用遞迴推進到最初的階段，
//最初階段取得答案，
//再依序往後返回。
func countFibonacci(amount int) int {
	if amount < 2 {
		return amount
	}
	return countFibonacci(amount-2) + countFibonacci(amount-1)
}
func countFibonacciByLoop(amount int) int {
	a, b := amount%2, 1
	for i := 0; i < amount/2; i++ {
		a += b
		b += a
	}
	return a
}

//印出斐波那契數列
//資料定義規則重複
func printFibonacci(times, amount, last2, last1 int) {

	if times > amount {
		return
	}
	if times == 0 {
		fmt.Printf("printFibonacci %d\n", last2)
		times++
		printFibonacci(times, amount, last2, last1)
	} else if times == 1 {
		last1 = 1
		times++
		fmt.Printf("printFibonacci %d\n", last1)
		printFibonacci(times, amount, last2, last1)
	} else {
		times++
		fmt.Printf("printFibonacci %d\n", last2+last1)
		printFibonacci(times, amount, last1, last2+last1)
	}
}

//河內塔
//計算移動次數
//公式解法
//符合遞迴的解法規則重複
func countStepOfTowerOfHanoiByFormula(plateNum int) int {
	sum := 0

	if plateNum == 1 {
		sum += plateNum
	} else {
		sum = 2*countStepOfTowerOfHanoiByFormula(plateNum-1) + 1
	}

	return sum
}

//河內塔
//印出步驟
//智慧解法(並非真的依據河內塔一步一步去算)
func towerOfHanoiByTrick(plateNum int, pillarA, pillarB, pillarC rune) int {

	if plateNum == 1 {
		fmt.Printf("Hanoi step: Move a plate from %c to %c.\n", pillarA, pillarC)
		return 1
	} else {
		a := towerOfHanoiByTrick(plateNum-1, pillarA, pillarC, pillarB)
		b := towerOfHanoiByTrick(1, pillarA, pillarB, pillarC)
		c := towerOfHanoiByTrick(plateNum-1, pillarB, pillarA, pillarC)
		return a + b + c
	}
}

//求最大公因數
//解法規則不斷重複
func greatestCommonFactor(a, b int) int {
	if a < b {
		//a,b交換
		a = a ^ b
		b = a ^ b
		a = a ^ b
		//a=a+b
		//b=a-b
		//a=a-b
	}

	if a%b == 0 {
		return b
	}
	return greatestCommonFactor(a%b, b)
}

//二元樹遍歷
//二元樹結構不斷重複(也是規則重複)
