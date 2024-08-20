---
title: Chapter 1 - Get started
weight: 11
---

## Go 的語言特性

 - 沒有內建的 GUI 框架或套件。
 - Concurrency。
 - 函式可回傳多個值。
 - 標準函式庫提供了常用的工具套件，包括網路通訊、HTTP、序列化、加解密等等。

> 如果需要開發跨平台的 GUI 應用程式，可以試試開源專案 [Wails](https://wails.io/)。

## Module

一個 Go module 就是一個 project。

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

接著在此目錄中建立一個 `main.go` 檔案，內容為：

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

About modules:

- 所有小寫字母開頭的型別、變數、或常數，只有在模組內部才能使用。
- 所有大寫字母開頭的型別、變數、或常數，都會 export 供外界使用。

## References

- [Go in Action, 2nd Edition](https://www.manning.com/books/go-in-action-second-edition)
- [Go in Practice, 2nd Edition](https://www.manning.com/books/go-in-practice-second-edition)
