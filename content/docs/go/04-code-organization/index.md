---
title: 04 Code organization
tags: [Go]
---

## Packages

Go 應用程式是由多個 packages 組成，一個 **package** 在檔案系統中就是一個資料夾，該資料夾底下的 .go 程式檔案必然隸屬同一個 package（否則無法通過編譯）。

> [!quote]
> A **package** is a collection of source files in the same directory that are compiled together. Functions, types, variables, and constants defined in one source file are visible to all other source files within the same package.
>
> Go 官方部落格：[How to Write Go Code](https://go.dev/doc/code)

換言之，package 是一個邏輯切割單位，讓不同用途的程式碼之間得以適度隔離。

範例：

```text
.                           -> 專案的根目錄
├─ go.mod                   -> 定義專案的名稱和 dependencies
├─ hello.go                 -> 實作 package main
└─ auth/                    -> auth 套件
    ├─ auth.go              -> auth 相關功能的實作
    ├─ auth_test            -> auth 相關功能的測試
    └─ token/               -> token 套件
        ├─ token.go         -> token 相關功能的實作
        └─ token_test.go    -> token 相關功能的測試
```

Package 的名稱通常會跟它所在的資料夾名稱相同，但也可以不同。例如檔案 `auth.go` 裡面可能會宣告套件名稱為 `authentication`：

```go
package authentication

....
```

那樣的話，`auth_test.go` 的套件名稱也必須是 `authentication`，因為同一個資料夾底下的 .go 檔案必須隸屬同一個套件。

另外要知道的是，Go 的 package 有兩種：

- 可執行套件：套件名稱一定是 `main`，而且不能被其他套件引用。
- 函式庫套件：套件名稱不是 `main` 的都是函式庫套件，可供其他套件引用。

至於不同的 package 之間要如何開放或隱藏某些資源或服務，稍後會再說明。

### Package 名稱 {#package-names}

<mark>套件名稱應簡潔明白，通常是名詞，而且全都是用英文小寫。注意不可以用底線（snake case）或大小寫混和（mixedCaps）。</mark>

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

> [!note]
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

> [!quote]
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

如果這個專案的原始碼是放在 GitHub 平台上的 repository，那麼剛才的 `go mod init` 指令可以這麼寫：

```text
go mod init github.com/michael/todoapp
```

其中的 `michael` 是 GitHub 使用者帳戶。

上述命令同樣只是在當前目錄建立一個 `go.mod` 檔案，內容會變成：

```text
module github.com/michael/todoapp

go 1.23.0
```

其中的 `github.com/michael/todoapp` 即此專案儲存於 GitHub 平台上的路徑，但本機的檔案系統不需要有這樣的路徑，例如它可能存放在 D:/work/todoapp/ 目錄下，有著類似底下的目錄結構：

```text
todoapp/            -> 專案的根目錄
  ├─ go.mod         -> 定義專案的名稱和 dependencies
  ├─ main.go        -> 主程式套件
  └─ auth/          -> auth 套件
      ├─ auth.go    -> auth 相關功能的實作
      ├─ auth_test  -> auth 相關功能的測試
```

當 `main.go` 要引用（import）`auth` 套件時，它的 `import` 陳述句會像這樣：

```go
import (
    "github.com/michael/todoapp/auth"
)
```

也就是說，在引用 local folder 的套件時，`import` 的完整路徑名稱會是 `<module path>/<package path>`。此例的 module path 是 `github.com/michael/todoapp/`，而 package path 是 `auth`。

**See also:** [Golang imports tutorial](https://www.youtube.com/watch?v=Nv8J_Ruc280)

### Module paths

<mark>模組路徑是模組的正式名稱（唯一識別名稱）</mark>，宣告於模組的 go.mod 檔案，而且模組路徑要能表達該模組的用途，以及可以從何處找到它。

模組路徑通常包含三個部分：

- repository root path
- repository 中的目錄名
- 主要的版本編號（只有在主要版本編號為 2 或更高的版本才需要）

範例：

```go
module example.com/mymodule
```

如果此範例的模組的版本是 v0.x.x 或 v1.x.x，那麼它的 v2.0.0 版（以及之後版本）的模組路徑就要加上主版本號，例如：

```go
module example.com/mymodule/v2
```

> [!note]
> 模組名稱雖然可以不包含 URL，但是帶有 URL 的模組名稱有助於找到並下載該模組，而且可以確保名稱唯一，避免跟其他模組名稱衝突或混淆。因此，建議的做法是以 URL 的寫法來指定模組名稱。

最後整理幾個重點：

- 每一個 module 都是以 module path 來作為唯一識別，這個模組路徑是宣告在一個 go.mod 檔案中。
- Modules 可以直接從版本控制儲存庫下載，或者從 module proxy 伺服器下載。
- 使用 `import` 來引用模組中的套件時，只能引用該模組 export 的（公開的）套件。

建議閱讀 Go 官方文件以了解更多有關 modules 的細節：

- [Go Modules Reference](https://go.dev/ref/mod)
- [go.mod file reference](https://go.dev/doc/modules/gomod-ref)

### Download a module

如果你的專案要使用某個外部模組，就必須先使用 `go get` 命令來把它下載至本機電腦。

範例：

```text
go get github.com/huanlin/learning-go
```

上述命令會直接到遠端的 GitHub 主機下載我的 `learning-go` 專案。

請注意，只有當模組符合此條件時，`go get` 命令才能成功下載：<mark>模組有包含 `go.mod` 檔案，而且 `go.mod` 檔案中的模組路徑與它所在的遠端 Git repository 的路徑相同。</mark>

> 除了 GitHub 之外，只要是 Git-based repositories，也都可以利用此命令來下載位於遠端主機的模組。例如：BitBucket。

**See also:** [What is 'go get' command in Go](https://gosamples.dev/go-get/).

#### Troubleshooting

我曾經在使用 `go get` 命令時碰到以下錯誤訊息：

```text
go: github.com/huanlin/learning-go@upgrade (v0.0.0-20240904141749-ce362a80bcf3)
    requires github.com/huanlin/learning-go@v0.0.0-20240904141749-ce362a80bcf3:
    parsing go.mod:
        module declares its path as: learning-go
                but was required as: github.com/huanlin/learning-go
```

後來發現是因為這個專案的 `go.mod` 檔案中最初的模組路徑是寫成 `learning-go`，像這樣：

```text
module learning-go
```

將檔案推送至遠端 GitHub 主機之後，如果用 `go get github.com/huanlin/learning-go` 命令來下載該模組，雖然 GitHub 主機上的確有這個 repository，可是 `go get` 命令發現該模組的 `go.mod` 檔案裡面寫的模組路徑並不是 `github.com/huanlin/learning-go`，於是顯示錯誤訊息並拒絕下載。

此時若修改 `go.mod` 檔案的內容，把模組路徑改為相符的 `github.com/huanlin/learning-go`，再推送至 GitHub 主機，然後再嘗試執行一遍剛才的 `go get` 命令，結果還是會得到同樣的錯誤訊息。這是 Go 在本機電腦的模組快取機制所造成的現象。

解決方法是隨意修改專案中的某個檔案，然後推送變更至 GitHub 主機，以產生一個新的 commit。接著以 `go get` 命令下載模組，並且在命令結尾處附加 @*<commit-hash>* 即可。比如說，commit hash 是 aa1ff21，那麼下載模組的命令會像這樣：

```text
go get github.com/huanlin/learning-go@aa1ff21
```

或者，也可以在 GitHub 上面建立一個 Release Tag，例如：`v0.0.1-beta`，然後告訴 `go get` 命令要下載這個版本的模組：

```text
go get github.com/huanlin/learning-go@v0.0.1-beta
```

### `go mod tidy` command

`go mod tidy` 命令會找到專案的 `go.mod` 檔案，並針對模組依賴關係執行以下工作：

- 刪除沒用到的模組。
- 下載需要的模組。
- 更新 `go.mod` 和 `go.sum` 檔案。

舉例來說，假設 go.mod 檔案內容是：

```text
module github.com/huanlin/learning-go

go 1.23.0

require github.com/shirou/gopsutil/v4 v4.24.8
```

執行 `go mod tidy` 命令時，它會自動下載需要的套件：

```console
go: downloading github.com/shirou/gopsutil/v4 v4.24.8
go: downloading github.com/shoenig/go-m1cpu v0.1.6
go: downloading github.com/tklauser/go-sysconf v0.3.12
go: downloading github.com/power-devops/perfstat v0.0.0-20210106213030-5aafc221ea8c

...(略)
```

並且更新 `go.mod` 檔案：

```text
module github.com/huanlin/learning-go

go 1.23.0

require github.com/shirou/gopsutil/v4 v4.24.8

require (
    github.com/go-ole/go-ole v1.2.6 // indirect
    github.com/lufia/plan9stats v0.0.0-20211012122336-39d0f177ccd0 // indirect
    github.com/power-devops/perfstat v0.0.0-20210106213030-5aafc221ea8c // indirect
    github.com/shoenig/go-m1cpu v0.1.6 // indirect
    ...(略)
)
```

可以看到 `go mod tidy` 命令自動替 `go.mod` 檔案加入了一堆間接依賴的模組。

最後是 `go.sum` 檔案：

```text
github.com/davecgh/go-spew v1.1.1 h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=
github.com/davecgh/go-spew v1.1.1/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/go-ole/go-ole v1.2.6 h1:/Fpf6oFPoeFik9ty7siob0G6Ke8QvQEuVcuChpwXzpY=
github.com/go-ole/go-ole v1.2.6/go.mod h1:pprOEPIfldk/42T2oK7lQ4v4JSDwmV0As9GaiUsvbm0=
github.com/google/go-cmp v0.5.6/go.mod h1:v8dTdLbMG2kIc/vJvl+f65V22dbkXbowE6jgT/gNBxE=
...(略)
```

## Project layout

GitHub 平台上面有一個用來展示 Go 專案結構的 repository 可以參考：[project-layout](https://github.com/golang-standards/project-layout)。

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