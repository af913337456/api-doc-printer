package apiDocPrinter

import "testing"

/**
    author : LinGuanHong
    github : https://github.com/af913337456
    blog   : http://www.cnblogs.com/linguanh
    time   : 12:31
*/

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