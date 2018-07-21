# 使用 VSTS 來自動建置 GitHub repo

按照這份微軟的逐步教學文件：Define a continuous integration (CI) build process for your GitHub repository，我把我在 GitHub 上面的一個 repo 設定成由 VSTS（Visual Studio Team Service） 來進行自動建置，設定的過程相當順利。

> VSTS 是個雲端服務，可免費申請。

那份文件是以 master branch 為例，而我實際操作時，是使用 develop branch。設定完成後，每當那個 GitHub  repo 有下列變動，就會觸發 VSTS 的建置程序：

* develop branch 有 commits 時。
* 有人對 develop branch 建立 pull request 時。

我順手抓了圖，附在本文後面（圖很多）。這裡先說一下我用來設定 CI 的 GitHub repo。

## 我的 GitHub repo

我用來設定 CI 的 GitHub repo 名稱是 [EasyBrailleEdit](https://github.com/huanlin/EasyBrailleEdit)。這個專案是從 2008 年開始，開發人員只有一名（就是我）。它是個 Windows desktop 應用程式，主要是供視覺障礙輔具相關機構或中小學的老師利用這個工具來製作給視障學生使用的點字書或雙視書（即紙上同時印有明眼字和點字）。

它原本是個商用軟體，過去幾年都是跟一家輔具廠商簽約合作，由對方負責銷售和安裝。直到今年，已經是第十個年頭，而且已經一兩年都沒有什麼大幅度的修改。於是，我跟合作廠商提議開放原始碼，如果有人想改善它（例如加上單元測試），那是最好了。不過，這麼多年前寫的程式，肯定有不少嚇人的坑，還是我自己找時間來慢慢重構吧 Orz

## 設定步驟

底下每個小節的標題，都是對應到微軟文件裡面的小節標題，方便兩相對照。

### Set up CI builds for your GitHub repository

前提：已經有一個 VSTS 帳號，而且在裡面建立好專案。我在 VSTS 上面建立的專案名稱是 Braille。

進入 VSTS 專案之後，可以從 Code 頁面看到右邊有個［Set up build］按鈕，如下圖：

![](images/github-ci-with-vsts-01.png)

點一下這個按鈕，會切換到 Build and Release 頁面，讓你開始設定一個新的 Build。

接著按下圖的指示操作：

![](images/github-ci-with-vsts-02.png)

然後會跳出 GitHub 的授權頁面，以便確認要授權 VSTS 存取你的 GitHub repo：

![](images/github-ci-with-vsts-edb706a5.png)

授權成功後，便可以指定要使用哪個 repo，以及哪個 branch：

![](images/github-ci-with-vsts-cf00ccd1.png)

然後要選擇 template。由於我的應用程式類型是 Windows Forms，所以我選擇「.NET Desktop」模板。

![](images/github-ci-with-vsts-1ed2391d.png)

［Agent queue］選擇 "Hosted VS2017"，［Solutio］可以指定某個 .sln 檔案，也可以使用萬用字元：

![](images/github-ci-with-vsts-0b7e17a7.png)

看一下其他步驟的設定頁面（沒有改動任何設定）：

![](images/github-ci-with-vsts-2f92f462.png)

![](images/github-ci-with-vsts-c4f4188f.png)

在 Triggers 頁面啟用此 GitHub repo 的持續整合功能，並指定 branch：

![](images/github-ci-with-vsts-c91aa652.png)

如此一來，每當 develop branch 有新的 commit 時，VSTS 就會自動建置專案。

接著在 Options 頁面設定 GitHub badge：

![](images/github-ci-with-vsts-2bc2183d.png)

點［Save & queue］之後，會彈出視窗讓你確認之後再儲存：

![](images/github-ci-with-vsts-033b66d3.png)

儲存成功後，Options 頁面會出現 badge 連結：

![](images/github-ci-with-vsts-b206dfd0.png)

把這個連結複製起來，接著就會用到。

### Create a VSTS build status with a GitHub README file

這個步驟是要把上一節所產生的 badge 連結加入至 GitHub repo 的 README.md 裡面。

我是把 badge 連結貼到 README.md 的最後一行。這個連結不是直接貼上去就行，還需要加工一下，讓它變成一個 <img> 標籤。其格式如下：

~~~~~~~~
[<img src="https://{your-account}.visualstudio.com/_apis/public/build/definitions/{guid}/{id}/badge"/>](https://{your-account}.visualstudio.com/{your-project}/_build/index?definitionId={id})
~~~~~~~~

底下是範例：

~~~~~~~~
[<img src="https://huanlin.visualstudio.com/_apis/public/build/definitions/4ea34b79-924b-4784-aca6-2668f8014439/1/badge" />](https://huanlin.visualstudio.com/Braille/_build/index?definitionId=1})
~~~~~~~~

修改好之後，儲存 README.md，commit 並且 push 至 GitHub。此時 VSTS 會知道 develop branch 有新的變動，於是便會開始執行建置程序。

過了一會兒，我就收到來自 VSTS 的新郵件，通知我有一項建置工作剛剛完成了。如下圖：

![](images/github-ci-with-vsts-a51d1b4c.png)

查看建置報告：

![](images/github-ci-with-vsts-335606d1.png)

此時到 GitHub 的 repo 頁面查看，可以看到 README.md 裡面，剛剛設定的 badge 顯示一個圖案，圖案裡面有字："build succeeded"：

![](images/github-ci-with-vsts-68448e43.png)

查看此 repo 的 Setting 頁面，可以看到裡面有一個 webhook：

![](images/github-ci-with-vsts-45ea782e.png)

這個 webhook 是由 VSTS 所建立，其作用是讓 GitHub 在每次有 commit 時自動通知 VSTS，以觸發建置程序。

> 先前我們在 VSTS 裡面設定 Build 時指定要使用這個 GitHub repo，當時 VSTS 便已經在 GitHub 建立了這個 webhook。

另外，我們也可以從 VSTS 的建置報告中找到觸發該建置程序的來源 commit：

![](images/github-ci-with-vsts-95d76bda.png)

![](images/github-ci-with-vsts-d2a85add.png)

### Create a pull request trigger for GitHub

在 VSTS 中，編輯先前建立好的 Build：

![](images/github-ci-with-vsts-02a695f2.png)

在 Triggers 頁面啟用 pull request validation：

![](images/github-ci-with-vsts-3c048860.png)

接著就來測試看看，剛才建立的 pull request trigger 有沒有作用。首先，在 GitHub repo 中建立一個新的 branch（基於 develop branch），隨便改一個檔案，commit 並且 push。

然後建立 pull request，請求將這個 branch 合併回 develop branch：

![](images/github-ci-with-vsts-aad64287.png)

剛建立好 pull request，就可以在 VSTS 頁面上看到一項新的建置工作正在進行：

![](images/github-ci-with-vsts-fbf2e41b.png)

另一方面，GitHub 頁面此時也看得到同樣的狀態：

![](images/github-ci-with-vsts-26f5aa71.png)

等到 VSTS 的建置工作完成，確認建置無誤，便可放心合併分支了：

![](images/github-ci-with-vsts-3a0a2eea.png)

呼～大功告成！
