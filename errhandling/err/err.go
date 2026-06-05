package main

//错误处理
//路径问题，包导不进，解决方法：原来在defer里面藏了一个mod.go，删除这个文件就好了
import (
	"bufio"
	"errors"
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
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	err = errors.New("this is a custom error") //自己写error，error也是一个interface
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok { //判断error什么类型
			panic(err) //不知道什么类型
		} else {
			fmt.Printf("%s,%s,%s\n",
				pathError.Op,   //open
				pathError.Path, //fib.txt
				pathError.Err)  // file exists
		}
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file) //缓冲区
	defer writer.Flush()            //写入

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f()) //Fprintln会自动添加换行符
	}
}

func main() {
	// tryDefer()
	writeFile("fib.txt")
}
