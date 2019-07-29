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
	"* 路由：`%s`\n" +
	"* 请求方式：`%s`\n" +
	"* 参数格式：`%s`\n" +
	"* 令牌需要：`%s`\n" +
	"* 输入参数：\n" +
	"```prolog\n" +
	"%s\n"+
	"```\n"+
	"* 输入参数描述：\n" +
	"```prolog\n" +
	"%s\n"+
	"```\n"+
	"* 输出：\n" +
	"```prolog\n" +
	"%s\n"+
	"```\n"+
	"* 输出字段描述：\n" +
	"```prolog\n" +
	"%s\n"+
	"```\n"

const (
	NameTag = "name:"
	RouterTag = "router:"
	HttpTag = "http:"
	FormatTag = "format:"
	TokenTag  = "token:"

	EndTag = "*/"
	InputDesTag  = "input-des:"
	OutputDesTag = "output-des:"
	InputJsonTag = "input-json:"

	InputXmlTag  = "input-xml:"
	InputTextTag = "input-text:"
	OutputJsonTag = "output-json:"
	OutputXmlTag  = "output-xml:"
	OutputTextTag = "output-text:"

	EmptyContent = "暂无"
)

var (
	AllTags = []string{
		NameTag,RouterTag,HttpTag,FormatTag,EndTag,
		TokenTag,InputDesTag,OutputDesTag,InputJsonTag,
		InputXmlTag,InputTextTag,OutputJsonTag,OutputXmlTag,OutputTextTag}

	OutputTags = []string{OutputJsonTag,OutputXmlTag,OutputTextTag}
)

type Temp1Obj struct {
	Name   string `json:"name"`
	Router string `json:"router"`
	ReqWay string `json:"req_way"`
	Format string `json:"format"`
	Token  string `json:"token"`

	JsonInputLines  []string `json:"json_input_lines"`
	XmlInputLines   []string `json:"xml_input_lines"`
	TextInputLines  []string `json:"text_input_lines"`
	InputDesLines   []string `json:"input_des_lines"`
	
	JsonOutputLines []string `json:"json_output_lines"`
	XmlOutputLines  []string `json:"xml_output_lines"`
	TextOutputLines []string `json:"text_output_lines"`
	OutputDesLines  []string `json:"output_des_lines"`
}

func ParseTag(line string) (string,string) {
	for _,item := range AllTags {
		if strings.Contains(line,item) {
			arr := strings.Split(line,":")
			if len(arr) < 2 {
				return "",""
			}
			return arr[0]+":",arr[1]
		}
	}
	return "",""
}

func isOutputTag(line,targetTag string) bool {
	line = strings.TrimSpace(line)
	for _,item := range OutputTags {
		if item == line && item != targetTag {
			return true
		}
	}
	return false
}

func isNewStartTag(line,targetTag string) bool {
	line = strings.TrimSpace(line)
	if line == EndTag && targetTag != EndTag {
		return true
	}
	if line == InputDesTag && targetTag != InputDesTag {
		return true
	}
	if isOutputTag(line,targetTag) {
		return true
	}
	if line == OutputDesTag && targetTag != OutputDesTag {
		return true
	}
	return false
}

func FormatTemp1(obj Temp1Obj) string {
	getParams := func(inputs []string) string {
		params := ""
		size := len(inputs)
		for i:=0;i<size;i++ {
			params = params + inputs[i]
			if i != size - 1 {
				params = params + "\n"
			}
		}
		if params == "" {
			params = EmptyContent
		}
		return params
	}
	inputs := ""
	inputDes := ""
	outputs := ""
	outputDes := ""
	if content := getParams(obj.JsonInputLines);content != EmptyContent {
		inputs = content
	}
	if content := getParams(obj.XmlInputLines);content != EmptyContent {
		inputs = content
	}
	if content := getParams(obj.TextInputLines);content != EmptyContent {
		inputs = content
	}
	if content := getParams(obj.InputDesLines);content != EmptyContent {
		inputDes = content
	}

	if content := getParams(obj.JsonOutputLines);content != EmptyContent {
		outputs = content
	}
	if content := getParams(obj.XmlOutputLines);content != EmptyContent {
		outputs = content
	}
	if content := getParams(obj.TextOutputLines);content != EmptyContent {
		outputs = content
	}
	if content := getParams(obj.OutputDesLines);content != EmptyContent {
		outputDes = content
	}
	return fmt.Sprintf(
		temp_1,
		obj.Name,obj.Router,obj.ReqWay,obj.Format,obj.Token,
		inputs,inputDes,outputs,outputDes)
}

/*
name:创建发布合约订单
router:CreateContractOrder
http:post
format:json
token:yes
input-json:
Symbol  string `json:"symbol"`  // 币符号
Name    string `json:"name"`    // 币名称
Decimal int64  `json:"decimal"` // 小数位
Support string `json:"support"` // 发行量
end:
*/

func GetName(line string,reader *bufio.Reader) string {
	return getObj(line,NameTag,reader)
}

func GetRouter(line string,reader *bufio.Reader) string {
	return getObj(line,RouterTag,reader)
}

func GetReqWay(line string,reader *bufio.Reader) string {
	return getObj(line,HttpTag,reader)
}

func GetFormat(line string,reader *bufio.Reader) string {
	return getObj(line,FormatTag,reader)
}

func GetToken(line string,reader *bufio.Reader) string {
	return getObj(line,TokenTag,reader)
}

func GetInputJson(line string,reader *bufio.Reader) (string,[]string) {
	return getInputs(line,InputJsonTag,reader)
}

func GetInputDes(line string,reader *bufio.Reader) (string,[]string) {
	return getInputs(line,InputDesTag,reader)
}

func GetInputXml(line string,reader *bufio.Reader) (string,[]string) {
	return getInputs(line,InputXmlTag,reader)
}

func GetInputText(line string,reader *bufio.Reader) (string,[]string) {
	return getInputs(line,InputTextTag,reader)
}

func getInputs(line,targetTag string,reader *bufio.Reader) (string,[]string)  {
	if !strings.Contains(line,targetTag) {
		return "",nil
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
		if isNewStartTag(line,targetTag) {
			return line,ret
		}
		ret = append(ret,line)
	}
	return line,ret
}

func getObj(line,tag string,reader *bufio.Reader) string {
	if strings.Contains(line,tag+":") {
		arr := strings.Split(line,":")
		if len(arr) < 2 {
			return ""
		}
		if arr[1] == "" {
			// 后面是空字符串，尝试获取下一行
			byt, _, err := reader.ReadLine()
			if err != nil {
				// 读完一个文件
				return ""
			}
			line := string(byt)
			if line == "" {
				return ""
			}
			return line
		}
		return arr[1]
	}
	return ""
}

