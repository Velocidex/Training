<!-- .slide: class="content " -->

## Review And Summary

* Velociraptor is essentially a collector of data
* Velociraptor has a number of ways to integrate and be controlled by
  other systems
* VQL provide `execve()` allowing Velociraptor to invoke other
  programs, and parse their output seamlessly.
* On the server VQL exposes server management functions allowing
  automating the server with artifacts.

---

<!-- .slide: class="content " -->

## Review And Summary

* The Velociraptor server exposes VQL via a streaming API - allowing
  external programs to Listen for events on the server
* Command the server to collect and respond
* Enrich and filter data on the server for better triaging and
  response.
