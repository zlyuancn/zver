# zver
> 版本解析模块

## 获取
> `go get -u -v github.com/zlyuancn/zver`

## 3级版本实例

```go
package main

import (
    "fmt"
    "github.com/zlyuancn/zver"
)

func main() {
    v1 := new(zver.Version3)
    if err := v1.Parser("3.4.2", "."); err != nil {
        panic(err)
    }

    fmt.Println(v1)
    fmt.Println(v1.ToText("_"))
    fmt.Println(v1.ToTextHasPrefix(".", "v"))

    v2 := new(zver.Version3)
    if err := v2.ParserHasPrefix("v3.4.3", ".", "v"); err != nil {
        panic(err)
    }

    fmt.Println(v1, ">", v2, ":", v1.Gt(v2))
    fmt.Println(v1, "<", v2, ":", v1.Le(v2))
    fmt.Println(v1, "=", v2, ":", v1.Eq(v2))
}
```
