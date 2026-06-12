// 走迷宫广度优先算法
package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col) //读取文件的第一行，把迷宫的行数和列数分别存进row和col

	maze := make([][]int, row) //创建一个外层切片，长度为行数row
	for i := range maze {      //遍历每一行
		maze[i] = make([]int, col) //为每一行创建一个内层切片，长度为列数col
		for j := range maze[i] {   //遍历这一行的每一列
			fmt.Fscanf(file, "%d", &maze[i][j]) //从文件里一次读取数字（0或1），填入二维数组的对应位置
		}
	}

	return maze //返回构建好的二维数组
}

// 定义“点”结构体，行坐标i和列坐标j
type point struct {
	i, j int
}

// 上左下右
var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

// 向量加法，计算移动后的新坐标
func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

// 探测网格是否越界
func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}

	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true //没越界返回真实值和true
}

func walk(maze [][]int,
	start, end point) [][]int {
	steps := make([][]int, len(maze)) //创建一个和迷宫一样大小的记步骤矩阵，初始为0
	for i := range steps {
		steps[i] = make([]int, len(maze[i])) //初始化记事本每一行
	}

	Q := []point{start} //创建一个队列Q，先把起点放进队列

	for len(Q) > 0 { //队列有点就一直循环
		cur := Q[0] //出队
		Q = Q[1:]   //更新队伍

		if cur == end { //如果当前位置刚好等于终点
			break
		}

		for _, dir := range dirs { //向 上 左 下 右 走
			next := cur.add(dir) //计算出走一步的新坐标

			val, ok := next.at(maze)
			if !ok || val == 1 { //新坐标是否撞到了墙
				continue //跳过
			}

			val, ok = next.at(steps)
			if !ok || val != 0 { //是否记录过
				continue
			}

			if next == start { //是否回到起点
				continue
			}
			//通过，合法的路
			curSteps, _ := cur.at(steps) //走到这里花了几步
			steps[next.i][next.j] =
				curSteps + 1 //走到新地块的步数

			Q = append(Q, next) //把新地块扔到队尾
		}
	}

	return steps //返回填满步数的记事本
}

func main() {
	maze := readMaze("maze/maze.in") //读文本文件

	steps := walk(maze, point{0, 0},
		point{len(maze) - 1, len(maze[0]) - 1}) //起点和终点

	for _, row := range steps { //遍历拿到结果矩阵的每一行
		for _, val := range row { //遍历每一行的每一个步数数字
			fmt.Printf("%3d", val) //打印，保证每个数字占3个字符宽度，对齐
		}
		fmt.Println()
	}

	// TODO: construct path from steps
}
