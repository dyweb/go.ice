# Design

[Previous design](design-old.md) focus on building API gateway, now we need to it to be a full web application framework
that just work, and don't focus on security and reliability (not that gateway).

Mainly follow https://github.com/gobuffalo/buffalo but with less features on frontend, auth etc.

- cli for generating code sketch or specific functions (json & protobuf)
  - [ ] requires gommon/runner to update

## directory layout


## FAQ

- Why not using existing framework?
  - not fun (just can't stop re inventing the wheel, never learn)
  - not using feature of go in newer version, i.e. gorilla toolkit