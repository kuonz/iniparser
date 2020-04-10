### 功能

解析`.ini`配置文件



### API

##### func Parser(path string) (IniConfig, error)

功能：用于解析`ini`配置文件，获取解析结果集

参数path: ini配置文件路径

返回值IniConfig: 解析结果

返回值error: 解析过程是否出错



##### func (ic IniConfig) Get(session string, key string) (string, bool)

参数session: 节名

参数key: 键名

返回值string: 键对应的值

返回值bool: 是否存在键对应的值



##### func (ic IniConfig) GetDefault(key string) (string, bool)

参数key: 键

返回值string: 键对应的值

返回值bool: 是否存在键对应的值



### 示例

`demo.ini`

```ini
# 默认节内容
name=kuonz 
version=0.0.1

[mysql]
username=root
password=123456
```

`main.go`

```go
package main

import (
  "github.com/kuonz/iniparser"
  "fmt"
)

func main() {
  // 解析 demo.ini
  ic, err := iniparser.Parse("./demo.ini")

  if err != nil {
    panic(err.Error())
  }

  // 获取默认节内容
  if name, ok := ic.GetDefault("name"); !ok {
    fmt.Println("don't exist the key name")
  } else {
    fmt.Println(name)
  }

  if author, ok := ic.GetDefault("author"); !ok {
    fmt.Println("don't exist key author")
  } else {
    fmt.Println(author)
  }

  // 获取指定节内容
  if username, ok := ic.Get("mysql", "username"); !ok {
    fmt.Println("session mysql don't exist the key username")
  } else {
    fmt.Println(username)
  }
}
```