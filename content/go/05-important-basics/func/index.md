---
title: 5.6 函式
---

## 基礎語法 {#func-basic}

- 呼叫同一套件的函式，只要寫函式名稱。
- 呼叫其他套件提供的（exported）函式，則須使用 `<package>.<function>` 語法。

範例：

```go
// 定義一個函式，名稱為 hello，沒有傳入參數，也沒有回傳值。
func hello() {
    fmt.Println("Hello, world!")
}

func main() {
    hello() // 呼叫 hello 函式。
}
```

## 函式與回傳值 {#writing-func}

本節內容：

- 撰寫具名函式。
- 撰寫匿名函式。

### 具名函式  {#named-func}

以下兩種寫法都可以：

```go
func parse1(filename string) ([]string, []byte, error )
func parse2(filename string) ([]string headers, []byte body, err error)
```

兩者差異在於回傳值的寫法：`parse1` 沒有預先宣告回傳的參數名稱，`parse2` 則有。

《100 Go Mistakes and How to Avoid Theme》書中的建議（#43: Never using named result parameters）如下：

- 如果不至於令程式碼不好理解，應避免使用具名的回傳參數。
- 有些情況，例如介面方法，或者函式的回傳參數有兩個以上的參數型別相同，使得閱讀程式碼的時候不容易辨識，則建議使用具名的回傳參數。

簡單來說，要不要使用具名的回傳參數，應該以程式碼的可讀性為主要考量。

值得留意的是，採用具名回傳參數的寫法時，那些回傳參數即等同於一進入函式就初始為 0 或 `nil` 的區域變數。所以函式返回時的 `return` 敘述可以不用寫回傳參數，像這樣：

```go
func foo(a int) (b int)
{
    b = a
    return // naked return
}
```

這種寫法叫做 **naked return**。

<mark>Naked return 的寫法應該只用於簡短的函式，因為 `return` 時沒有寫明回傳的參數，程式碼比較不好理解。</mark>至於多簡短的函式才比較能讓人接受 naked return 的寫法，則沒有標準，大原則是以程式碼是否容易閱讀和維護來決定。

> 如果一定要有一個具體準則，這裡給出一個參考依據：當一個函式的程式碼超出編輯器的顯示區域，而必須上下捲動才能看到完整的程式碼，則此函式可能不易「一看就懂」，故應避免使用具名的回傳參數和 naked return 語法。久而久之，可能自然而然就不會使用 naked return 了。

#### 具名回傳參數的副作用 {#side-effects-with-named-result-params}

使用具名回傳參數時，可能一不留神就寫出類似底下的程式碼：

```go
func parse(ctx AppContext) (title string, err error)
{
    if ctx.CheckError() != nil {
        return "", err
    }

    // ...
}
```

上面的程式碼可以通過編譯，但是有一個 bug：永遠不會回傳錯誤，因為 `err` 在一進入函式的時候就已經初始為 `nil` 了。底下是正確的寫法：

```go
func parse(ctx AppContext) (title string, err error)
{
    if err := ctx.CheckError(); err != nil {
        return "", err
    }

    // ...
}
```

### 匿名函式 {#anonymous-func}

```go
func main() {
    // 撰寫匿名函式，並指派給變數 f
    f := func() {
        fmt.Println("Inside a function")
    }

    f()  // 呼叫 f 指向的函式
}
```

## init 函式

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
