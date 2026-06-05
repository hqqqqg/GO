// 代码覆盖率go test -coverprofile=c.out
// 查看能看懂的东西 go tool cover -html=c.out
package main

import "testing"

func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		// 普通的测试
		{"abcabcbb", 3},
		{"pwwkew", 3},

		// 特殊的
		{"", 0},
		{"b", 1},
		{"bbbbbbbbb", 1},
		{"abcabcabcd", 4},

		// 中文
		{"这里是慕课网", 6},
		{"一二三二一", 3},
		{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8},
	}

	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for input %s; "+
				"expected %d",
				actual, tt.s, tt.ans)
		}
	}
}

// 性能测试go test -bench . 记得加点！！
// 生成文件 go test -bench. -cpuprofile cpu.out
// 使用工具进入另一个环境 go tool pprof cpu.out
// 在pprof里用web画图看性能消耗,画出来的图看不到全貌！！！！！
func BenchmarkSubstr(b *testing.B) { // BenchmarkXxx
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	for i := 0; i < 13; i++ {
		s = s + s
	}
	b.Logf("len(s) = %d", len(s)) //测试用来打印
	ans := 8
	b.ResetTimer() //重置时间

	for i := 0; i < b.N; i++ { //b.N是GO框架决定函数需要重复的次数
		actual := lengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("got %d for input %s; "+
				"expected %d",
				actual, s, ans)
		}
	}
}
