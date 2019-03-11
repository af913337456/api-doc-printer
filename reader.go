package apiDocPrinter

import (
	"os"
	"bufio"
	"strings"
)

/**
    author : LinGuanHong
    github : https://github.com/af913337456
    blog   : http://www.cnblogs.com/linguanh
    time   : 12:41
*/

func ReadCodeFile(fileName string) []Temp1Obj {
	file,err := os.OpenFile(fileName,os.O_RDONLY,0777)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	begin := false
	var item Temp1Obj
	objList := []Temp1Obj{}
	for {
		byt, _, err := reader.ReadLine()
		if err != nil {
			// 读完一个文件
			break
		}
		line := string(byt)
		if line == "" {
			continue
		}
		line = strings.TrimSpace(line)
		if begin {
			if line == "*/" {
				begin = false
			}else{
				if name := GetName(line);name != "" {
					item.Name = name
				}else if reqWay := GetReqWay(line);reqWay != "" {
					item.ReqWay = reqWay
				}else if format := GetFormat(line);format != "" {
					item.Format = format
				}else if router := GetRouter(line);router != "" {
					item.Router = router
				}else if token  := GetToken(line);token != "" {
					item.Token = token
				}else if inputs := GetInput(line,reader);inputs != nil {
					item.Lines = inputs
					objList = append(objList,item)
				}
			}
			continue
		}
		if line == "/*" || line == "/**" {
			// 开始
			begin = true
			item = Temp1Obj{}
		}else{
			begin = false
		}
	}
	return objList
}

















