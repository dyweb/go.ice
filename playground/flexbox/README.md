# Flexbox

Before flexbox

- inline element display from left to right, i.e. `<span>`
- block display from top to down, i.e. `<div>` and `<p>`

Sticky footer

- the key is `flex-grow` set to 1 for the main content

Vertical Centering

- [file](center-box.html) 
- [blog](https://philipwalton.github.io/solved-by-flexbox/demos/vertical-centering/)
- use `justify-content` and `align-items` on different axis

## Ref

- https://philipwalton.github.io/solved-by-flexbox/ real example
- https://css-tricks.com/snippets/css/a-guide-to-flexbox/ detail explanation
- https://www.bitovi.com/blog/use-flexbox-to-create-a-sticky-header-and-sidebar-with-flexible-content
  - `flex-grow: 1;  /*ensures that the container will take up the full height of the parent container*/`