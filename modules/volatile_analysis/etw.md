<!-- .slide: class="title" -->

# Event Tracing for Windows

---

<!-- .slide: class="content small-font" -->

## What is ETW

* ETW is the underlying system by which event logs are generated and collected.
https://docs.microsoft.com/en-us/windows-hardware/test/weg/instrumenting-your-code-with-etw

<img src="etw_arch.png" style="width: 50%"  />

---

<!-- .slide: class="content" -->

## ETW Providers

Show all registered ETW providers

```
logman query providers
```


Show details about each provider

```
logman query providers Microsoft-Windows-DNS-Client
```

---

<!-- .slide: class="content small-font" -->

## ETW for event driven logs

* ETW and event logs are just two sides of the same coin

<div class="container">
<div class="col">

* Log providers are just ETW providers
   * In VQL `watch_etw()` can be used instead of `watch_evtx()

* See `Windows.Sysinternals.SysmonLogForward` for an example

</div>
<div class="col">

![](etw_event_log.png)

</div>

---

<!-- .slide: class="content" -->

## Exercise - Monitor DNS queries

* Use ETW to monitor all clients' DNS queries.

* Stream queries to server

---

<!-- .slide: class="full_screen_diagram" -->

## Exercise - Monitor DNS queries

![](etw_follow_dns.png)

---


<!-- .slide: class="content" -->

## Windows Management Instrumentation

* A framework to export internal windows state information using a query language (WQL)
* Consists of classes (providers) and objects
* Lots of hooks into many internal system features
* Being able to inspect system state using a consistent interface allows a tool to query a wide range of services.


---

<!-- .slide: class="full_screen_diagram" -->

## WMI Explorer

![](wmie_github.png)

---

<!-- .slide: class="full_screen_diagram" -->

## WMI Explorer

![](wmie.png)

---

<!-- .slide: class="full_screen_diagram" -->

```sql
SELECT * FROM wmi(query="SELECT * FROM win32_diskdrive")
```

![](wmi_diskdrive.png)
