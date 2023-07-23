# MkDocs vs. Hugo


## 檔案結構

對我來說，MkDocs 與 Hugo 在檔案結構方面的一個重要差異是圖片的存放位置：

- MkDocs 的圖片可以存放在 Markdown 文件所在資料夾或其下的子目錄，也可以統一放在根目錄之下的某個子目錄。
- Hugo 的圖片則有一些規則必須了解，否則很容易碰到在 Markdown 文件中插入的圖片無法在網頁上顯示。

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

- 需要共用的圖片，要把圖片放在專門用來[存放靜態檔案的資料夾](https://gohugo.io/content-management/static-files/)。
- 如果圖片要跟文件放在同一個資料夾下，或者放在文件的子目錄下（例如 images），那麼該文件的檔案名稱就必須是 index.md 或 _index.md。在 Hugo 中，若資料夾裡面有 index.md 檔案，該資料夾會被視為樹葉節點（Leave Bundle）；若資料夾裡面有 _index.md 檔案，該資料夾會被視為分支節點（Branch Bundle）。總之，想要把圖片和文件放在一起，就必須把文件放在樹葉或樹枝節點裡面，即檔名是 index.md 或 _index.md。詳情可參考官方文件的 [Page bundles](https://gohugo.io/content-management/page-bundles/)。

如果選擇把圖片要放在靜態檔案資料夾，目錄結構看起來會像這樣：

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

在 docker-overview.md 中插入圖片的寫法是：

```
![](images/figure1.png)
```

這種寫法有個缺點：Markdown 編輯器的預覽功能將無法得知圖片的正確路徑，因而無法預覽圖片。因此，我個人偏好把圖片放在文件所在的目錄或子目錄中。
