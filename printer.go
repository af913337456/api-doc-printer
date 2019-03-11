package apiDocPrinter

import (
	"os"
	"fmt"
	"time"
)

/**
    author : LinGuanHong
    github : https://github.com/af913337456
    blog   : http://www.cnblogs.com/linguanh
    time   : 12:29
*/

func PrintToMD(data string)  {
	fileName := fmt.Sprintf("temp_%d.md",time.Now().Unix())
	file,err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	file.WriteString(data)
}