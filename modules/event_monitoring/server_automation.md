<!-- .slide: class="title" -->

# Server Automation With Event Queries

---

<!-- .slide: class="content" -->

## Server Automation With Event Queries

* It is possible to automate things based on server side events.
* When collections complete they emit `System.Flow.Completion` which
  can be watched using the `watch_monitoring()` plugin.

---

<!-- .slide: class="content" -->

## Exercise: Automatically archive Windows.KapeFiles.Targets

* When the user collects a `Windows.KapeFiles.Targets` artifact,
  zip the collection up (with [create_flow_download](https://docs.velociraptor.app/vql_reference/server/create_flow_download/)) and archive it into a directory

---

<!-- .slide: class="content" -->
## Exercise: Import offline collections

* Write an artifact that will automatically import any new offline
  collection thatt appear in a directory.
* Also add the collections to a running hunt.
