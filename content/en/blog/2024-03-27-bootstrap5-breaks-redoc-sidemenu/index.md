---
title: Redoc side menu woes with Bootstrap 5.x
slug: "redoc-side-menu-woes-with-bootstrap-5"
date: 2024-03-27
tags: [Hugo, Docsy, Redoc, OpenAPI]
---

## Symptom

In a static website built with Hugo and Docsy theme, I add a page to render an OpenAPI yaml file using [Redoc](https://redocly.com/docs/redoc/). The markdown file of the page is as simple as below:

```markdown
---
title: API reference
---

<redoc spec-url="http://petstore.swagger.io/v2/swagger.json"></redoc>
<script src="https://cdn.redoc.ly/redoc/latest/bundles/redoc.standalone.js"> </script>
```

The result page should work like [the official Redoc demo](https://redocly.github.io/redoc/), that is when an item in the left side menu is clicked, the content area should scroll to the corresponding section. However, when I click the side menu, the content area doesn't scroll at all.

> I should create a repository on GitHub to show this issue in action. But I'm pretty tired by the time of writing. Maybe some other day.

Tools and versions:

- [Redoc 2.1.3](https://github.com/Redocly/redoc)
- [Docsy 0.9.0](https://github.com/google/docsy)
- [Bootstrap 5.3.1](https://github.com/twbs/bootstrap)

## Investigation

I'm not proficient in front-end technologies, so I used a rudimentary method to find the cause -- I removed some code from Docsy, bit by bit, until the side menu's scrolling behavior works.

It took me almost an entire day just to find that the issue could be resolved by removing the following code from Docsy's `/asset/scss/main.scss`:

```scss
@import "../vendor/bootstrap/scss/bootstrap";
@import "support/bootstrap_vers_test";
```

Note that when the above code is removed, Hugo will fail to build the website because there are many modules depending on it. Therefore, there are more lines need to be removed in the same file. [Click here to see more code on GitHub](https://github.com/google/docsy/blob/v0.9.0/assets/scss/main.scss#L8-#L9)

At the end of the day, I managed to remove Bootstrap and related code. The website was built successfuly and the side menu in the Redoc page worked normally.

### Possible related issues

After I came up with my workaround and went home, I searched Google and found that the following issues seems relevant:

- Docsy issue #1628: [Redoc left navigation auto scrolling and link changing on scrolling is not working anymore](https://github.com/google/docsy/issues/1628)
- Redoc issue #1235: [Left navigation auto scrolling and link changing on scrolling is not working](https://github.com/Redocly/redoc/issues/1235)
- Redoc issue #1987 [Redoc sidemenu doesn't auto-activate for certain styled outer container](https://github.com/Redocly/redoc/issues/1987)

## Solution

As mentioned in the last section, creating a layout dedicated for Redoc could be a workaround, just like [the Swagger layout in Docsy](https://github.com/google/docsy/tree/v0.9.0/layouts/swagger), without including Bootstrap components in the page.

However, it's just a workaround, not a perfect solution. I'll come back and update this post if I find something new.

Until then.

## Demo

[Demo 1](/docs/demo/redoc)