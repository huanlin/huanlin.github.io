---
title: Docsy - Click Image and Zoom-in
date: 2023-08-24
slug: docsy-image-zoom-in
tags: [hugo, docsy]
---

I use [Medium Zoom](https://medium-zoom.francoischalifour.com/) in my Docsy website to make images show full size when they are clicked.

## Demo

Click the following image to show it with full size:

![Taipei 101](images/taipei-101.jpg?width=600#center)

## Implementation

First, add the following code in `/layouts/partials/hooks/body-end.html`:

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
```

Then, add the following css to `/assets/scss/_styles_projects.scss`:

```css
/*
  For MediumZoom.js to force the zoomed image to be displayed on top of everything.
  Without these settings, the zoomed image will display under some elements on the page.

  Ref: https://github.com/francoischalifour/medium-zoom#debugging
*/
.medium-zoom-overlay,
.medium-zoom-image--opened {
  z-index: 999;
}
```

That's it. Now the images are automatically zoomed when they are clicked.

## Other Choices

- [hugo-shortcode-gallery](https://github.com/mfg92/hugo-shortcode-gallery), see [demo](https://matze.rocks/images/)
