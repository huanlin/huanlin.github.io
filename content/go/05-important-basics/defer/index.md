---
title: 5.7 defer 陳述式
---

Go 的 `defer` 關鍵字可用來將一個函式呼叫的執行時機延後至包覆函式（surrounding function）結束之前才執行，常用於清理資源（例如確保關閉資料庫連線）。

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

另外要注意的是，延後執行的時機除了函式正常返回，還有一種情況：goroutine 發生了執行時期的 panics。相關細節與注意事項可參閱官方文件：[Defer statements](https://go.dev/ref/spec#Defer_statements)。
