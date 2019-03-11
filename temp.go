package apiDocPrinter

import (
	"fmt"
	"strings"
	"bufio"
)

/**
    author : LinGuanHong
    github : https://github.com/af913337456
    blog   : http://www.cnblogs.com/linguanh
    time   : 11:53
*/

const temp_1 =
	"" +
	"#### %s\n" +
	"* 路由：`/%s`\n" +
	"* 请求方式：`%s`\n" +
	"* 参数格式：`%s`\n" +
	"* 令牌需要：`%s`\n" +
	"* 输入参数：\n" +
	"```prolog\n" +
	"%s"+
	"```"

type Temp1Obj struct {
	Name   string `json:"name"`
	Router string `json:"router"`
	ReqWay string `json:"req_way"`
	Format string `json:"format"`
	Token  string `json:"token"`
	Lines  []string `json:"lines"`
}

func FormatTemp1(obj Temp1Obj) string {
	params := ""
	for _,item := range obj.Lines {
		params = params + item + "\n"
	}
	if params == "" {
		params = "暂无"
	}
	return fmt.Sprintf(temp_1,obj.Name,obj.Router,obj.ReqWay,obj.Format,obj.Token,params)
}

/*
name:创建发布合约订单
router:CreateContractOrder
http:post
format:json
token:yes
input:
Symbol  string `json:"symbol"`  // 币符号
Name    string `json:"name"`    // 币名称
Decimal int64  `json:"decimal"` // 小数位
Support string `json:"support"` // 发行量
end:
*/

func GetName(line string) string {
	return getObj(line,"name")
}

func GetRouter(line string) string {
	return getObj(line,"router")
}

func GetReqWay(line string) string {
	return getObj(line,"http")
}

func GetFormat(line string) string {
	return getObj(line,"format")
}

func GetToken(line string) string {
	return getObj(line,"token")
}

func GetInput(line string,reader *bufio.Reader) []string {
	if !strings.Contains(line,"input:") {
		return nil
	}
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
		if strings.Contains(line,"end:") {
			break
		}
		ret = append(ret,line)
	}
	return ret
}

func getObj(line,tag string) string {
	if strings.Contains(line,tag+":") {
		arr := strings.Split(line,":")
		if len(arr) < 2 {
			return ""
		}
		return arr[1]
	}
	return ""
}

