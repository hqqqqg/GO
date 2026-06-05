package main

//路径问题，包导不进，解决方法：原来在defer里面藏了一个mod.go，删除这个文件就好了
import (
	"bufio"
	"fmt"
	"google/errhandling/defer/fib"
	"os"
)

func tryDefer() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("oops")
		}
	}
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()              //记得关上
	writer := bufio.NewWriter(file) //缓冲区
	defer writer.Flush()            //写入

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	// tryDefer()
	writeFile("fib.txt")
}
