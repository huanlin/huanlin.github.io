---
title: "2: 建立 Python 和 JavaScript web apps"
draft: true
---

在學習建立和使用容器之前，我們必須先知道如何建立基本的應用程式。

本章用來建立 web 應用程式的兩大基石：

- JavaScript runtime NodeJS
- Python 程式語言 + FastAPI

## Designing a basic FastAPI web app in Python

FastAPI 是一個 Python web app 框架，用來協助開發 web-based 或 API-based 應用程式。稍後會用到的工具如下，依照安裝的順序列出：

- **Python version 3.8+**
- **venv** - A Python module for creating and managing virtual environments for isolating code.
- **FastAPI** - A popular Python web framework that comes with a minimal amount of features to build a highly functional application.
- **uvicorn** — The web server we will use to handle web traffic with FastAPI.

如果您是經驗豐富的 Python 開發人員，可自行選擇慣用的工具。若是 Python 新手，則建議跟著本書按部就班來逐一完成各項工具的安裝與環境設定。

### Python project setup with virtual environments

Python virtual environments：開發 Python 應用程式時，用來隔離 Python 專案的虛擬環境，以免各專案之間因為依賴不同版本的 third-party 套件而互相干擾。

這裡會使用 Python 內建的虛擬環境管理員，叫做 `venv`。

If you’re on macOS or Linux, you will likely use python3 instead of python and if you’re on Windows you will likely use python instead of python3

```shell
cd work/py
python -m venv venv
```

