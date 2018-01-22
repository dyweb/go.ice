# Design

[Previous design](design-old.md) focus on building API gateway, now we need to it to be a full web application framework
that just work, and don't focus on security and reliability (not that gateway).

Mainly follow https://github.com/gobuffalo/buffalo but with less features on frontend, auth etc.

- cli for generating code sketch or specific functions (json & protobuf)
  - [ ] requires gommon/runner to update

## directory layout

### ice

- app application struct with common structs, avoid scatter config, dbmgr around in main.go,
  - client
  - server
- cache kv cache, remote & in process cache (might do something like group cache? https://github.com/golang/groupcache)
- config configuration structs to avoid cycle import
- db relational database
  - drivers wrapper around existing sql drivers
  - migration 
  - cmd.go handy db command to be used in application
  - manager.go a singleton for each application
- [ ] interface and util for server client, might look at swagger?
- util
  - logutil logger registry for the library
  
### application using ice

- db
  - migration migration tasks
- server server implementation
- client client wrapper
- service interface and common error code etc.

## FAQ

- Why not using existing framework?
  - not fun (just can't stop re inventing the wheel, never learn)
  - not using feature of go in newer version, i.e. gorilla toolkit