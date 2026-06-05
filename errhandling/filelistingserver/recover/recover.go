//recover处理

package main

import (
	"fmt"
)

func tryRecover() {
	defer func() { //一定要在defer里面调用
		r := recover() //底层会把panic接住
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else { //不是error也不知道该怎么办
			// panic(r)
			panic(fmt.Sprintf(
				"I don't know what to do: %v", r))
		}
	}() //后面加括号
	// panic(errors.New("this is an error"))

	// b := 0
	// a := 5 / b
	// fmt.Println(a)

	panic(123)

}

func main() {
	tryRecover()
}
