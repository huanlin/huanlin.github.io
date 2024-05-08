---
title: "Center images with URI fragment"
date: "2024-05-07"
slug: center-images-with-uri-fragment
tags: ["Hugo", "Docsy"]
---

{{< admonition type=warning title="Update" >}}
**2024-05-08:** After I [post a question](https://discourse.gohugo.io/t/the-embeded-image-render-hook-removes-uri-fragment-of-src-attribute/49684) on the Hugo discourse site, Joe Mooring replied and confirmed it is a bug in Hugo. He also created a [ticket](https://github.com/gohugoio/hugo/issues/12468) and a [pull request](https://github.com/gohugoio/hugo/pull/12469) for it. Therefore, this post will be outdated when his pull request is merged and released.
{{< /admonition >}}

## How

My website is built with Hugo and Docsy, and I use a [URI fragment](https://en.wikipedia.org/wiki/URI_fragment) `#center` to indicate that an image should be horizontally centered. Here is an example:

```markdown
![](images/figuer-1.png#center)
```

In the [_styles_project.scss file](https://www.docsy.dev/docs/adding-content/lookandfeel/#project-style-files), I have the following CSS with the "Attribute Ends With" selector `$`:

```css
img[src$="#center"] {
    display: block;
    margin: 1.0rem auto;
    max-width: 100%;
    height: auto;
}
```

## The issue

Using the above appoach, <mark>images are not centered anymore with Hugo v0.124.x and v0.125.6.</mark>

After some tests, I've found that it's because [Hugo's built-in image hook](https://gohugo.io/render-hooks/images/) removed the URI fragment `#center` when converting markdown to HTML.

The generated HTML should look something like below:

```HTML
<img src="images/figure-1.png#center">
```

But the result is:

```HTML
<img src="images/figure-1.png">
```

I've tried two approaches that can solve this issue. One is to disable the default image hook, and another is to write a custom image hook.

## Approach 1: Disable the default image book

According to the Hugo document: [Image render hooks](https://gohugo.io/render-hooks/images/), we can disable the default image hook in the site configuration file:

```toml
[markup]
  [markup.goldmark]
    [markup.goldmark.renderHooks]
      [markup.goldmark.renderHooks.image]
        enableDefault = false
```

Once the default image hook is disabled, the URI fragment `#cener` is correctly render in the result HTML, hence the image is centered.

## Approach 2: Custom image hook

To learn more about Hugo, I've tried to write a custom image hook to solve this issue. Here is how I do it.

Download the source code of the [embeded image render hook](https://github.com/gohugoio/hugo/blob/master/tpl/tplimpl/embedded/templates/_default/_markup/render-image.html) and save it as `/layouts/_defaul/_markup/render-image.html`. Then modify the file content as below:

```go {linenos=table, hl_lines=["6-12"]}
{{- $u := urls.Parse .Destination -}}
{{- $src := $u.String -}}
{{- if not $u.IsAbs -}}
  {{- with or (.PageInner.Resources.Get $u.Path) (resources.Get $u.Path) -}}
    {{- $src = .RelPermalink -}}
    {{/* keep the URI fragment "#center" */}}
    {{- with $u.RawQuery -}}
      {{- $src = printf "%s?%s" $src . -}}
    {{- end -}}
    {{- with $u.Fragment -}}
      {{- $src = printf "%s#%s" $src . -}}
    {{- end -}}
  {{- end -}}
{{- end -}}
{{- $attributes := merge .Attributes (dict "alt" .Text "src" $src "title" (.Title | transform.HTMLEscape)) -}}
<img
  {{- range $k, $v := $attributes -}}
    {{- if $v -}}
      {{- printf " %s=%q" $k $v | safeHTMLAttr -}}
    {{- end -}}
  {{- end -}}>
{{- /**/ -}}
```

Keep writing!

## See also

- [The embeded image render hook removes URI fragment of src attribute](https://discourse.gohugo.io/t/the-embeded-image-render-hook-removes-uri-fragment-of-src-attribute/49684)
- [Hugo embeded link render hook](https://github.com/gohugoio/hugo/blob/master/tpl/tplimpl/embedded/templates/_default/_markup/render-link.html)