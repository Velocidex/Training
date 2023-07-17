<!-- .slide: class="content" -->

## Review And Summary

* Event logs are very important for Incident Response
* The Windows Event logs format is structured, but does not contain
  critical information.
* Using automated tools to apply detection logic on event logs helps
  to quickly identify an incident.
* Hunting the EVTX logs can reveal important but unrelated
  information.

---

<!-- .slide: class="content" -->

## Review And Summary

* On Linux syslog events are not structured.
* We can use `Grok` or regular expressions to extract information but
  this is fragile.
* Syslog can be carved from disk - even after rotation.
* Newer systems are using journal logs as well.
