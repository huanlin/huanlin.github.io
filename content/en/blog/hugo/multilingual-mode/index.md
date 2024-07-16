
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
