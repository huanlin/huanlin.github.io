---
title: 04 Code organization
tags: [Go]
---

## Packages

Go 應用程式是由多個 packages 組成，而一個 **package** 是一個或多個 .go 程式檔案所組成；這些程式檔案會放在同一個資料夾底下，而這個資料夾的名稱通常會跟 package 名稱一樣。

> A **package** is a collection of source files in the same directory that are compiled together. Functions, types, variables, and constants defined in one source file are visible to all other source files within the same package.
>
> Go 官方部落格：[How to Write Go Code](https://go.dev/doc/code)

換言之，package 一個邏輯切割單位，讓不同用途的程式碼之間得以適度隔離。

範例：

```text
.                -> 專案的根目錄
├── go.mod       -> 定義專案的名稱和 dependencies
├── hello.go     -> 實作 package main
└── cart         -> 用來放 cart 套件的程式碼
    └── cart.go  -> 實作 cart 套件
```

另外要知道的是，Go 的 package 有兩種：

- 可執行套件：套件名稱一定是 `main`，而且不能被其他套件引用。
- 函式庫套件：套件名稱不是 `main` 的都是函式庫套件，可供其他套件引用。

至於不同的 package 之間要如何開放或隱藏某些資源或服務，請看下一節的說明。

### Package 的命名 {#package-names}

好的套件名稱應簡潔明白，通常是名詞，而且全都是用英文小寫。注意不可以用底線（snake case）或大小寫混和（mixedCaps）。

範例：

- `list`
- `http`
- `strconv` （兩個單字的縮寫組合: string conversion）
- `syscall` （兩個單字的縮寫組合：system call）
- `fmt` （format 的縮寫）

詳見 Go 官方部落格：[Package names](https://go.dev/blog/package-names)。

### Scope

程式裡面的識別字（identifiers），像是變數、函式、型別等等，依照它們宣告時的所在位置和寫法，分為三種可見範圍：

| 範圍 | 說明 |
|-----|------|
| block | 宣告在 `{...}` 區塊裡面的識別字只有該區塊的程式碼可存取。 |
| package | 同一個 package 內的 .go 程式檔案皆可存取。 |
| global | 任何程式碼皆可存取。 |

有兩種情況會是 global 範圍：

- Go 的內建函式，例如 `panic()`。
- 在 package 層級宣告的識別字，名稱以英文大寫字母開頭來命名，就會被編譯器視為 **exported**，亦即公開的。

參考以下範例：

```go
package config

var ConfigFileName string = "d:/work/config.yaml" // 任何套件皆可存取。
var encoding string = "UTF-8" // 僅相同 package 的程式碼可以存取。

func createConfig() { // 僅相同 package 的程式碼可以存取。
    num := 100 // 只在此函式內可見。
}
```

那麼，如果有兩個 .go 程式檔案放在同一個 package 裡面，有辦法讓其中一個 .go 程式檔案中的全域變數隱藏起來，不讓另一個 .go 程式檔案存取嗎？

答案是不行。（這裡的全域變數指的是沒有包在任何 `{...}` 區塊中的變數）

這是因為，同一套件裡面的任何東西只要沒有被 `{...}` 包起來，在同一個套件範圍內都是共享的。這是 Go 語言賦予 package 的特性和規則。

> 也許有人會覺得這是 Go 語言的一個限制或缺點，但從另一個角度來看，寫程式的時候不用老是費心去考慮某些變數或函式到底要隱藏到什麼程度，也能讓事情變得簡單一點。

### Variable shadowing

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

## Modules

一個 Go 專案通常只包含一個模組（**module**），亦可能包含多個 modules，而每個 module 是由一個或多個 packages 所組成。

專案的主模組會放在 repository 的根目錄，其名稱通常會跟專案的 repository 名稱相同。模組的根目錄之下需要一個 `go.mod` 檔案來設定專案的基本資訊（名稱、版本）以及管理它所依賴的外部模組。簡單來說，一個 module 通常代表一個應用程式專案。

> A repository contains one or more modules. A **module** is a collection of related Go packages that are released together. A Go repository typically contains only one module, located at the root of the repository.
>
> Go 官方部落格：[How to Write Go Code](https://go.dev/doc/code)

建立 `go.mod` 檔案的方式，是在專案根目錄底下執行 `go mod init` 命令。比如說，專案名稱是 `todoapp`，便可使用以下命令來建立 `go.mod` 檔案：

```shell
mkdir todoapp
cd todoapp
go mod init todoapp
```

上述命令所建立的 `go.mod` 檔案，其內容會像這樣：

```text
module todoapp

go 1.23.0
```

如果這個專案的原始碼是放在 GitHub 平台上的一個名為 "learning-go" 的 repository，那麼剛才的 `go mod init` 指令會這麼寫：

```text
mkdir todoapp
cd todoapp
go mod init github.com/todoapp
```

上述命令同樣只是在當前目錄建立一個 `go.mod` 檔案，內容會變成：

```text
module github.com/todoapp

go 1.23.0
```

### Module paths

模組路徑是模組的正式名稱（唯一識別名稱），宣告於模組的 go.mod 檔案，而且模組路徑要能表達該模組的用途，以及可以從何處找到它。

模組路徑通常包含三個部分：

- repository root path
- repository 中的目錄
- 主要的版本編號（只有在主要版本編號為 2 或更高的版本才需要）

範例：

```go
module example.com/mymodule
```

如果此範例的模組的版本是 v0.x.x 或 v1.x.x，那麼它的 v2.0.0 版（以及之後版本）的模組路徑就要加上主版本號，例如：

```go
module example.com/mymodule/v2
```

{{% admonition type=note title="Note" %}}
模組名稱雖然可以不包含 URL，但是帶有 URL 的模組名稱有助於找到並下載該模組，而且可以確保名稱唯一，避免跟其他模組名稱衝突或混淆。因此，建議的做法是以 URL 的寫法來指定模組名稱。
{{% /admonition %}}

最後整理幾個重點：

- 每一個 module 都是以 module path 來作為唯一識別，這個模組路徑是宣告在一個 go.mod 檔案中。
- Modules 可以直接從版本控制儲存庫下載，或者從 module proxy 伺服器下載。
- 使用 `import` 來引用模組中的套件時，只能引用該模組 export 的（公開的）套件。

建議閱讀 Go 官方文件以了解更多有關 modules 的細節：

- [Go Modules Reference](https://go.dev/ref/mod)
- [go.mod file reference](https://go.dev/doc/modules/gomod-ref)

## Summary

- Go 是以 package 來作為隔離的基本單位。
- 隸屬同一個 package 的程式碼可以互相存取宣告於 package 層級的名稱，如變數、函式、型別等等。
- 不同 package 的程式碼只能使用對方 export 出來的東西。
  - Go 語言沒有 `public`、`private` 或 `protected` 等識別字，而是根據變數名稱的第一個字母大小寫來判斷能否被外部引用。
  - 大寫字母開頭的名稱會被 export，即可供外界使用。
  - 小寫字母開頭的名稱無法被外界存取。

## References

- Go 官方部落格：[How to Write Go Code](https://go.dev/doc/code)

[100-mistakes]: https://www.manning.com/books/100-go-mistakes-and-how-to-avoid-them