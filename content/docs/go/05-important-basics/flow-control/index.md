---
title: 5.4 流程控制
---

與流程控制有關的語法和範例。

## `for` loop {#for-loop}

底下是幾種常見的寫法：

```go
i := 1
for i <= 3 {
    fmt.Println(i)
}

for j := 0; j < 3; j++ {
    fmt.Println(j)
}

for {  // 無限迴圈
    fmt.Println("loop")
}
```

迴圈裡面可以用 `continue` 來進行下一圈，以及用 `break` 來跳離迴圈。

### For-each range loop

使用 `range` 關鍵字來指定索引值的範圍：

```go
for i := range 3 {  // i = 0, 1, 2
    fmt.Println("range", i)
}
```

常用來處理 arrays、slices、maps、channels 等結構：

```go
strings := []string{"hello", "world"}
for i, s := range strings {
    fmt.Println(i, s)
}
```

執行結果：

```text
0 hello
1 world
```

上例中，若不在乎陣列的索引值，可使用 blank identifier `_` 取代 `i`：

```go
strings := []string{"hello", "world"}
for _, s := range strings {
    fmt.Println(s)
}
```

執行結果：

```text
hello
world
```

## `if` statement

Go 的 `if` 陳述式不需要使用小括弧 `()`，但必須使用大括號 `{}`。雖然某些簡單的表達式可以使用小括弧，例如 `if (j > 10) {}` 這樣的寫法可以通過編譯，但與常見的寫法不一致，而且多此一舉。

> [!note]
> 事實上，在 VS Code 中編寫 Go 程式的時候，即使在 `if` 陳述式中使用了小括弧，那些小括弧會在存檔時被 [gofmt](https://pkg.go.dev/cmd/gofmt) 工具自動消除。這是因為 Go 工具鍊在 VS Code 中的預設配置為存檔時自動格式化。**參閱:** [01 Get started > VS Code](https://huanlin.cc/docs/go/01-get-started/#vs-code)。

`else` 敘述一定要跟在 `if` 區塊結尾的大括號後面，不能寫成單獨一行，像底下這樣的寫法將無法通過編譯：

```go
if j > 10 {
}
else { // 編譯錯誤! else 必須寫在 if 區塊結尾的同一行。
}
```

### If with a short statement {#if-with-a-statement}

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
