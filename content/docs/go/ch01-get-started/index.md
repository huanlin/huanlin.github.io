---
title: Chapter 1 - Get started
weight: 11
---

Ref: [[Golang] Modules and Packages](https://pjchender.dev/golang/modules-and-packages/)

## Go 的語言特性

- 沒有內建的 GUI 框架或套件。
- Concurrency。
- 函式可回傳多個值。
- 標準函式庫提供了常用的工具套件，包括網路通訊、HTTP、序列化、加解密等等。

> 如果需要開發跨平台的 GUI 應用程式，可以試試開源專案 [Wails](https://wails.io/)。

## Module

什麼是 Go module，以及它的幾個重要特性：

- 一個 module 是一個 project，有一個版本編號。
- 一個 module 包含一個或多個 packages。
- Modules 可以直接從版本控制儲存庫下載，或者從 module proxy 伺服器下載。
- 每一個 module 都是以 module path 來作為唯一識別，這個模組路徑是宣告在一個 go.mod 檔案中。

### Hello, World

使用 `go mod init` 命令來建立模組：

```shell
cd hellogo
go mod init hellogo
```

此命令會在當前目前建立一個 `go.mod` 檔案，內容是該模組的資訊，以及描述它依賴哪些模組（如果有的話）。上述指令所建立的檔案內容會像這樣：

```text
module hellogo

go 1.23.0
```

以此範例而言，`hellogo` 目錄即成為你的 project 的 main module，而 Go 編譯器在建置應用程式時，便會參考此目錄下的 `go.mod` 檔案。

> `go.mod` 檔案所在的目錄稱為 **模組根目錄**（module root directory）。

接著在此目錄中建立一個 `hello.go` 檔案，內容為：

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

- `package` 表明這個模組的名稱叫做 `main`。
- `import` 表明此模組需要引用 `fmt` 模組。
- `main()` 函式為每一個 Go 應用程式的進入點。

使用 `go run` 命令來執行此程式：

```shell
go run main.go
```

也可以用 `go build` 命令來將程式碼編譯成可執行檔：

```shell
go build main.go
```

### Packages

Go 的 package 有兩種：

- executable package：會編譯成可執行的應用程式，其主模組的名稱必須是 `main`，而且會包含程式的進入點：`main` 函式。
- library package：供其他套件引用，不會編譯成可執行檔。套件名稱不用是 `main`。

### Module paths

模組路徑是模組的正式名稱（唯一識別名稱），宣告於模組的 go.mod 檔案，而且模組路徑要能表達該模組的用途，以及可以從何處找到它。

模組路徑通常包含三個部分：

- repository root path
- repository 中的目錄
- 主要的版本編號（只有在主要版本編號為 2 或更高的版本才需要）

範例：

```go
module example.com/x/mod
```

如果此範例的模組的版本是 v1.0.0，那麼它的 v2.0.0 版（以及之後版本）的模組路徑就要加上主版本號，例如：

```go
module example.com/x/mod/v2
```

## More about modules

- Go 語言沒有 `public`、`private` 或 `protected` 等識別字，而是根據變數名稱的第一個字母大小寫來判斷能否被外部引用。
- 所有大寫字母開頭的名稱都會被 export，即可供外界使用。（等同其他物件導向語言的 `public` 存取範圍）
- 所有小寫字母開頭的名稱只能在模組內部使用。
- 使用 `import` 來引用模組中的套件時，只能引用該模組 export 的（公開的）套件。

## Recommended reading

建議閱讀 Go 官方文件以了解更多細節：

- [Go Modules Reference](https://go.dev/ref/mod)
- [go.mod file reference](https://go.dev/doc/modules/gomod-ref)

## References

- [Go in Action, 2nd Edition](https://www.manning.com/books/go-in-action-second-edition)
- [Go in Practice, 2nd Edition](https://www.manning.com/books/go-in-practice-second-edition)
