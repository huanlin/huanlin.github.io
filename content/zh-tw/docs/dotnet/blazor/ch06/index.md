---
title: 第 6 章：表單與欄位驗證 Part 1
weight: 6
---

## 使用元件來建構表單

在 Blazor 應用程式中雖然可以直接使用標準的 HTML 表單來讓使用者輸入資料，但此作法有個缺點：沒有方便好用的欄位驗證機制。因此，Blazor 的表單元件是比較好的選擇。

主要的元件叫做 `EditForm`。

`EditForm` 內部會建立一個名為 `EditContext` 的物件。``EditContext` 是整個表單元件系統的中樞，它會追蹤表單內的所有輸入元件，並記錄資料模型的狀態。當資料模型有欄位值變動，它會觸發欄位驗證事件。

> 這裡的「資料模型」用英文來說就是 model，實際上就是一個 C# 類別，而類別中的屬性是對應（綁定）至表單的各個輸入欄位。討論相關議題時，為求簡潔，我偶爾會說「模型」。

Blazor 提供的欄位驗證器叫做 `DataAnnotationsValidator`，它讓我們能夠使用 Data Annotations 的方式撰寫欄位的驗證規則。

`EditForm` 提供了三個事件來處理表單提交的動作：`OnSubmit`、`OnValidSubmit`、`OnInvalidSubmit`。

- `OnSubmit`：跟標準 HTML 表單的提交（submit）事件沒有兩樣，每當使用者點擊 Submit 按鈕，便會觸發此事件。
- `OnValidSubmit`：當使用者點擊 Submit 按鈕，`EditForm` 會先透過 `EditContext` 確認表單的模型是否都能通過欄位驗證檢查；只有當模型皆通過驗證才會觸發此事件。這這表示我們通常不需要在此事件的處理常式中撰寫欄位驗證的程式碼。
- `OnInvalidSubmit`：呈上，只有當表單的模型未能通過驗證才會觸發此事件。

### 建立模型

### `EditForm` 基礎設定

### 蒐集輸入元件的資料


## 驗證模型

## 提交資料至伺服器

