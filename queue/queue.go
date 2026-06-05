// 生成文档
// go doc 查看当前包的概览
// go doc Queue 查看包以及函数
// go doc Queue.IsEmpty 查看这个函数
// go help doc
// go doc fmt.Println
// pkgsite -http :6060  godoc弃用，用这个新的东西找不到我的文档
package queue

// 1
type Queue []int

// 2
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

// 3
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

// 4
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
