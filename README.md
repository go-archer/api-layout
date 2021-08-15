# API项目布局
## 目录结构
```bash
.
├── config      -配置文件目录
├── service     -业务逻辑实现层，关联实体层数据，使用数据操作层实现业务逻辑
├── entity      -数据实体层，包含请求、响应数据定义
├── middleware  -中间件
├── router      -路由配置
├── controller  -控制器
├── docs        -项目文档，通过swag init 生成
├── repository  -数据操作层，关联模型层数据
├── model       -数据模型层
├── main.go    
├── go.mod
├── go.sum
└── README.md
```
## 安装文档插件
```bash
go get -u github.com/swaggo/swag/cmd/swag
```
## 请求验证使用说明
### 自定义验证

```go
// 结构体
func customFunc(fl validator.FieldLevel) bool {
    if fl.Field().String() == "invalid" {
        return false
    }
    return true
}
validate.RegisterValidation("custom tag name", customFunc)
//　注意：使用与现有函数相同的标记名称将覆盖现有函数
```

### 字段说明

| 字段                          | 说明                                                                                                                                                                     |
| ----------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| -                             | 忽略字段，告诉验证跳过这个struct字段                                                                                                                                     |
| &#124;                        | 'or'运算符，允许使用和接受多个验证器                                                                                                                                     |
| structonly                    | 当遇到嵌套结构的字段并包含此标志时，将运行嵌套结构上的任何验证，但不会验                                                                                                 |
| nostructlevel                 | 与structonly标记相同，但不会运行任何结构级别验证                                                                                                                         |
| omitempty                     | 允许条件验证                                                                                                                                                             |
| dive                          | 告诉验证者潜入切片，数组或映射，并使用后面的验证标记验证切片，数组或映射的该级别                                                                                         |
| required                      | 验证该值不是数据类型的默认零值。数字不为０，字符串不为 " ", slices, maps, pointers, interfaces, channels and functions不为nil                                            |
| isdefault                     | 验证该值是默认值                                                                                                                                                         |
| len=10                        | 对于数字，长度将确保该值等于给定的参数。对于字符串，它会检查字符串长度是否与字符数完全相同。对于切片，数组和map，验证元素个数                                            |
| max=10                        | 对于数字，max将确保该值小于或等于给定的参数。对于字符串，它会检查字符串长度是否最多为该字符数。对于切片，数组和map，验证元素个数                                         |
| min=10                        | 与max相反                                                                                                                                                                |
| eq=10                         | 对于字符串和数字，eq将确保该值等于给定的参数。对于切片，数组和map，验证元素个数                                                                                          |
| ne=10                         | 与eq相反                                                                                                                                                                 |
| oneof=red green (oneof=5 7 9) | 对于字符串，整数和uint，oneof将确保该值是参数中的值之一。参数应该是由空格分隔的值列表。值可以是字符串或数字                                                              |
| gt=10                         | 大于，对于数字，确保该值大于给定的参数。对于字符串，它会检查字符串长度是否大于该字符数。对于切片，数组和map，它会验证元素个数。对于time.Time确保时间值大于time.Now.UTC() |
| gte=10                        | 大于等于，对于time.Time确保时间值大于或等于time.Now.UTC()                                                                                                                |
| lt=10                         | 小于，对于time.Time确保时间值小于time.Now.UTC()                                                                                                                          |
| lte=10                        | 小于等于，对于time.Time确保时间值小于等于time.Now.UTC()                                                                                                                  |
| unique                        | 对于数组和切片，unique将确保没有重复项。对于map，unique将确保没有重复值                                                                                                  |
| alpha                         | 验证字符串值是否仅包含ASCII字母字符                                                                                                                                      |
| alphanum                      | 验证字符串值是否仅包含ASCII字母数字字符                                                                                                                                  |
| alphaunicode                  | 验证字符串值是否仅包含unicode字符                                                                                                                                        |
| alphanumunicode               | 验证字符串值是否仅包含unicode字母数字字符                                                                                                                                |
| numeric                       | 验证字符串值是否包含基本数值。基本排除指数等...对于整数或浮点数，返回true                                                                                                |
| hexadecimal                   | 验证字符串值是否包含有效的十六进制                                                                                                                                       |
| hexcolor                      | 验证字符串值包含有效的十六进制颜色，包括＃标签                                                                                                                           |
| rgb                           | 验证字符串值是否包含有效的rgb颜色                                                                                                                                        |
| rgba                          | 验证字符串值是否包含有效的rgba颜色                                                                                                                                       |
| hsl                           | 验证字符串值是否包含有效的hsl颜色                                                                                                                                        |
| hsla                          | 验证字符串值是否包含有效的hsla颜色                                                                                                                                       |
| email                         | 验证字符串值包含有效的电子邮件                                                                                                                                           |
| file                          | 验证字符串值是否包含有效的文件路径，并且该文件存在于计算机上                                                                                                             |
| url                           | 验证字符串值是否包含有效的url接受golang请求uri接受的任何url，但必须包含一个模式，例如http://或rtmp://                                                                    |
| uri                           | 验证了字符串值包含有效的uri。接受uri接受的golang请求的任何uri                                                                                                            |
| base64                        | 验证字符串值是否包含有效的base64值                                                                                                                                       |
| base64url                     | 根据RFC4648规范验证字符串值是否包含有效的base64 URL安全值                                                                                                                |
| btc_addr                      | 验证字符串值是否包含有效的比特币地址                                                                                                                                     |
| btc_addr_bech32               | 验证了字符串值包含bip-0173定义的有效比特币Bech32地址                                                                                                                     |
| eth_addr                      | 验证字符串值是否包含有效的以太坊地址                                                                                                                                     |
| contains=@                    | 验证字符串值是否包含子字符串值                                                                                                                                           |
| containsany=!@#?              | 验证字符串值是否包含子字符串值中的任何Unicode code points                                                                                                                |
| containsrune=@                | 验证字符串值是否包含提供的符文值                                                                                                                                         |
| excludes=@                    | 验证字符串值不包含子字符串值                                                                                                                                             |
| excludesall=!@#?              | 将验证字符串值在子字符串值中是否包含任何Unicode code points                                                                                                              |
| excludesrune=@                | 验证字符串值是否包含提供的符文值                                                                                                                                         |
| isbn                          | 验证字符串值是否包含有效的isbn10或isbn13值。                                                                                                                             |
| isbn10                        | 验证字符串值是否包含有效的isbn10值。                                                                                                                                     |
| isbn13                        | 验证字符串值是否包含有效的isbn13值。                                                                                                                                     |
| uuid                          | 验证字符串值是否包含有效的UUID。                                                                                                                                         |
| uuid3                         | 验证字符串值是否包含有效的版本3 UUID。                                                                                                                                   |
| uuid4                         | 验证字符串值是否包含有效的版本4 UUID。                                                                                                                                   |
| uuid5                         | 验证字符串值是否包含有效的版本5 UUID。                                                                                                                                   |
| ascii                         | 验证字符串值是否仅包含ASCII字符。注意：如果字符串为空，则验证为true                                                                                                      |
| printascii                    | 验证字符串值是否仅包含可打印的ASCII字符。注意：如果字符串为空，则验证为true。                                                                                            |
| multibyte                     | 验证字符串值是否包含一个或多个多字节字符。注意：如果字符串为空，则验证为true                                                                                             |
| datauri                       | 验证字符串值是否包含有效的DataURI。注意：这也将验证数据部分是否有效base64                                                                                                |
| latitude                      | 验证字符串值是否包含有效的纬度。                                                                                                                                         |
| longitude                     | 验证字符串值是否包含有效经度。                                                                                                                                           |
| ssn                           | 验证字符串值是否包含有效的美国社会安全号码。                                                                                                                             |
| ip                            | 验证字符串值是否包含有效的IP地址                                                                                                                                         |
| ipv4                          | 验证字符串值是否包含有效的v4 IP地址                                                                                                                                      |
| ipv6                          | 验证字符串值是否包含有效的v6 IP地址                                                                                                                                      |
| cidr                          | 验证字符串值是否包含有效的CIDR地址                                                                                                                                       |
| cidrv4                        | 验证字符串值是否包含有效的v4 CIDR地址                                                                                                                                    |
| cidrv5                        | 验证字符串值是否包含有效的v5 CIDR地址                                                                                                                                    |
| tcp_addr                      | 验证字符串值是否包含有效的可解析TCP地址                                                                                                                                  |
| tcp4_addr                     | 验证字符串值是否包含有效的可解析v4 TCP地址                                                                                                                               |
| tcp6_addr                     | 验证字符串值是否包含有效的可解析v6 TCP地址                                                                                                                               |
| udp_addr                      | 验证字符串值是否包含有效的可解析UDP地址                                                                                                                                  |
| udp4_addr                     | 验证字符串值是否包含有效的可解析v4 UDP地址                                                                                                                               |
| udp6_addr                     | 验证字符串值是否包含有效的可解析v6 UDP地址                                                                                                                               |
| ip_addr                       | 验证字符串值是否包含有效的可解析IP地址                                                                                                                                   |
| ip4_addr                      | 验证字符串值是否包含有效的可解析v4 IP地址                                                                                                                                |
| ip6_addr                      | 验证字符串值是否包含有效的可解析v6 IP地址                                                                                                                                |
| unix_addr                     | 验证字符串值是否包含有效的Unix地址                                                                                                                                       |
| mac                           | 验证字符串值是否包含有效的MAC地址。注意：有关可接受的格式和类型，请参阅Go的ParseMAC: http://golang.org/src/net/mac.go?s=866:918#L29                                      |
| hostname                      | 根据RFC 952 https://tools.ietf.org/html/rfc952验证字符串值是否为有效主机名                                                                                               |
| fqdn                          | 验证字符串值是否包含有效的FQDN (完全合格的有效域名)，Full Qualified Domain Name (FQDN)                                                                                   |
| html                          | 验证字符串值是否为HTML元素标记，包括https://developer.mozilla.org/en-US/docs/Web/HTML/Element中描述的标记。                                                              |
| html_encoded                  | 验证字符串值是十进制或十六进制格式的正确字符引用                                                                                                                         |
| url_encoded                   | 这验证了根据https://tools.ietf.org/html/rfc3986#section-2.1对字符串值进行了百分比编码（URL编码）                                                                         |

