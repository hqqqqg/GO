// 斐波那契数列
package fib

func Fibonacci() func() int { //返回的是函数
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
