// 斐波那契数列
package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func fib() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

// 为函数实现接口
func (g intGen) Read(
	p []byte) (n int, err error) {
	next := g() //自己调用自己
	if next > 100 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)      //将数字变成字符串\n
	return strings.NewReader(s).Read(p) //将字符串读到p中
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() { //这行代码底层调用read
		fmt.Println(scanner.Text())
	}
}

func main() {
	var f intGen = fib()
	printFileContents(f)
}
