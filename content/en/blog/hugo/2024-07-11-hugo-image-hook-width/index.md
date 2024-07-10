---
title: "Hugo image render hook with width parameter"
date: "2024-07-10"
slug: center-images-with-uri-fragment
tags: ["Hugo"]
---

Hugo version: 0.128.2

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

Add a file "render-image.html" under the website's `/layouts/_default/_markup/` folder. Add the following code to the file:

```hugo
{{- $u := urls.Parse .Destination -}}
{{- $src := $u.String -}}
{{- if not $u.IsAbs -}}
  {{- $path := strings.TrimPrefix "./" $u.Path }}
  {{- with or (.PageInner.Resources.Get $path) (resources.Get $path) -}}
    {{- $src = .RelPermalink -}}
    {{/* Commented because there is no need to keep the query string in the src attribute
    {{- with $u.RawQuery -}}
      {{- $src = printf "%s?%s" $src . -}}
    {{- end -}}
    */}}
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

The above code is modified from the [Hugo embedded image render hook](https://github.com/gohugoio/hugo/blob/master/tpl/tplimpl/embedded/templates/_default/_markup/render-image.html).

## See also

- [Center images with URL fragment]({{< ref "../2024-05-07-hugo-image-center/index.md" >}})
- [Hugo embeded link render hook](https://github.com/gohugoio/hugo/blob/master/tpl/tplimpl/embedded/templates/_default/_markup/render-link.html)