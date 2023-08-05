---
title: MkDocs vs. Hugo
date: 2023-7-23
---

## 檔案結構

對我來說，MkDocs 與 Hugo 在檔案結構方面的一個差異是圖片的存放位置：

- MkDocs 的圖片可以存放在 Markdown 文件所在資料夾或其下的子目錄，也可以統一放在根目錄之下的某個子目錄。
- Hugo 雖然也是以 Markdown 語法，但是它在插入圖片的時候有一些特殊規必須了解，否則會摸不著頭腦，搞不清楚為何在 Markdown 文件中插入的圖片無法在網頁上顯示。

由於上述差異，編寫 MkDocs 文件時，圖片的存放位置比較直觀，沒有任何特殊之處。例如以下目錄結構：

```
docs\
    docker\
        images\
            figure1.png
            figure2.png
        docker-overview.md    
```

在 docker-overview.md 中插入圖片 images/figure1.png 的寫法是：

```
![](images/figure1.png)
```

另一方面，Hugo 的圖片存放位置有以下規則：

- 對於共用的圖片檔案，應放在專門用來[存放靜態檔案的資料夾](https://gohugo.io/content-management/static-files/)，例如根目錄下的 `static` 或 `asset` 子目錄。
- 如果圖片要跟文件放在同一個資料夾下，或者放在文件所在位置的子目錄下（例如 images），那麼該文件的檔案名稱就必須是 index.md 或 _index.md。在 Hugo 中，若資料夾裡面有 `index.md` 檔案，該資料夾會被視為**樹葉節點（Leave Bundle）**；若資料夾裡面有 `_index.md` 檔案，該資料夾會被視為**分支節點（Branch Bundle）**。總之，想要把圖片和文件放在一起，就必須把文件放在樹葉或樹枝節點裡面，而文件的檔名必須是 index.md 或 _index.md。詳情可參考官方文件的 [Page bundles](https://gohugo.io/content-management/page-bundles/)。

若選擇把圖片放在 `static`` 資料夾，目錄結構看起來會像這樣：

```
\
    static\
        images\
            figure1.png
            figure2.png
    content\            
        docs\
            docker\
                docker-overview.md    
```

在上述結構中，欲在 docker-overview.md 檔案中插入圖片的寫法是：

```
![](images/figure1.png)
```

注意圖片的路徑前面並不需要加上 `static/`，因為 Hugo 會自動去 `static` 目錄下尋找圖片，並於生成 HTML 的時候使用正確的圖片連結。

此寫法有個缺點：市面上大部分的 Markdown 編輯器的預覽功能將無法找到圖片，因而在預覽視窗中無法顯示圖片。基於這個緣故，我偏好把圖片放在文件所在的目錄或 `images` 子目錄中。

