---
title: Keep Sidebar Menu's Scroll Position
date: "2023-08-08"
description: >
  Docsy and many other Hugo themes have the same UI issue. Here is a solution.
tags: ["hugo", "docsy"]
lastmod: "2023-08-09"
---

## The Issue

I'll save my time describing the issue. Just read the following Docsy issues on GitHub:

- [#257 - Sidebar menu reloads to the top and user has to scroll down to find an item](https://github.com/google/docsy/issues/257) 
- [#348 - Left-side menu / TOC jumps around after you click a menu](https://github.com/google/docsy/issues/348)

Both issues are created in 2020. They remain unresolved while I'm writing this note right now, and it is August 2023!

## How I Solve It

Note the following solution can only be applied to Hugo with Docsy theme.

First, add a file `body-end.html` in the folder `/layouts/partials/hooks/` of your website. This is a customization mechanism provided by Docsy theme.

Then add the following code to `body-end.html`:

<script > 
  (function() {
    var a = document.querySelector("#td-section-nav");
    addEventListener("beforeunload", function(b) {
        localStorage.setItem("menu.scrollTop", a.scrollTop)
    }), a.scrollTop = localStorage.getItem("menu.scrollTop")
  })()
</script>

Job done.

One final note: the id `td-section-nav` is a `nav` element used in Docsy for the sidebar menu. You can find this element with Chrome DevTools. It looks like:

```html
<nav class="collapse td-sidebar-nav" id="td-section-nav">
    <ul class="td-sidebar-nav__section pe-md-3 ul-0">
          <li> .... </li>          
```

If you're using other Hugo themes other than Docsy, you can use the same trick to save and restore the scroll position.

{{% admonition type=note title="Note" %}}
I found the above code from a [demo website](https://hugo-book-demo.netlify.app/) of Hugo Book theme.
{{% /admonition %}}