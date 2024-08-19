---
title: Docsy - Show The Last Modified Date
date: 2023-08-11
slug: docsy-show-last-modified-date
tags: ["hugo", "docsy"]
---

## Background

Hugo supports two ways to get the last modified date of posts:

- Using the last commit date in Git.
- Using the `lastmod` property in the frontmatter of posts.

To use the Git way, we need to enable it in the Hugo configuration file `hugo.toml`:

```toml
enableGitInfo = true
```

With the above configuration, it seems we can have both. That means if we defined `date` or `lastmod` property in the frontmatter of a post, Hugo is supposed to use them, otherwise Hugo will use Git information instead.

However, from my test, Hugo always uses Git information for the last modifition date of posts.

## The Issue

In the Docsy theme, the last modified date is processed in `/layouts/partials/page-meta-lastmod.html`, shown below:

```go
{{ if and (.GitInfo) (.Site.Params.github_repo) -}}
<div class="text-muted mt-5 pt-3 border-top">
  {{ T "post_last_mod" }} {{ .Lastmod.Format .Site.Params.time_format_default -}}
  {{ with .GitInfo }}: {{/* Trim WS */ -}}
    <a href="{{ $.Site.Params.github_repo }}/commit/{{ .Hash }}">
      {{- .Subject }} ({{ .AbbreviatedHash }}) {{- /* Trim WS */ -}}
    </a>
  {{- end }}
</div>
{{ end -}}
```

It doesn't work for me. The last modified date never shows on my posts because my GitHub repo is private and I didn't provide the repo's URL in `hugo.toml`.

## How I Do It

Here is my customized version of `/layouts/partials/page-meta-lastmod.html`:

```go
{{ if .Lastmod }}  
  <div class="text-muted pt-3 border-top">
    {{ T "post_last_mod" }}: {{ .Lastmod.Format .Site.Params.time_format_default }}
  </div>
{{ else if .PublishDate }} 
  <div class="text-muted pt-3 border-top"> 
    {{ T "post_last_mod" }}: {{ .PublishDate.Format .Site.Params.time_format_default }}
  </div>    
{{ end }}
```

I know the code isn't pretty. It just works.

Now if Hugo can fetch the last commit date of a post from my Git repo, the `.Lastmod` will have a value, hence displayed at the bottom of my posts. Or, if anything goes wrong with the `.Lastmod` variable, I can still use the `date` property in the frontmatter.

**See also:** Hugo document > [Configure front matter](https://gohugo.io/getting-started/configuration/#configure-front-matter)