---
title: "A Hugo image render hook that supports width parameter"
linkTitle: Image hook
date: "2024-07-10"
slug: hugo-image-render-hook-width-param
tags: ["Hugo"]
---

This post shows how I support image width with the URL query parameter `width`.

**Hugo version:** 0.128.2

## Requirement

I want to specify a image's width via the URL query string. For example:

```markdown
![](images/figuer-1.png?width=650px "Figure 1")
```

Rendered HTML:

```html
<img src="/path/to/images/figure-1.png" title="Figure 1" width="650">
```

## Implementation

Add a file "**render-image.html**" under the website's `/layouts/_default/_markup/` folder. Add the following code to the file:

```go
{{- $u := urls.Parse .Destination -}}
{{- $src := $u.String -}}
{{- if not $u.IsAbs -}}
  {{- $path := strings.TrimPrefix "./" $u.Path }}
  {{- with or (.PageInner.Resources.Get $path) (resources.Get $path) -}}
    {{- $src = .RelPermalink -}}
    {{- with $u.RawQuery -}}
      {{- $src = printf "%s?%s" $src . -}}
    {{- end -}}
    {{- with $u.Fragment -}}
      {{- $src = printf "%s#%s" $src . -}}
    {{- end -}}
  {{- end -}}
{{- end -}}

{{- $params := $u.Query -}}
{{- $width := default "" ($params.Get "width" | strings.TrimSuffix "px") }}

{{- $attributes := merge .Attributes
    (dict "alt" .Text "src" $src "title" (.Title | transform.HTMLEscape))
    (dict "width" $width) -}}

<img
    {{- range $k, $v := $attributes -}}
      {{- if $v -}}
        {{- printf " %s=%q" $k $v | safeHTMLAttr -}}
      {{- end -}}
    {{- end -}}>
```

Job done!

The above code is modified from the [Hugo embedded image render hook](https://github.com/gohugoio/hugo/blob/master/tpl/tplimpl/embedded/templates/_default/_markup/render-image.html).

## A note about Hugomods Images

I've also tried [HugoMods Images module](https://images.hugomods.com/). Its image render hook supports URL query parameters such as `width` and `height`. However, it has the following catches:

- 不認識以 "." 開頭的路徑，例如 "./images/..."。採用此寫法的圖片都會無法顯示（被 image render hook 直接略過）。
- 雖然可以用 URL query parameter "width" 指定圖片大小，但卻是真的把圖片縮小了，使用者無法點擊圖片來查看原尺寸的圖片。
- 生成網頁的時候，會額外產生新的圖片檔案（預設是 .webp），額外占用磁碟空間（如以下範例）。

```html
<picture class="d-block text-center">
  <source srcset="/path/to/my-post/images/my-post_hu987...be6b0d89b30645f8.webp" type="image/webp">
  <img class="img-fluid medium-zoom-image"
       src="/path/to/my-post/images/my-post_hu987...e8a_83422_250x0_resize_catmullrom_3.7727d...bd2.png"
       alt="" loading="lazy" height="149" width="250">
</picture>
```

Therefore, I wouldn't recommend HugoMods Images module for technical documentation websites. It could be helpful for other types of blogs.

## See also

- [Hugo - Center images with URL fragment]({{< ref "../2024-05-07-hugo-image-center/index.md" >}})
- [Hugo embeded link render hook](https://github.com/gohugoio/hugo/blob/master/tpl/tplimpl/embedded/templates/_default/_markup/render-link.html)
- [HugoMods Images module's render hook](https://github.com/hugomods/images/blob/main/layouts/_default/_markup/render-image.html)
