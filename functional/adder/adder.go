// 闭包=局部函数

package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x //x就是传进来的i,sum在函数内一直被用着，直到main函数结束，sum才会被销毁
		return sum
	}
}

func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(a(i))
	}
}
