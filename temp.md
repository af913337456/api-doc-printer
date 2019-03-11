#### 创建发布合约订单
* 路由：`/CreateContractOrder`
* 请求方式：`post`
* 参数格式：`json`
* 令牌需要：`yes`
* 输入参数：
```prolog
Symbol  string `json:"symbol"`  // 币符号
Name    string `json:"name"`    // 币名称
Decimal int64  `json:"decimal"` // 小数位
Support string `json:"support"` // 发行量
```
