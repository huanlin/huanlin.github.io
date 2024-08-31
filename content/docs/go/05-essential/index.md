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

## Variable shadowing

以下範例程式可以編譯和執行，但寫法容易令人 confuse：

```go
var case1 bool = true
var sum int = 100

func main() {
    if case1 {
        sum := add(5, 5) // 區域變數
        fmt.Println(sum)
    } else {
        m := add(10, 10) // 區域變數
        fmt.Println(sum)
    }

    fmt.Println(sum) // 使用全域變數
}

func add(x, y int) int {
    return x + y
}
```

程式中有幾處 `sum` 變數，有的是全域變數，有的是區域變數。雖然能通過編譯，但人眼容易誤讀，因為 `:=` 運算子可以同時宣告變數且賦值，使其左側的變數成為區域變數。如果使用 `=` 運算子，則會使用先前宣告過的變數，在此範例便是全域的 `sum`。

參見：[100 Go Mistakes and How to Avoid Them][100-mistakes] 的第 1 條：Unintended variable shadowing。

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


[100-mistakes]: https://www.manning.com/books/100-go-mistakes-and-how-to-avoid-them
[go-in-action]: https://www.manning.com/books/go-in-action-second-edition
[go-in-practice]: https://www.manning.com/books/go-in-practice-second-edition
[go-by-example]: https://www.manning.com/books/go-by-example