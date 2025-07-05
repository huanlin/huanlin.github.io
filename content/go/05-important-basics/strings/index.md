---
title: 5.8 字串
---

## 概要

- Go 的字串內部不是字元陣列，而是代表每個 UTF-8 字元的一串 bytes。
- Go 的原始碼以及官方文件是以 rune（讀音近似「潤」）來指稱一個 UTF-8 字元的 code point。簡單起見，可以把它理解為一個 UTF-8 字元。
- 使用 string 宣告變數若無指定初值，預設值為空字串 ""。
- 比較兩個字串，可以用：
  - `==`、`!=`、`<`、`>`、`<=`、`>=` 運算子。
  - `strings.EqualFold()` 函式：用於不區分大小寫的比較。
  - `strings.Compare()` 函式：應該只用於三向比較（three-way comparison）的場合。
- Raw string 的寫法是以 backtick 字元 (`\``) 包住字串。

## 格式化字串 {#fmt}

格式化字串的相關函式與參數，可參閱官方文件: [fmt package](https://pkg.go.dev/fmt)。其中包括常用的字串格式化參數、如何建立錯誤訊息、以及掃描（scan）字串等等，都有詳細的說明。網路上也有一些整理好的小抄，例如：[GoLang fmt Printing Cheat Sheet](https://cheatography.com/gpascual/cheat-sheets/golang-fmt-printing/)。

底下是幾個常用的格式化參數（verbs）：

| Verbs | 用途說明 |
|-------|---------|
| `%d`  | 輸出整數。範例：45 以 "`%04d`" 格式化的結果為 "0045"。 |
| `%f`  | 輸出浮點數。範例："`%.2f`" 指定小數點到第二位，"`%9.f`" 指定整數部分寬度為 9，不顯示小數。 |
| `%s`  | 輸出字串內容。|
| `%q`  | 輸出以雙引號包住的字串，而且會以跳脫字元（escape character）來呈現無法印出的字元，例如 `\n`。 |
| `%v`  | 輸出型別的預設格式。 |
| `%+v` | 類似 `%v`，特別用於輸出結構內容，而且會輸出欄位名稱。|
| `%%`  | 輸出一個 `%` 字元。 |
| `%t`  | 輸出 `bool` 型別的值：`true` 或 `false`。 |
| `%T`  | 輸出型別名稱。 |
| `%x`  | 以英文小寫顯示 16 進制的數值。例如整數值 255 會輸出 `ff`。 |
| `%X`  | 以英文大寫顯示 16 進制的數值。例如整數值 255 會輸出 `FF`。 |

## 字串長度 {#str-length}

Go 的字串內部不是字元陣列，而是代表每個 UTF-8 字元的一串 bytes。因此，若以內建函式 `len` 試圖取得字串長度，得到的不會是字元個數，而是其內部 bytes 區塊的長度。如欲取得字串長度，應使用標準函式庫的 `utf8.RuneCountInString()`。

範例：

```go
import (
    "fmt"
    "unicode/utf8"
)

func main() {
    str := "地鼠"
    fmt.Println(len(str))                    // 輸出: 6
    fmt.Println(utf8.RuneCountInString(str)) // 輸出: 2
}
```

Try it: <https://go.dev/play/p/365ZZEx2uGz>

> [!quote]
> In Go, a string is in effect a read-only slice of bytes.
>
> -- The Go Blog: [Strings, bytes, runes and characters in Go](https://go.dev/blog/strings)

## Rune

如欲存取字串中的個別字元，應使用 `rune` 型別，而且不能使用陣列索引的語法，否則結果不會是我們想要的。

範例：

```go
str := "地鼠"
for i := 0; i < len(str); i++ {
    fmt.Print(string(str[i]) + " ")
}
fmt.Println() // 輸出:  å  ° é ¼
```

要取出字串中的字元，可以用 `range` 來取出型別為 `rune` 的字元：

```go
unicodeCharStr := "地鼠"
for i, rune := range unicodeCharStr {
    fmt.Printf("%d:%s ", i, string(rune))
}
fmt.Println() // 輸出: 0:地 3:鼠
```

在 Go 的原始碼以及官方文件中都是以 rune（讀音近似「潤」）來指稱一個 UTF-8 字元的 code point。型別 `rune` 只是 `int32` 的別名，換言之，一個 rune 就是一個 32 位元的整數，其數值範圍已足夠容納所有的 [Unicode](https://home.unicode.org/) code point。

> [!quote]
> "Code point" is a bit of a mouthful, so Go introduces a shorter term for the concept: rune. The term appears in the libraries and source code, and means exactly the same as "code point", with one interesting addition.
>
> The Go language defines the word rune as an alias for the type int32, so programs can be clear when an integer value represents a code point.
>
> -- The Go Blog: [Strings, bytes, runes and characters in Go](https://go.dev/blog/strings)

## 字串比較 {#str-compare}

一般的字串比較，建議使用運算子：`==`、`!=`、`<`、`>`、`<=`、`>=`。

如果比較時不區分英文大小寫，則使用 `strings.EqualFold()` 函式。另一種方法是把兩個字串先用 `strings.ToLower()` 轉成全部小寫，然後再用 `==` 比較。

如果需要**三向比較**（three-way comparison），亦即需要判斷兩個字串是大於、小於、還是等於，則可以使用 `strings.Compare()`。基於效率考量，此函式應該只用於三向比較的場合。在 Go 原始碼裡面也有這樣的建議（參見 [src/strings/compare.go](https://go.dev/src/strings/compare.go)）：

```go
// Compare returns an integer comparing two strings lexicographically.
// The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
//
// Use Compare when you need to perform a three-way comparison (with
// [slices.SortFunc], for example). It is usually clearer and always faster
// to use the built-in string comparison operators ==, <, >, and so on.
func Compare(a, b string) int {
    return bytealg.CompareString(a, b)
}
```

> [!note]
> 這個 `strings.Compare()` 函式原本效率較差，直到 Go v1.23 終於有了改善。詳見：[strings: intrinsify and optimize Compare](https://github.com/golang/go/commit/fd999fda5941f215ef082c6ef70e44e648db5485)。

## 練習：打亂字串內容 {#string-shuffle}

練習內容：

1. 寫一個函式將傳入的字串內容打亂（shuffle），將此函式命名為 `shuffle1`。
2. 再寫一個函式 `shuffle2`，作用與 `shuffle1` 相同，只是用不同的做法來達成相同目的。
3. 撰寫效能測試來觀察 `shuffle1` 和 `shuffle2` 的效能差異。

此練習的目的如下：

- 處理字串中的字元（`rune`）。
- 產生隨機數字（使用 `math/rand` 套件）。
- 撰寫和執行效能測試。

### 兩個打亂字串的函式 {#shuffle-functions}

將以下程式碼儲存為 `main.go`。

```go
package main

import (
    "fmt"
    "math/rand"
    "strings"
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

**Try it:** <https://go.dev/play/p/6r0H7l-ktAB>

說明：

- `shuffle1()` 先把傳入的字串 `s` 轉換成一個 `rune` 切片（`[]rune`，亦即 `[]int32`），然後把切片中的每一個字元跟另一個隨機位置的字元交換。
- `shuffle2()` 則是用 `string.Split()` 把字串分割成 `string` 切片（`[]string`），使得切片中的每一個字串都只包含一個 UTF-8 字元。然後再把切片中的每一個字串取出，跟另一個隨機位置的字串交換。最後再以 `strings.Join()` 把字串切片組合成一個字串。

程式執行結果（每次都不一樣）：

```text
2153097864
3596708124
```

### 撰寫效能測試 {#benchmark-code}

接著要撰寫效能測試來了解兩個函式的效能表現。首先，新增一個 Go 程式檔案，命名為 `main_test.go`。程式碼如下：

```go
package main

import (
    "testing"
)

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

### 執行效能測試 {#run-benchmark}

接著用 `go test` 命令來執行效能測試：

```shell
go test -bench . -benchmem
```

加上 `-bench` 選項即表示要執行效能測試，而 `-benchmem` 選項表示要一併觀察記憶體的使用情況。

> [!note] 備註
> 許多文件寫的測試命令是 `go test -bench=.`，也就是 `-bench` 選項後面是等於符號 '`=`'，而非空白字元。在我的 Windows 機器上使用 `-bench=.` 來執行測試，結果會說找不到任何測試："no tests to run"。

執行上述命令時，Go 測試工具會尋找 `*_test.go` 檔案中所有以 `Benchmark` 開頭的函式，並且對這些函式發出好幾輪的的呼叫；每一輪測試都會傳入一個型別為 `*testing.B` 的參數 `b`，而 `b.N` 就是測試工具對測試函式的指示：「請執行你的工作 `b.N` 次。」

> **重點整理**
>
> 單元測試的函式名稱是以 `Test` 開頭，效能測試的函式名稱則是以 `Benchmark` 開頭。這兩種測試函式都是寫在 `*_test.go` 檔案中。

每一輪測試完成後，測試工具會根據那一輪測試所耗費的時間來決定下一輪的 `b.N` 要增加至多少。越到後面，`b.N` 數值增加得越快。比如說，可能會以 1, 2, 3, 5, 10, 20, 30, 50, 100 這樣的速度遞增（只是舉例，方便了解）。

執行結果：

```text
goos: windows
goarch: amd64
pkg: demostring
cpu: 11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz
BenchmarkShuffle1-8   1758118   638.2 ns/op   192 B/op   2 allocs/op
BenchmarkShuffle2-8   1226618   999.8 ns/op   688 B/op   2 allocs/op
PASS
ok      demostring    3.097s
```

## References

- The Go Blog: [Strings, bytes, runes and characters in Go](https://go.dev/blog/strings) by Rob Pike (2013-10-23)
- <https://pkg.go.dev/testing>
- [Benchmarking in Golang: Improving function performance](https://blog.logrocket.com/benchmarking-golang-improve-function-performance/)
- [Go 语言高性能编程 - benchmark 基准测试](https://geektutu.com/post/hpg-benchmark.html)
