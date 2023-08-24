---
title: Click Image and Zoom-in
---

## Medium Zoom

Home page: [Medium Zoom](https://medium-zoom.francoischalifour.com/)

### Docsy Example

Add the following code in the file `/layouts/partials/hooks/body-end.html`:

```html
<script src="https://cdnjs.cloudflare.com/ajax/libs/medium-zoom/1.0.7/medium-zoom.min.js" crossorigin="anonymous" referrerpolicy="no-referrer"></script>

<script>
const images = Array.from(document.querySelectorAll("img"));
images.forEach(img => {
  mediumZoom(img, {
    margin: 0, /* The space outside the zoomed image */
    scrollOffset: 40, /* The number of pixels to scroll to close the zoom */
    container: null, /* The viewport to render the zoom in */
    template: null /* The template element to display on zoom */
  });
});
</script>

<script > 
  /* */
  (function() {
    var a = document.querySelector("#td-section-nav");
    addEventListener("beforeunload", function(b) {
        localStorage.setItem("menu.scrollTop", a.scrollTop)
    }), a.scrollTop = localStorage.getItem("menu.scrollTop")
  })()
</script>
```

That's it. Now the images are automatically clickable and zoomable.

### Another Example

[An example for Hugo](https://github.com/russmckendrick/blog/blob/428e00b236fd3fcf484190d1d32759b9a51643b6/layouts/partials/extend_footer.html#L1-L13)

```html
<script src="https://cdnjs.cloudflare.com/ajax/libs/medium-zoom/1.0.6/medium-zoom.min.js" integrity="sha512-N9IJRoc3LaP3NDoiGkcPa4gG94kapGpaA5Zq9/Dr04uf5TbLFU5q0o8AbRhLKUUlp8QFS2u7S+Yti0U7QtuZvQ==" crossorigin="anonymous" referrerpolicy="no-referrer"></script>

<script>
const images = Array.from(document.querySelectorAll(".post-content img"));
images.forEach(img => {
  mediumZoom(img, {
    margin: 0, /* The space outside the zoomed image */
    scrollOffset: 40, /* The number of pixels to scroll to close the zoom */
    container: null, /* The viewport to render the zoom in */
    template: null /* The template element to display on zoom */
  });
});
</script>

<!-- https://ionic.io/ionicons -->

<script type="module" src="https://unpkg.com/ionicons@5.5.2/dist/ionicons/ionicons.esm.js"></script>
<script nomodule src="https://unpkg.com/ionicons@5.5.2/dist/ionicons/ionicons.js"></script>
```

## Other Choices

- [hugo-shortcode-gallery](https://github.com/mfg92/hugo-shortcode-gallery), see [demo](https://matze.rocks/images/)
