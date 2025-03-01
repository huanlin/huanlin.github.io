---
title: Hugo Blockquote Render Hook
date: 2024-09-09
tags: ["Hugo"]
---

Hugo supports [blockquote render hooks](https://gohugo.io/render-hooks/blockquotes/) since v0.132.0, which makes it easy to create beautiful alert boxes with the simple blockquote syntax. This is great because I don't need to use shortcodes for creating alert boxes anymore.

Here are supported alert types:

- `NOTE`
- `TIP`
- `WARNING`
- `IMPORTANT`
- `CAUTION`
- `QUOTE`

P.S. They can be written with lowercase.

## Examples

### Note

```text
> [!note]
> This is a note.
```

> [!note]
> It's not who I'm underneath, but what I do, that defines me.

### Tip

```text
> [!tip]
> It's not who I'm underneath, but what I do, that defines me.
```

> [!tip]
> It's not who I'm underneath, but what I do, that defines me.

### Important

```text
> [!important]
> It's not who I'm underneath, but what I do, that defines me.
```

> [!important]
> It's not who I'm underneath, but what I do, that defines me.

### Warning

```text
> [!warning]
> It's not who I'm underneath, but what I do, that defines me.
```

> [!warning]
> It's not who I'm underneath, but what I do, that defines me.

### Caution

```text
> [!caution]
> It's not who I'm underneath, but what I do, that defines me.
```

> [!caution]
> It's not who I'm underneath, but what I do, that defines me.

### Quote

```text
> [!quote]
> It's not who I'm underneath, but what I do, that defines me.
```

> [!quote]
> It's not who I'm underneath, but what I do, that defines me.

## Implementation

### render-blockquote-alert.html

**File:** `/layouts/_default/_markup/render-blockquote-alert.html`

```html
{{- $iconMap := dict "note" "fas fa-pencil-alt fa-fw" -}}
{{- $iconMap  = dict "info" "fas fa-info-circle fa-fw" | merge $iconMap -}}
{{- $iconMap  = dict "tip" "fas fa-lightbulb fa-fw" | merge $iconMap -}}
{{- $iconMap  = dict "question" "fas fa-question-circle fa-fw" | merge $iconMap -}}
{{- $iconMap  = dict "warning" "fas fa-exclamation-triangle fa-fw" | merge $iconMap -}}
{{- $iconMap  = dict "caution" "fas fa-exclamation-triangle fa-fw" | merge $iconMap -}}
{{- $iconMap  = dict "important" "fa-solid fa-star fa-fw" | merge $iconMap -}}
{{- $iconMap  = dict "failure" "fas fa-times-circle fa-fw" | merge $iconMap -}}
{{- $iconMap  = dict "danger" "fas fa-skull-crossbones fa-fw" | merge $iconMap -}}
{{- $iconMap  = dict "bug" "fas fa-bug fa-fw" | merge $iconMap -}}
{{- $iconMap  = dict "example" "fas fa-list-ol fa-fw" | merge $iconMap -}}
{{- $iconMap  = dict "quote" "fas fa-quote-right fa-fw" | merge $iconMap -}}
{{- $iconDetails := "fas fa-angle-right fa-fw" -}}
<!--
  Modified from https://github.com/HEIGE-PCloud/DoIt/blob/main/layouts/shortcodes/admonition.html
  2024-09-09: Michael Tsai - Removed Open/Close icon, so it's always displayed as open.
-->

{{- $type := .AlertType | default "note" -}}
<div class="td-max-width-on-larger-screens details admonition {{ $type }}">
    <div class="details-summary admonition-title">
        <i class='icon {{ index $iconMap $type | default (index $iconMap "note") }}'></i>
        {{ with .AlertTitle }}
            {{ . }}
        {{ else }}
            {{ or (i18n .AlertType) (title .AlertType) }}
        {{ end }}
    </div>
    <div class="details-content">
        <div class="admonition-content">
            {{ .Text | safeHTML -}}
        </div>
    </div>
</div>
```

### _admonition_variables.scss

**File:** `/assets/scss/_admonition_variables.scss`

```scss
/*
  Modified from DoIt theme:
  https://github.com/HEIGE-PCloud/DoIt/blob/main/assets/css/_variables.scss
*/

// ========== Admonition ========== //
// Color map of the admonition
$admonition-color-map: (
  'note': #448aff,
  'abstract': #00b0ff,
  'info': #00b8d4,
  'tip': #00bfa5,
  'success': #00c853,
  'question': #64dd17,
  'warning': #ff9100,
  'failure': #ff5252,
  'danger': #ff1744,
  'bug': #f50057,
  'example': #651fff,
  'quote': #9e9e9e,
  'important': #651fff,
  'caution': #ff1744,
) !default;

// Color map of the admonition background
$admonition-background-color-map: (
  'note': rgba(68, 138, 255, 0.1),
  'abstract': rgba(0, 176, 255, 0.1),
  'info': rgba(0, 184, 212, 0.1),
  'tip': rgba(0, 191, 165, 0.1),
  'success': rgba(0, 200, 83, 0.1),
  'question': rgba(100, 221, 23, 0.1),
  'warning': rgba(255, 145, 0, 0.1),
  'failure': rgba(255, 82, 82, 0.1),
  'danger': rgba(255, 23, 68, 0.1),
  'bug': rgba(245, 0, 87, 0.1),
  'example': rgba(101, 31, 255, 0.1),
  'quote': rgba(159, 159, 159, 0.1),
  'important': rgba(101, 31, 255, 0.1),
  'caution': rgba(255, 23, 68, 0.1),
) !default;
```

### _admonition.scss

**File:** `/assets/scss/_admonition.scss`

```scss
/*
  Source: https://github.com/HEIGE-PCloud/DoIt/blob/main/assets/css/_partial/_single/_admonition.scss
  Modified by Michael Tsai (2024-09-09), to make it similar to MkDocs Material's admonitions.
*/
.admonition {
  position: relative;
//  margin: 1rem 0;
//  padding: 0 .75rem;
//  background-color: map-get($admonition-background-color-map, 'note');
//  border-left: .25rem solid map-get($admonition-color-map, 'note');
//  overflow: auto;

  background-color: map-get($admonition-background-color-map, 'note');
  border: 0.05rem solid #448aff;
  border-radius: 0.2rem;
  box-shadow: var(--md-shadow-z2);
//  color: var(--md-admonition-fg-color);
  display: flow-root;
  font-size: .87rem;
  margin: 1.5625em 0;
  padding: 0 0.8rem 0.8rem;
  page-break-inside: avoid;


  .admonition-title {
    font-weight: bold;
    margin: 0 -0.75rem;
    padding: .25rem 1.8rem;
    border-bottom: 1px solid map-get($admonition-background-color-map, 'note');
    background-color: opacify(map-get($admonition-background-color-map, 'note'), 0.15);
  }

  &.open .admonition-title {
    background-color: map-get($admonition-background-color-map, 'note');
  }

  .admonition-content {
    padding: .9rem 0 0.1rem;
  }

  // 當 callout 方塊裡面有分段落時，必須調整段落的上下邊界，否則最後一行底下的留白區域會太多。
  .admonition-content p {
    margin-top: 0.2rem;
    margin-bottom: 0rem;
    padding-bottom: 0.2rem;
  }

  .admonition-content code {
    background-color: inherit;
  }

  i.icon {
    font-size: 0.85rem;
    color: map-get($admonition-color-map, 'note');
    position: absolute;
    top: .6rem;
    left: .4rem;
  }

  i.details-icon {
    position: absolute;
    top: .6rem;
    right: .3rem;
  }

  @each $type, $color in $admonition-color-map {
    &.#{$type} {
      border-left-color: $color;

      i.icon {
        color: $color;
      }
    }
  }

  @each $type, $color in $admonition-background-color-map {
    &.#{$type} {
      background-color: $color;

      .admonition-title {
        border-bottom-color: $color;
        background-color: opacify($color, 0.15);
      }

      &.open .admonition-title {
        background-color: $color;
      }
    }
  }

  &:last-child {
    margin-bottom: .75rem;
  }
}
```

### i18n

**File:** /i18n/en.toml

```toml
caution = 'Caution'
important = 'Important'
note = 'Note'
tip = 'Tip'
warning = 'Warning'
quote = "Quote"
bug = "Bug"
danger = "Danger"
```

### Hugo configuration

**File:** /hugo.toml

```toml
[markup.goldmark.parser.attribute]
block = true
```

## Reference

- Hugo Documentation: [Blockquote render hooks](https://gohugo.io/render-hooks/blockquotes/)