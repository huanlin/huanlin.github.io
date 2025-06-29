---
title: Gemini CLI 提示詞
---

## 限制與注意事項

- 無法將「非 UTF-8」編碼的檔案改為 UTF-8 編碼。

## .NET and C#

### Solution and Project

```text
分析此 solution 的程式架構，列出可改進之處。
將此 repository 的 .NET 專案都改用 Central Package Management (CPM)。
更新目標框架為 net9.0，並修正不相容的 API。
幫我看看 GitHub Actions workflow 配置檔案有沒有不恰當之處，例如打包 package 的時候不應該包含測試專案。
```

### Code file

```text
我已經復原了上一次的變更。請重新讀取每一個 C# 檔案，並使用 file-scoped namespaces 來減少巢狀。When modifying files, plese use UTF-8 encoding with signatue.  
```

