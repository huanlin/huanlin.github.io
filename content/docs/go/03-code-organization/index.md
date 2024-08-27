---
title: 03 Code organization
tags: [Go]
---

## Scope

程式裡面有許多變數、函式、型別等識別字，依照它們宣告時的所在位置和寫法，可分為三種可見範圍：

- block：宣告在 `{...}` 區塊裡面的變數只有該區塊的程式碼可存取。
- package：同一個 package 內的 .go 程式檔案可存取彼此的變數（以及函式、型別等等），無論它們是否為 exported（公開成員）。
- global：只要是 exported 變數（名稱以大寫英文開頭來命名的都是），就能夠被任何程式碼存取。

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


[100-mistakes]: https://www.manning.com/books/100-go-mistakes-and-how-to-avoid-them