---
title: ML.NET 簡介
linkTitle: ML.NET 簡介
---

## 什麼是 ML.NET

ML.NET 是由 Microsoft 開發的免費、開源、且跨平台的機器學習框架。透過 ML.NET，開發人員只需具備一點 ML 基礎概念和演算法，便可建構自訂的 ML 模型，並將其整合至 .NET 應用程式中。

這裡的「模型」指的是一個複雜的數位結構，該結構使我們能夠傳入一些資訊，並獲取輸出結果。例如，輸入的資訊可能是數位影像，輸出則可能是一段描述影像內容的文字。模型內部可能包含多種資料結構和數學公式，而其確切組成會在訓練過程中借助 ML.NET 之類的工具逐漸成形。

ML.NET 支援多種機器學習場景，包括二元分類（binary classification）、多類別分類（multi-class classification）、回歸（regression）、群集（clustering）、異常偵測（anomaly detection）和推薦（recommendation）。只需執行一個簡單的命令和幾個參數，便可針對一些常見的場景進行訓練，例如圖片分類和情感分析。

> 之後再找時間介紹上述幾個專有名詞。

ML.NET 的一個主要優點是它與 .NET 生態系統的整合，包括 Visual Studio 和 Azure Cloud Platform 等受歡迎的工具。這使得 .NET 開發人員能夠輕鬆地將機器學習納入現有的工作流程和應用程式中。不過，Azure 和 Visual Studio 並非必要條件，因為 ML.NET 也有提供方便的命令列工具（CLI），而這工具是跨平台的，可安裝運行於 Windows、Mac 或 Linux 作業環境。

ML.NET 的目標是提供類似 Visual Model Builder 那樣方便好用的功能，讓使用者可以輕鬆創建和訓練機器學習模型，而無需深入了解機器學習的演算法或複雜知識。

## 為何要學 ML.NET？

ChatGPT 和 GitHub Copilot 都是機器學習（以下簡稱 ML）的典型應用，它們廣受歡迎，並成功展示了機器學習的巨大潛力。ChatGPT 是由 OpenAI 開發的語言模型，利用深度學習技術來生成類似人類的自然語言回應。自 2022 年發佈以來，它引起了極大關注，部分原因在於它能夠針對各種提問產生連貫且具備上下文感知（context-aware）的回答。

另一方面，GitHub Copilot 是由 GitHub 和 OpenAI 開發的 AI 助理，用於提升開發人員的編程效率。它會分析程式碼片段、上下文、以及開發人員的歷史記錄和偏好來提供建議。微軟的 Visual Studio 也有內建一個類似的功能，叫做 Intellicode。

然而，企業的競爭優勢來自於有能力建構自己專屬的 ML 模型，而無需依賴第三方解決方案。其原因有二：首先，前面提到的幾個常見的 ML 工具主要是為了解決一些通用的問題，而不擅長處理特定行業或商業領域的需求。其次，出於商業機密和資訊安全的考量，企業內部的敏感資訊不能上傳至第三方平台，以免隨著網際網路四處傳播（這也是為什麼有些企業不允許、或有限度開放員工使用 ChatGPT 的主要原因）。

有了 ML.NET，熟悉 C# 的開發人員便可針對企業內部的需求來打造專屬的 ML 應用程式。

## ML.NET 的強大之處



## Reference

- "Machine Learning for C# Developers" by Fiodar Sazanavets