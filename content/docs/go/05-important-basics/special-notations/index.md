---
title: 5.3 特殊表示法
---

## Blank identifier ( `_` ) {#blank-identifier}

呼叫函式時，如果某個回傳值無需處理，可以用一個 blank identifier 字元，也就是底線（ `_` ）來承接該回傳值。

範例：

```go
-, err = ReadFile("no/file)
if (err != nil) {
    fmt.Println("Error: err)
}
```

此範例所要表達的是：我不在乎 `ReadFile()` 執行成功時回傳的結果，而只看它是否返回錯誤。

## 省略符號 ( `...` ) {#ellipsis}

在 Go 語言中，`...` 是一個特殊語法，主要用於兩個場合：

- 函式的可變參數（variadic parameters），也就是可以傳入任意數量的參數。
- 用於展開 slice。

### 不定個數的參數 {#variadic-params}

宣告函式的參數時，可以使用 `...` 來表示可接受任意數量的參數。這些參數會被視為一個 `slice`。

範例：

```go
// 定義一個可變參數函式
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

func main() {
    fmt.Println(sum(1, 2, 3)) // 輸出: 6
    fmt.Println(sum(4, 5, 6, 7)) // 輸出: 22
}
```

### 展開 slice

範例：

```go
func printNumbers(nums ...int) {
    for _, num := range nums {
        fmt.Println(num)
    }
}

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    printNumbers(numbers...) // 展開 slice 並傳遞給函式
}
```

