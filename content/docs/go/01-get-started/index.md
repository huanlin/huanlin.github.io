---
title: 01 Get started
tags: [Go]
---

## Go 語言簡介

Go 於 2007 年誕生，由 Google 創建。

Go 程式語言沒有以下語法或機制：

- 沒有型別繼承。
- 沒有 exception handling 語法（沒有 `try...catch`，也沒有 `throw`）。
- 沒有巨集（macro）。
- 沒有 enum。（可以用具名常數）
- 沒有局部函式（partial functions）。
- 不支援變數延遲估值（lazy variable evaluation）。
- 沒有運算子多載（operator overloading）。
- 沒有樣式匹配（pattern matching）。
- 沒有內建的 GUI 框架或套件。

如欲了解為什麼 Go 不支援某些語言特性，可參閱官方文件：[Go FAQ](https://go.dev/doc/faq)。

Go 的優點與強項：

- 很適合開發 CLI 和伺服器端應用程式。
- 函式可回傳多個值。於是，函式可以輕易回傳錯誤，故也就不需要 throw exceptions 了。就如 Rob Pike 於 2015 年發表的文章所說，[errors are values](https://go.dev/blog/errors-are-values)。直到現在（2024 年）依然如此。
- Concurrency。非同步呼叫的語法非常簡單直觀，跟循序呼叫的語法幾乎一樣。
- 單元測試在 Go 語言中是一級公民：測試程式的檔案名稱一律命名為「*欲測試之程式檔名*`_test`.go」，而且兩個檔案要放在同一個目錄下。例如 hello.go 的測試程式會叫做 hello_test.go。
- 標準函式庫提供了常用的工具套件，包括網路通訊、HTTP、序列化、加解密等等。

> 如果需要開發跨平台的 GUI 應用程式，可以試試開源專案 [Wails](https://wails.io/)。

## 建立開發環境 {#setup-dev-env}

### 安裝 Go {#installing-go}

請參閱官方文件：[Download and install](https://go.dev/doc/install)

安裝完成後，開啟命令列視窗，執行 `go version` 命令查看版本。

> 撰寫本文時，我安裝的 Go 版本是 v1.23.0。

### VS Code

比較常聽到建議使用的 IDE 有這幾個：

- Visual Studio Code
- GoLand by JetBrains
- Neovim

對鍵盤操控和 coding 效率極為講究的人可能會喜歡 Neovim 或 JetBrains。由於我用的是 VS Code，所以這裡只介紹它的相關設定。

與 Go 有關的 VS Code extensions：

- [Go](https://marketplace.visualstudio.com/items?itemName=golang.go) by the Go Team at Google
- [Go Test Explorer](https://marketplace.visualstudio.com/items?itemName=premparihar.gotestexplorer)

底下截圖展示了我撰寫本文時的 VS Code 工作環境：

![](images/vscode-go.png)

左下角的 Go 面板可以查看 Go 環境變數以及安裝了哪些 Go tools。

VS Code 官方文件有更詳細的介紹：[Go in Visual Studio Code](https://code.visualstudio.com/docs/languages/go)。

如果已經有正確安裝 Go 工具鍊的相關工具，在預設情況下，按 Ctrl+S 存檔時會自動重新排版程式碼，可輕鬆維持一致的程式碼風格。如欲查看預設的自動排版選項，可以按 `F1` 或 `Crtl+Shift+P` 開啟命令面板，輸入 `Preferences: Open Default Settings (JSON)`，便可以找到所有跟 Go 有關的預設選項。底下僅摘錄其中一部份：

```json
  // Configure settings to be overridden for the go language.
  "[go]":  {
    "editor.insertSpaces": false,
    "editor.formatOnSave": true,
    "editor.codeActionsOnSave": {
        "source.organizeImports": true
    }
  },

  // Enable intellisense, code navigation, refactoring, formatting & diagnostics
  // for Go. The features are powered by the Go language server "gopls".
  "go.useLanguageServer": true,
```

其中提到的 "gopls"（讀作 "go please"）是官方提供的、用於 VS Code 的 Language Server。只要有安裝 gopls，在 VS Code 中撰寫程式就會有許多方便的編輯功能，像是 intellisense、重構、排版程式碼等等。

**參閱：** [gopls 官方文件](https://pkg.go.dev/golang.org/x/tools/gopls#section-readme)

順便提及，Go 提供的程式碼排版工具預設會使用 `tab` 來縮排，而不是插入空白字元，故剛才展示的預設選項中，`editor.insertSpace` 預設為 `false`。建議不要更改這個選項，以確保所有的 Go 程式碼維持同樣的風格。

#### 除錯

欲在 VS Code 中除錯 Go 程式，通常需要建立 **launch.json** 來提供一些必要的參數。

舉例來說，如果命令列應用程式執行的過程中有用到 `fmt.Scanf()` 來獲取使用者輸入的字元，在預設情況下，VS Code 除錯應用程式的時候是以整合式終端機視窗（integrated terminal）來顯示應用程式的執行過程，而這個整合式終端機並沒有辦法接受使用者輸入的字元。像這種情況，就會需要告訴 VS Code：除錯我的應用程式時，請改用外部的終端機視窗（external terminal）。底下是一個 launch.json 範例：

```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${fileDirname}",
            "console": "externalTerminal"
        }
    ]
}
```

有關建立 `launch.json` 的方法以及詳細的參數說明，請參閱 Go Wiki : [debugging](https://github.com/golang/vscode-go/wiki/debugging#configure)。
