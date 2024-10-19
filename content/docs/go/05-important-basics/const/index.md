---
title: 5.2 常數
---

建議閱讀官方文件：

- [Constant expressions](https://go.dev/ref/spec#Constant_expressions)
- [Constant declarations](https://go.dev/ref/spec#Constant_declarations)

特別值得留意的是關鍵字 `iota`，它能夠讓編譯器產生特定運算規則的常數值。以下是一些範例。

## Example 1

```go
const (
    Sunday =  0         // 0
    Monday =  iota + 1  // 1 (iota = 0)
    Tuesday             // 2 (iota = 1)
    Wednesday           // 3 (iota = 2)
    Thursday            // 4 (iota = 3)
    Friday              // 5 (iota = 4)
    Saturday            // 6 (iota = 5)
)
```

## Example 2

```go
const (
    Black      = iota + 1       // 1 = 0+1  (iota = 0)
    Red                         // 2 = 1+1  (iota = 1)
    Yellow                      // 3 = 2+1  (iota = 2)
    Green      = iota + 2       // 5 = 3+2  (iota = 3)
    Blue                        // 6 = 4+2  (iota = 4)
    Pink, Gray = iota, iota + 1 // 5, 6     (iota = 5)
    Brown      = iota + 3       // 9 = 6+3  (iota = 6)
    White                       // 10 = 7+3 (iota = 7)
)
```

**Try it:** <https://go.dev/play/p/F-XPYY3g2Hm>

## Example 3

定義常數時，也可以使用自訂型別。

```go
package main

import "fmt"

type language = string
type MyString string

const (
    EN language = "English"
    ZH          = "Chinese"
)

func main() {
    var lang language = "English"
    var str MyString = "Hello"
    fmt.Printf("Type of lang: %T\n", lang)
    fmt.Printf("Type of str: %T\n", str)

    fmt.Printf("Type of const EN: %T\n", EN)
}
```

執行結果：

```text
Type of lang: string
Type of str: main.MyString
Type of const EN: string
```

**Try it:** <https://go.dev/play/p/u2V08y0BmHq>
