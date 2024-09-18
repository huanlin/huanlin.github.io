---
title: 5.5 init 函式
---

Go 有一個特殊用途的函式，名稱固定叫做 `init()`。此函式會在一個 package 載入時自動執行，故通常會把一些初始化的操作寫在此函式中。

特性與用法：

- `init` 函式不能有參數和回傳值。
- 每個 package 可以有多個 `init` 函式，而各 packages 的 `init` 函式的執行順序便是按照 packages 之間的依賴關係來決定。換言之，先載入哪個 package，就會先執行那個 package 的 `init` 函式。
- 每個 .go 檔案也可以有多個 `init` 函式，這些函式會按照它們在檔案中出現的順序執行。但一般來說，通常不會在一個 .go 檔案裡面寫多個 `init` 函式，以免讓程式碼更難理解和維護。
- 應避免在 `init` 函式中執行耗時工作。

範例：

```go
var greeting string

func init() {
    fmt.Println("Hello")
}

func main() {
    fmt.Println("world")
}
```

執行結果：

```text
Hello
world
```

### Using `init()` as side effects {#init-as-side-effects}

有時候，我們的程式不會呼叫某個套件裡面的任何函式或變數，但是卻需要應用程式載入時執行那個套件的 `init()` 函式，以便完成某些預設的配置或初始化操作。碰到這種情形，就必須在 import 該套件時使用 blank identifier (`_`)，像這樣：

```go
import (
    _ "github.com/lib/pq"
    _ "image/png"
    ...
)
```

這種在引用套件的名稱前面加一個底線字元的寫法稱為 importing side effects。換言之，那些套件的 `init()` 函式便是我們的應用程式所需要的「副作用」。

> [!note]
> 如果應用程式沒有直接使用某個套件裡面的任何東西，該套件名稱就不能出現在 `import` 區塊裡，除非前面加上底線字元 `_` 字元來告訴 Go 編譯器：我需要這個套件載入時自動執行的 `init()`。
