---
title: "Hugo Tip & Tricks"
date: "2023-08-09"
description: >
  Some tips & tricks that I found useful with Hugo.
tags: ["hugo"]
weight: 1
---

## Shortcodes

### Escaping Shortcodes

Sometimes I need to demonstrate how to use a shortcode in my article. Simply put the shortcode example in a markdown code block won't work because the shortcode will be interpreted and executed anyway.

The solution is to add a pair of `/*` and `*/` in both the beginning and end lines of the shortcode block. Here is an example:

![](images/escaping-shortcode.png)

It is rendered like this:

```
{{%/* admonition type=note title="This is a note" */%}}
It's not who you are underneath, it's what you do that defines you.
{{%/* /admonition */%}}
```



