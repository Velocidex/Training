<!-- .slide: class="title" -->

# Extending VQL and the Velociraptor API

---

<!-- .slide: class="content small-font" -->

## Module overview

* VQL is really a glue language - we rely on VQL plugins and functions
  to do all the heavy lifting.
* To take full advantage of the power of VQL, we need to be able to
  easily extend its functionality.

* This module illustrates how VQL can be extended by including powershell scripts, external binaries and extending VQL in Golang.

* For the ultimate level of control and automation, the Velociraptor
API can be used to interface directly to the Velociraptor server
utilizing many supported languages (like Java, C++, C#, Python).
