
In Hugo configuration file:

```
# Language settings
contentDir = "content/en"
defaultContentLanguage = "en"

# 注意：defaultContentLanguageInSubdir 不可為 true! 否則搜尋功能會 404 Error。
# 參見：https://github.com/google/docsy/issues/1646
defaultContentLanguageInSubdir = false

[languages]
[languages.en]
  languageName ="English"
  weight = 1  # Weight used for sorting.
[languages.en.params]
  title = "ICSD Documentation"
  description = "Documentation site"

[languages.zh]
  languageName = "中文 (Chinese)"
  Title = "ICSD 文件"
  weight = 2
  #contentDir = "content/zh"
[languages.zh.params]
  subtitle = "文件網站"
  time_format_default = "2006-01-02"
  time_format_blog = "2006-01-02"

[module]
  [[module.mounts]]
    source = 'content/en'
    target = 'content'
    lang = 'en'
  [[module.mounts]]
    source = 'content/en' # Use content/en as the fallback folder to find pages when the selected languae is 'zh'
    target = 'content/'
    lang = 'zh'
```    

使用檔名模式時，檔案名稱必須使用符合組態檔案中的語言代號——只有預設的語言除外。例如，預設的語言是 `en`，而 `languages.zh` 是用來決定中文版的相關參數，那麼每一個中文版的檔案就必須是 `*.zh.md`。如果是英文版文件，由於英文是預設語言，故檔案名稱有沒有加上 `en` 皆可，亦即 `*.md` 和 `*.en.md` 皆可。方便起見，我對於預設語言的文件都不特別加上語言代號。

要提醒的是，對於非預設語系的文件，如果命名檔案時不小心打錯字，應該是 `*.zh.md` 卻命名為 `*.zh-Hant.md`，由於 `zh-Hant` 沒有在組態檔案中定義，這會造成一些奇怪的狀況或者錯誤。在撰寫此文件時， 如果有這種情形，Hugo 在編譯網站檔案的時候會先比平常花更多時間，並以錯誤訊息中止：

```console
ERROR render of "404" failed: 
  "...\layouts\_default\baseof.html:8:9": execute of template failed: ....
ERROR render of "page" failed: 
  "...\themes\...\baseof.html:8:9": execute of template failed...
...
Built in 61749 ms
Error: error building site: render: 
  "...\layouts\_default\baseof.html:8:9": ... timed out after 30s. 
  This is most likely due to infinite recursion....
```

可以看到上面的錯誤訊息顯示 Hugo 建置網站的時間超過 61 秒，最終以「逾時」（timeout）中止。
