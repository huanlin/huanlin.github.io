---
title: 5.7 字串
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

## References

- The Go Blog: [Strings, bytes, runes and characters in Go](https://go.dev/blog/strings) by Rob Pike (2013-10-23)