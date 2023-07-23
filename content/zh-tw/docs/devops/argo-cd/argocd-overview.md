# Argo CD 簡介

摘要：

先備知識：[GitOps](gitops.md)

## What

Argo CD 是一個基於 Kubernetes 的宣告式 GitOps 工具。

關於 GitOps 與宣告式工具，可參考另一篇文章的說明：[GitOps 簡介](gitops.md)。至於 Kubernetes，則可參考 [Kubernetes 簡介](../k8s/overview/k8s-overview.md)。主要是了解我們的程式會部署到 Pods 中，而 Kubernetes 會根據我們在 YAML 檔案中描述的期望狀態來調整實際作業環境的狀態。

## Why

## Argo CD 的運作方式

Argo CD 遵循 GitOps 模式，使用 Git 儲存庫作為單一資訊來源，以獲取應用程式預期狀態的定義。

Argo CD 本身也是一個 Kubernetes controller，它會持續監測運行中的應用程式，並拿當前的狀態去跟預期的目標狀態（保存於 Git repo）比較。一旦發現當前狀態與目標狀態有差異，Argo CD 能夠以自動或手動的方式來讓應用程式當前的運行狀態回到預期的目標狀態。換言之，在 Git repo 中對目標狀態所做的任何修改都能自動反映至目標環境。

底下是 Argo CD 的架構圖（摘自 [Argo CD 文件](https://argo-cd.readthedocs.io/en/stable/)）：

![](images/argocd-architecture.png)

## 參考資料

- [Argo CD Overview](https://argo-cd.readthedocs.io/en/stable/)
- [What is Argo CD](https://www.youtube.com/watch?v=p-kAqxuJNik)