---
title: Go Cookbook
draft: true
---

## Benchmark

假設要對一個字串的內容隨機編排，並且有兩個函式 `shuffle1` 和 `shuffle2` 都能滿足此需求。程式碼如下：

```go
package main

import (
    "fmt"
    "math/rand"
    "strings"
    "testing"
)

func main() {
    s := "0123456789"
    s1 := shuffle1(s)
    fmt.Println(s1)

    s2 := shuffle2(s)
    fmt.Println(s2)
}

func shuffle1(s string) string {
    runes := []rune(s)
    for i := range runes {
        j := rand.Intn(len(runes))
        runes[i], runes[j] = runes[j], runes[i]
    }

    return string(runes)
}

func shuffle2(s string) string {
    sSlice := strings.Split(s, "")
    for i := range sSlice {
        j := rand.Intn(len(sSlice))
        sSlice[i], sSlice[j] = sSlice[j], sSlice[i]
    }
    return strings.Join(sSlice, "")
}
```

### 撰寫效能測試 {#benchmark-code}

現在我們想撰寫效能測試來了解這兩個函式的效能表現。做法是新增一個測試檔案，例如 main_test.go。程式碼如下：

```go
func BenchmarkShuffle1(b *testing.B) {
    for i := 0; i < b.N; i++ {
        shuffle1("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
    }
}

func BenchmarkShuffle2(b *testing.B) {
    for i := 0; i < b.N; i++ {
        shuffle2("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
    }
}
```

### 執行效能測試 {#running-benchmark}

```shell
go test -bench .
```

參考資料：

- [Go by Example: Testing and Benchmarking](https://gobyexample.com/testing-and-benchmarking)
- [Benchmarking in Golang: Improving function performance](https://blog.logrocket.com/benchmarking-golang-improve-function-performance/)

## 環境變數 {#env-var}

使用套件：<https://github.com/joho/godotenv>

