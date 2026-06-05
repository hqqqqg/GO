// 表格驱动测试，所在直属目录下，go test 运行
package main

import "testing"

func TestTriangle(t *testing.T) { //TestXXX
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 37},
	}
	for _, tt := range tests {
		if actual := calcTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calcTriangle(%d,%d);"+"got %d; expected %d", tt.a, tt.b, actual, tt.c)
		}
	}
}
