---
title: 04 建立命令列應用程式
linkTitle: 04 建立 CLI 應用程式
draft: true
---

Go 標準函式庫中的命令列參數和選項的處理方式是基於 Plan 9 作業系統，這與現今廣泛使用的 GNU/Linux 和 BSD (Berkeley Software Distribution) 的系統（如 Mac OS X 和 FreeBSD）有所不同。

例如，在 Linux 和 BSD 系統上，您可以使用命令 `ls -la` 來列出目錄中的所有文件。`ls` 是命令，而 `-la` 包含了兩個選項，或者兩個旗標（flags，這是 Go 函式庫的用語）。`l` 旗標告訴 `ls` 使用長格式列出檔案，而 `a` 旗標（代表 "all"）則使列表包含隱藏文件。

Go 的旗標系統不允許結合多個旗標，所以它會把 `-la` 視為一個旗標。Go 之所以如此設計，部分原因是 Go 將一個減號開頭的短命令行旗標（`-la`）與兩個減號開頭的長旗標（`--la`）視為相同。

另一方面，GNU 風格的命令（如 `ls`）支援長選項。例如 `--color`，它是以兩個減號來告訴應用程式：「後面跟著的字串 `color` 不是五個選項，而是一個選項。」

為了讓我們開發的 Go 應用程式支援 Linux 風格的命令選項，比較好的解法是使用現成的套件。底下是兩個頗受歡迎的套件：

- [corba](https://github.com/spf13/cobra) - 許多 Go 專案都有使用，如 Kubernetes、Hugo、和 GitHub CLI 等等。
- [cli](https://github.com/urfave/cli) - 開源的 platform-as-a-service (PaaS) 專案 Cloud Foundry 有使用此套件。

ref: https://livebook.manning.com/book/go-in-practice-second-edition/chapter-2/v-6/21
