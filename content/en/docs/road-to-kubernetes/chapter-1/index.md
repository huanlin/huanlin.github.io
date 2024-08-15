---
title: "Chapter 1: 概述"
slug: "road-to-kubernetes-start"
draft: true
---

## 每個人都能學會

Kubernetes 和 Docker 有許多術語和一大堆工具，導致初學者很容易一開始就被一大堆名詞弄得暈頭轉向，不知到底從哪個方向切入才能順利跨越入門的那道門檻。

只要先記住一件事：Kubernetes 的核心觀念就是**部署**。

而且，每個人都能學會使用 K8s 部署。我們只是需要找到正確的學習切入點，先學會基礎的觀念和操作，按部就班，逐漸進階。

第二個重要觀念是：Kubernetes 和 Docker 都是**宣告式工具**（declarative tools）。

宣告式工具：用一個文字檔（通常是 YAML 格式）描述你想要達成的結果，然後工具就會幫你達成。注意是描述**結果為何**（what should be done），而不是**如何達成**（how it should be done）。

在 Kubernetes 中，這些用來宣告結果的檔案稱為 **Manifests**。在 Docker 則叫做 **Dockerfiles**。

> IaC (infrastructure-as-code) 工具如 Ansible 和 Terraform 也是宣告式工具。

## Our path to deployment

本書的學習步驟：

1. 用 Python 和 JavaScript 寫一兩個小程式。
2. 把程式的原始碼放到 GitHub 平台來做版本控制，並使用 GitHub Actions 來執行自動整合與部署。
3. 加入容器技術。我們會實作一個整合流程來為我們的程式自動建立一個 Docker 容器。完成後，再將容器上傳至 DockerHub。
4. 一旦把容器上傳至 DockerHub，我們可以進入下一個階段：container-based deployment。
5. 現在，我們應該會逐漸體會一次要部署多個容器的麻煩，並開始思考如何提升容器應用程式的效率、資源的配置等等。為了處理這些議題，開始加入 Kubernetes。

我們將會了解，使用雲端平台托管的（managed） Kubernetes 服務來執行一個單純的 Kubernetes 部署任務是很簡單的。比較麻煩的是加入更多可能需要或不需要存取 internet 資源的服務。

本書的目的就是要幫助你解決這一路上出現的許多挑戰。

## Why K8s? 有 Docker 還不夠嗎？

基於容器的部署方式有其限制。當我們需要一次部署多個容器時，儘管可以使用 Docker Compose 和 Watchtower 等工具，但是版本的升級、負載管理、資源配置等工作還是會令 DevOps 工程師焦頭爛額。Kubernetes 可以減輕這方面的負擔。

關於版本升級的問題，其根源在於<mark>管理 dependencies 的複雜性</mark>。

## 容器是用來幹嘛？

容器可以將你的軟體及其依賴的套件全部捆成一包，以便在更容易部署、共享、和運行於多種運算裝置。換句話說，容器可讓你的軟體具備可移植性。

當我們說建置容器、共享容器、或運行容器，通常是在說 **Docker**。但 Docker 並非唯一的容器工具。許多開放原始碼專案都有支援容器，而只要能使用容器，就能夠使用 Kubernetes。

## 沒有 Kubernetes 會怎樣？

剛開始建置容器 image 的過程相當簡單，麻煩的是要找出所有的 dependencies，包括 OS 層級和應用程式層級的依賴。一旦你能確認軟體依賴哪些相關套件或執行環境，就可以開始建立容器的組態檔，即所謂的 Dockerfile。Dockerfile 的內容可以非常單純，也可以搞到超級複雜。

Dockerfile 通常會導致錯誤的信心，因為只有等到你在建置和執行容器的時候才會發現 bugs。即使建置成功，運行時也可能出錯；即使第一次運行順利成功，後來仍有可能出錯。

Here’s a few highlights of challenges that we will likely face when using just containers without Kubernetes:

- Updating or upgrading container versions
- Managing three or more containers at once
- Scaling container-based apps and workloads
- Failure recovery
- Container-to-container communication (interoperability)

以下是沒有使用 Kubernetes、僅使用容器的情況下可能面臨的挑戰：

- 更新或升級容器的版本
- 同時管理三個或更多容器
- 擴展（scaling）基於容器的應用程式和 workloads
- 故障復原
- 容器間的通訊（互通性；interoperability）

## 使用 Kubernetes 就從此幸福快樂嗎？

雖然 Kubernetes 能夠協助解決一些部署和管理容器的麻煩，但隨之而來的挑戰還是不少，例如：

- Kubernetes 本身也需要安裝和更新（所以採用 managed Kubernetes 會比較輕鬆）
- 組態爆炸（configuration overload）
- Stateful vs. stateless applications
- Security and permissions
- 容器間的通訊
- Storage
- Recovery and debugging

一旦成功運行 Kubernetes，接著要面對的就是一堆組態檔的更新與維護。即使只是更新一個應用程式，我們也得學習如何使用 CI/CD 工具，並了解何謂 role-based access control (RBAC)，才有辦法順利將變更推送至我們的 Kubernetes cluster。

容器本身的設計是無狀態的（stateless），這表示資料庫應用程式這類需要保存狀態的使用情境會牽涉到更複雜的技術。

總之，沒有萬靈丹。

## 學習方式

1. 手動部署
2. 使用容器來部署
3. 使用 Kubernetes 來部署

一定要先學會手動部署，這很重要，因為你才會知道自動化部署的背後其實都做了哪些事情。要是你在一知半解的情況下完成自動化部署，那麼將來一旦出狀況，你就會感到相當棘手，不知如何解決。而且，任何東西遲早都會出狀況。

