package apiDocPrinter

import (
	"bufio"
	"strings"
)

/**
    author : LinGuanHong
    github : https://github.com/af913337456
    blog   : http://www.cnblogs.com/linguanh
    time   : 16:49
*/

func getVal(val string,reader *bufio.Reader) string {
	val = strings.TrimSpace(val)
	if val != "" {
		return val
	}
	// 读下一行
	byt, _, err := reader.ReadLine()
	if err != nil {
		// 读完一个文件
		return ""
	}
	line := string(byt)
	if line == "" {
		return ""
	}
	return strings.TrimSpace(line)
}

func GetContents(targetTag string,reader *bufio.Reader) (string,[]string)  {
	ret := []string{}
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
		if isNewStartTag(line,targetTag) {
			return line,ret
		}
		ret = append(ret,line)
	}
	return "",ret
}












