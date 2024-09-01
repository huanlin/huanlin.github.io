---
title: 05 雜七雜八但是重要
tags: [Go]
---

整理一些重要或者比較 tricky 的 Go 語法或標準函式庫用法。

## String

留意 Unicode 的字串長度：

```go
unicodeCharStr := "地鼠"
fmt.Println(len(unicodeCharStr)) // output: 6
```

程式印出的結果是 6 而不是 2。這是因為 <mark>Go 的字串內部不是字元陣列，而是代表每個 UTF-8 字元的 byte 陣列。</mark>

因此，如果取出字串中的某個字元，不能以陣列索引的語法，否則結果不會是我們想要的：

```go
unicodeCharStr := "地鼠"
for i := 0; i < len(unicodeCharStr); i++ {
    fmt.Print(string(unicodeCharStr[i]) + " ")
}
fmt.Println() // 輸出:  å  ° é ¼
```

要取出字串中的字元，可以用 `range`：

```go
unicodeCharStr := "地鼠"
for i, r := range unicodeCharStr {
    fmt.Printf("%d:%s ", i, string(r))
}
fmt.Println() // 輸出: 0:地 3:鼠
```

## Blank identifier

呼叫函式時，如果某個回傳值無需處理，可以用一個 blank identifier 字元（ `_` ）來承接該回傳值。

範例：

```go
-, err = ReadFile("no/file)
if (err != nil) {
    fmt.Println("Error: err)
}
```

此範例所要表達的是：我不在乎 `ReadFile()` 執行成功時回傳的結果，而只看它是否返回錯誤。

## If with a short statement {#if-with-statement}

類似 `for` 迴圈，`if` 敘述也可以先有一個短敘述（short statement），然後才跟著判斷式。

範例：

```go
func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
        return v
    }
    return lim
}
```

第 2 行的意思是先把 `math.Pow()` 的結果指派給變數 `v`，然後判斷 `v` 是否小於 `lim`。

注意：由 `if` 的短敘述所宣告的變數只活在那個 `if` 區塊內。

## 取得物件或變數的型別 {#get-type}

這裡示範三種方法：

- 使用 fmt.Printf 的 %T 旗號。
- 使用 reflect 套件。
- 使用 type assertion。

### 使用 fmt.Printf 的 %T 旗號 {#printf-t-flag}

```go
var count int = 42
fmt.Printf("variable count=%v is of type %T \n", count, count)
```

### 使用 reflect 套件 {#reflect-package}

使用 `reflect.TypeOf()` 方法：

```go
fmt.Printf("%v", reflect.TypeOf(10))   // int
fmt.Printf("%v", reflect.TypeOf("Go")) // string
```

### 使用 type assertion {#type-assertion}

```go
var x interface{} = 7

switch x.(type) {
case int:
    fmt.Println("int")
}
```

參閱 A Tour of Go: [Type assertions](https://go.dev/tour/methods/15)

## `defer` 關鍵字 {#defer}

Go 的 `defer` 關鍵字可用來將函式的執行時機延後至包覆函式返回之前才執行，常用於清理資源（例如確保關閉資料庫連線）。

範例：

```go
func main() {
    defer fmt.Println("World") // 離開 main 函式之前才執行此敘述。
    fmt.Println("Hello")
}
```

輸出結果：

```text
Hello
World
```

### 清理資源 {#defer-cleanup}

範例：

```go
func doSomething() error {
  f, err := os.Open("test.txt")
  if err != nil {
    return err
  }
  defer f.Close()

  // 繼續處理檔案內容
}
```

注意：一旦檔案開啟成功，**接著立刻加上** `defer f.Close()`，然後才處理後續的檔案操作，如此便可確保此函式離開之前會關閉檔案。

### 後進先出 {#defer-lifo}

如果在一個函式中使用了多次 `defer`，那些被延遲的函式呼叫將會以後進先出的順序執行。

範例：

```go
func main() {
    defer fmt.Println(1)
    defer fmt.Println(2)
    defer fmt.Println(3)
}
```

輸出結果：

```text
3
2
1
```


[100-mistakes]: https://www.manning.com/books/100-go-mistakes-and-how-to-avoid-them
[go-in-action]: https://www.manning.com/books/go-in-action-second-edition
[go-in-practice]: https://www.manning.com/books/go-in-practice-second-edition
[go-by-example]: https://www.manning.com/books/go-by-example