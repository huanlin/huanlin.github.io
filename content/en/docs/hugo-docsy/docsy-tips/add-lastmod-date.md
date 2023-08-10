---
title: Show The Last Modified Date
date: "2023-08-10" 
tags: ["hugo", "docsy"]
---

## Background

Hugo supports two ways to get the last modified date of posts:

- Using the last commit date in Git.
- Using the `lastmod` property in the frontmatter of posts.

To use the Git way, you need to enable it in your Hugo configuration file `hugo.toml`:

```toml
enableGitInfo = true
```

You can have both though. For example, use the date information from Git first, and if Git is disabled, use `lastmod` property instead.

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

As you can see, Docsy only read date information from a Git repository. I want the `lastmod` date to be shown on my posts.

## How I Do It

Thanks to the post on Make with Hugo: [Add a Last Edited Date to Posts](https://makewithhugo.com/add-a-last-edited-date/), I used the code and combined them with Docsy's template.

Here is the result in my `/layouts/partials/page-meta-lastmod.html`:

```go
<!-- Created Date -->
{{ $pubdate := .PublishDate.Format "2006-01-02" }}
Created: 
<time datetime="{{ .PublishDate }}" title="{{ .PublishDate }}">
    {{ $pubdate }}
</time>

{{ if .Lastmod }}
    {{ $lastmod := .Lastmod.Format "2006-01-02" }}
    {{ if ne $lastmod $pubdate }}
        <span class="post-info-last-mod">
            &nbsp; (Updated: 
            <time datetime="{{ .Lastmod }}" title="{{ .Lastmod }}">
                {{ $lastmod }}
            </time>)
          </span>
    {{ end }}
{{ else if and (.GitInfo) (.Site.Params.github_repo) }}
<div class="text-muted mt-5 pt-3 border-top">
  {{ T "post_last_mod" }} {{ .Lastmod.Format .Site.Params.time_format_default }}
  {{ with .GitInfo }}: {{- /* Trim WS */ -}}
    <a href="{{ $.Site.Params.github_repo }}/commit/{{ .Hash }}">
      {{ .Subject }} ({{ .AbbreviatedHash }}) {{- /* Trim WS */ -}}
    </a>
  {{ end }}
</div>
{{ end }}
```

Now if I have `date` or `lastmod` defined in the frontmatter of a post, the date information will be displayed at the bottom of the post. Job done.

**See also:** Hugo document > [Configure front matter](https://gohugo.io/getting-started/configuration/#configure-front-matter)