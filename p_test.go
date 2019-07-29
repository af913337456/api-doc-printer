package apiDocPrinter

import (
	"testing"
	"strings"
	"fmt"
)

/**
    author : LinGuanHong
    github : https://github.com/af913337456
    blog   : http://www.cnblogs.com/linguanh
    time   : 12:31
*/

func TestSplit(t *testing.T) {
	str := "name:"
	arr := strings.Split(str,":")
	fmt.Println(arr[1])
}

func Test_Printer(t *testing.T) {
	PrintToMD(temp_1)
}

func Test_ReadCodeFile(t *testing.T) {
	list := ReadCodeFile("C:/go_1.11/lgh/src/NormalXGoServerApi/api/account/order.go")
	lines := ""
	for _,item := range list {
		lines = lines + FormatTemp1(item) + "\n"
	}
	PrintToMD(lines)
}

func Test_ReadCodeFile_2(t *testing.T) {
	path := "C:/Go-1.12/lgh/src/NormalXGoServerApi/api/uploader/qiniuUploader.go"
	list := ReadCodeFile(path)
	lines := ""
	for _,item := range list {
		lines = lines + FormatTemp1(item) + "\n"
	}
	PrintToMD(lines)
}

