---
title: "Visual Studio Code 快捷鍵"
tags: ["VS Code"]
---

底下是官方提供的各作業系統平台的 VS Code 快捷鍵速查表：

- [Windows](https://go.microsoft.com/fwlink/?linkid=832145)
- [macOS](https://go.microsoft.com/fwlink/?linkid=832143)
- [Linux](https://go.microsoft.com/fwlink/?linkid=832144)

底下是我自己常用的快捷鍵（for Windows），有些按鍵因為已經熟記而未列出，例如除錯相關按鍵（F9、F10、F11 等等）。

![](images/vscode-shortcut-keys.png#center)

## General and files

| 快捷鍵                  | 用途說明 |
|-------------------------|----------------------|
| Ctrl+Shift+P, F1        | 顯示 Command Palette |
| Ctrl+K Ctrl+S           | 顯示快捷鍵列表 |
| Ctrl+P                  | 切換至檔案... |
| Ctrl+K P | 把目前檔案的所在路徑複製到剪貼簿 |
| Ctrl+K R | 在檔案總管中顯示目前檔案 |
| Ctrl+K S | 儲存全部檔案 |

## Basic editing

| 快捷鍵                  | 用途說明 |
|-------------------------|----------------------|
| Ctrl+X | 在沒有選取文字的情況下直接剪下一列 |
| Ctrl+C | 在沒有選取文字的情況下直接複製一列 |
| Alt+ `↑` / `↓` | 把當前所在列往上／往下移動 |
| Shift+Alt + `↓` / `↑` | 把當前所在列往上／往下複製 |
| Ctrl+Enter | 往下插入一列 |
| Ctrl+ `↑` / `↓` | 上下捲動 |

## Navigation

| 快捷鍵                  | 用途說明 |
|-------------------------|----------------------|
| Ctrl+T | 顯示所有符號 |
| Ctrl+G | Go to Line... |
| Ctrl+P | 切換至檔案... |
| Ctrl+Shift+O | 顯示 Problems 面板 |
| F8 | 到下一個錯誤或警告 |
| Shift+F8 | 到上一個錯誤或警告 |
| Alt+ `←` / `→` | 回上一頁／到下一頁 |

## Multi-cursor and selection

| 快捷鍵                  | 用途說明 |
|-------------------------|----------------------|
| Ctrl+Alt+ `↑` / `↓` | 往上／往下增加游標 |
| Ctrl+Shift+L | 選取文字之後再用此按鍵一次選取檔案內所有相同的文字 |
| Ctrl+F2 | 一次選取檔案內所有「跟游標所在位置之字詞」相同的文字 |
| Shift+Alt+`→` | 擴大選取範圍 |
| Shift+Alt+`←` | 縮小選取範圍 |
| Shift+Alt+滑鼠拖曳 | 垂直選取（方塊選取） |
| Ctrl+Shift+Alt+方向鍵 | 垂直選取（方塊選取） |

## Display

| 快捷鍵                  | 用途說明 |
|-------------------------|----------------------|
| Ctrl+ `=` / `-` | Zoom in/out（用九宮格的 `+` 和 `-` 也行） |
| Ctrl+Shift+E | 到 Explorer 面板 |
| Ctrl+Alt+F   | 在 Explorer 面板中搜尋檔案名（不含資料夾名） |
| Ctrl+Shift+F | 到 Search 面板 |
| Ctrl+Shift+H | 對全部檔案進行搜尋並替換 |
| Ctrl+K V     | 在側邊面板顯示 Markdown 預覽結果 |

## Rich languages editing

| 快捷鍵                  | 用途說明 |
|-------------------------|----------------------|
| Shift+Alt+F | 格式化整個檔案 |
| Ctrl+K Ctrl+F | 格式化選取的文字 |
| F12 | 跳至定義（變數、函式宣告處） |
| Alt+F12 | 查看定義 |
| Shift+Alt+. | 自動修正 (`Auto Fix...`) |
| Ctrl+Alt+. | 顯示修改建議 (`Quick Fix...`) |
| Shift+F12 | 顯示所有參考（誰用到這個變數／函式？） |
| F2 | 更改符號名稱（變數、函式等等） |

> **註：** 顯示修改建議（`Quick Fix...`）的預設按鍵是 `Ctrl+.`，但此按鍵與微軟注音輸入法衝突，故我把它改成 `Ctrl+Alt+.`。我更常用 Shift+Alt+.（`Auto Fix...`），因為此功能經常跟 `Quick Fix...` 一樣會顯示修改建議。

## Integrated terminal

| 快捷鍵                  | 用途說明 |
|-------------------------|----------------------|
| Ctrl+` | 顯示 integrated terminal |
| Ctrl+Shift+` | 開一個新的 terminal |

## See also

- [官方 Visual Studio Code 快速鍵一覽表](https://blog.poychang.net/vscode-shortcuts/) by Poy Chang
- [Visual Studio Code 常用的快捷鍵](https://wyatthoho.medium.com/visual-studio-code-%E5%B8%B8%E7%94%A8%E7%9A%84%E5%BF%AB%E6%8D%B7%E9%8D%B5-894ff940a2c1) by wyatthoho
