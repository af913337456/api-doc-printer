package apiDocPrinter

import (
	"os"
	"bufio"
	"strings"
	"fmt"
	"encoding/json"
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
	var item Temp1Obj
	reader := bufio.NewReader(file)
	begin := false
	objList := []Temp1Obj{}
	endLine := ""
	for {
		var line = ""
		if endLine != "" {
			line = endLine
		}else{
			byt, _, err := reader.ReadLine()
			if err != nil {
				// 读完一个文件
				break
			}
			line = string(byt)
		}
		if line == "" {
			continue
		}
		line = strings.TrimSpace(line)
		if begin {
			fmt.Println(line)
			if line == "*/" {
				begin = false
				bys,_ := json.Marshal(item)
				fmt.Println(string(bys))
				objList = append(objList,item)
				endLine = ""
				item = Temp1Obj{}
			}else{
				tag,val := ParseTag(line)
				switch tag {
				case NameTag:
					item.Name = getVal(val,reader)
					break
				case RouterTag:
					item.Router = getVal(val,reader)
					break
				case HttpTag:
					item.ReqWay = getVal(val,reader)
					break
				case FormatTag:
					item.Format = getVal(val,reader)
					break
				case TokenTag:
					item.Token = getVal(val,reader)
					break
				case EndTag:
					break
				case InputJsonTag:
					endLine,item.JsonInputLines = GetContents(InputJsonTag,reader)
					break
				case InputXmlTag:
					endLine,item.XmlInputLines  = GetContents(InputXmlTag,reader)
					break
				case InputTextTag:
					endLine,item.TextInputLines = GetContents(InputTextTag,reader)
					break
				case InputDesTag:
					endLine,item.InputDesLines = GetContents(InputDesTag,reader)
					break
				case OutputJsonTag:
					endLine,item.JsonOutputLines = GetContents(OutputJsonTag,reader)
					break
				case OutputXmlTag:
					endLine,item.XmlOutputLines = GetContents(OutputXmlTag,reader)
					break
				case OutputTextTag:
					endLine,item.TextOutputLines = GetContents(OutputTextTag,reader)
					break
				case OutputDesTag:
					endLine,item.OutputDesLines = GetContents(OutputDesTag,reader)
					break
				default:

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















