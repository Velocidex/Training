<!-- .slide: class="title" -->

# Volatile machine state

---

<!-- .slide: class="content" -->

## Volatile machine state

* We have seen how to collect files from the endpoint
* However there is so much more than files:
  * Processes, Threads, Handles, Mutant
  * Process Memory can contain important information
  * Network connections

---

<!-- .slide: class="content" -->
## The order of volatility

* Volatile information changes rapidly
* Some things change more rapidly than others
* Acquiring information in the correct order helps to extract maximum
  value
  1. Call APIs like process listing, handles etc
  2. Scan process memory for signs of compromise
  3. Collect files like event logs, prefetch files etc.
  4. Maybe full memory capture - very rare!
