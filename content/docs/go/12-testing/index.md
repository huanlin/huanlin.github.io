---
title: Testing
draft: true
---

Test functions 名稱必須以 "Test" 開頭，否則 VS Code 的 Test Explorer 無法自動發現測試函式，故也無法執行它們。

以下測試函式的名稱都不合法：

- `MyFuncTest`
- `testMyFunc`
- `TESTMyFunc`

以下兩種測試函式的命名方式都不會出現 linter（靜態語法分析）警告：

- `TestMyFunc`
- `Test_MyFunc`

雖然測試函式名稱可以有底線，但不建議。

## 測試套件

- https://github.com/stretchr/testify

## References

- [Visual Studio Code User Guide: Testing](https://code.visualstudio.com/docs/editor/testing)
- [Testing in Go: Best Practices and Tips](https://grid.gg/testing-in-go-best-practices-and-tips/)
- [Testify 單元測試](https://wshs0713.github.io/posts/5793cee/)
- [Naming tests in Golang](https://medium.com/getground/naming-tests-in-golang-c58c188bb9a1)

