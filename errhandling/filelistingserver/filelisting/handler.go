package filelisting

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

func HandleFileList(writer http.ResponseWriter, //输出，输入浏览器显示
	request *http.Request) error { //输入，用户请求信息
	fmt.Println()
	if strings.Index(
		request.URL.Path, prefix) != 0 { //找有没有/list/输出什么
		return userError("path must start" + "with" + prefix)
	}

	path := request.URL.Path[len("/list/"):] //拿到想读的文件
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := io.ReadAll(file)
	if err != nil {
		panic(err)

	}
	writer.Write(all) //写进writer里面
	return nil
}
