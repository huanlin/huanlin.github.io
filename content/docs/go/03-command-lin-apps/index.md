---
title: 03 命令列應用程式
tags: [Go]
draft: true
---

## 使用 os.Args 取得命令列參數 {#os-args}

Go 標準函式庫的 `os` 套件有一個公開變數 `Args` 可用來取得應用程式執行時所傳入的命令列參數。

- `os.Args[0]` 是執行應用程式時的檔案名稱。
- `os.Args[1]` 是第 1 個命令列參數。
- `os.Args[2]` 是第 2 個命令列參數。依此類推。

範例：

```go
ackage main

import (
    "fmt"
    "log"
    "os"
    "strings"
)

func main() {
    if len(os.Args) < 2 {
        log.Println("need to provide filename!")
        os.Exit(1)
    }

    fileContents, err := os.ReadFile(os.Args[1])
    if err != nil {
        log.Println(err)
        os.Exit(1)
    }

    words := strings.Fields(string(fileContents))

    fmt.Println("Found", len(words), "words")
}
```

## 剖析命令列參數 {#command-line-flags}

Go 標準函式庫中的命令列參數和選項的處理方式是基於 Plan 9 作業系統，這與現今廣泛使用的 GNU/Linux 和 BSD (Berkeley Software Distribution) 的系統（如 Mac OS X 和 FreeBSD）有所不同。

例如，在 Linux 和 BSD 系統上，您可以使用命令 `ls -la` 來列出目錄中的所有文件。`ls` 是命令，而 `-la` 包含了兩個選項，或者兩個旗標（flags，這是 Go 函式庫的用語）。`l` 旗標告訴 `ls` 使用長格式列出檔案，而 `a` 旗標（代表 "all"）則使列表包含隱藏文件。

Go 的旗標系統不允許結合多個旗標，所以它會把 `-la` 視為一個旗標。Go 之所以如此設計，部分原因是 Go 將一個減號開頭的短命令行旗標（`-la`）與兩個減號開頭的長旗標（`--la`）視為相同。

另一方面，GNU 風格的命令（如 `ls`）支援長選項。例如 `--color`，它是以兩個減號來告訴應用程式：「後面跟著的字串 `color` 不是五個選項，而是一個選項。」

為了讓我們開發的 Go 應用程式支援 Linux 風格的命令選項，比較好的解法是使用現成的套件。底下是兩個頗受歡迎的套件：

- [cobra](https://github.com/spf13/cobra) - 許多 Go 專案都有使用，如 Kubernetes、Hugo、和 GitHub CLI 等等。
- [cli](https://github.com/urfave/cli) - 開源的 platform-as-a-service (PaaS) 專案 Cloud Foundry 有使用此套件。

## 應用程式的組態 {#app-config}

- **問題**：

    應用程式需要保存組態，因為這些組態可能數量眾多，命令列參數難以應付。

- **解決方案**：

    目前比較受歡迎的一種組態檔案格式是 JSON (JavaScript Object Notation)。Go 標準函式庫提供了內建的 JSON 解析、反序列化和序列化功能。另外常見的兩種組態檔格式為 YAML 和 INI 檔案。跟 JSON 比起來，YAML 和 INI 的好處是人類更容易閱讀（而且 JSON 不能寫註解）。

> [12-factor apps](http://12factor.net/) 方法所建議的模式：透過環境變數來傳遞 configuration。

### JSON

```json
{
     "enabled": true,
     "path": "/usr/local"
}
```

### YAML

```yaml
# A comment line
enabled: true
path: /usr/local
```

### INI

```ini
; A comment line
[Section]
enabled = true
path = /usr/local # another comment
```

### 使用環境變數來保存組態 {#config-via-env-vars}


## Working with web servers {#web-servers}

