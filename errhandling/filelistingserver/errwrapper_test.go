// http测试，堆代码
// 视频里面http上下反了
package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func errPanic(_ http.ResponseWriter,
	_ *http.Request) error {
	panic(123)
}

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}

func errUserError(_ http.ResponseWriter,
	_ *http.Request) error {
	return testingUserError("user error")
}

func errNotFound(_ http.ResponseWriter,
	_ *http.Request) error {
	return os.ErrNotExist
}

func errNoPermission(_ http.ResponseWriter,
	_ *http.Request) error {
	return os.ErrPermission
}

func errUnknown(_ http.ResponseWriter,
	_ *http.Request) error {
	return errors.New("unknown error")
}

func noError(writer http.ResponseWriter,
	_ *http.Request) error {
	fmt.Fprintln(writer, "no error")
	return nil
}

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNotFound, 404, "Not Found"},
	{errNoPermission, 403, "Forbidden"},
	{errUnknown, 500, "Internal Server Error"},
	{noError, 200, "no error"},
}

// 假的request/response
func TestErrWrapper(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)              //就测这个函数
		response := httptest.NewRecorder() //返回*httptest.ResponseRecorder结构体，实现http.ResponseWriter接口，在web.go里面
		request := httptest.NewRequest(    //假装客户发来的
			http.MethodGet,              //请求方法
			"http://www.imooc.com", nil) //网址和body
		f(response, request)

		verifyResponse(response.Result(),
			tt.code, tt.message, t)
	}
}

// 起整个服务器，测试的覆盖率大
func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		server := httptest.NewServer( //自动分配空闲端口
			http.HandlerFunc(f)) //套上类型转换，包装成合法接口
		resp, _ := http.Get(server.URL) //客户端

		verifyResponse(
			resp, tt.code, tt.message, t)
	}
}

func verifyResponse(resp *http.Response, //response
	expectedCode int, //预期的正确状态码
	expectedMsg string, //预期的正确报错文本
	t *testing.T) {
	b, _ := io.ReadAll(resp.Body)         //忽略error
	body := strings.Trim(string(b), "\n") //要对比所以要修剪
	if resp.StatusCode != expectedCode ||
		body != expectedMsg { //状态码不对或者真实的文本内容不对
		t.Errorf("expect (%d, %s); "+ //打印
			"got (%d, %s)",
			expectedCode, expectedMsg,
			resp.StatusCode, body)
	}
}
