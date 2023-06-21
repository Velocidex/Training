<!-- .slide: class="title" -->

# Event queries for monitoring

## Using Event Queries for detection.

---

<!-- .slide: class="content" -->

## Example: Lateral Movement


* WMI may be used to [create processes remotely](https://www.blackhat.com/docs/us-15/materials/us-15-Graeber-Abusing-Windows-Management-Instrumentation-WMI-To-Build-A-Persistent%20Asynchronous-And-Fileless-Backdoor-wp.pdf):
```
wmic process call create "notepad.exe"
```

This works by invoking the `Create` method of the `Win32_Process` WMI
class.  This is very suspicious. Lets implement an Event Artifact to
detect this.

---

<!-- .slide: class="content" -->

## Lateral Movement: WMI Win32_Process.Create

* We can use WMI eventing to install a watcher on WMI calls
  themselves! The following WMI query will generate an event for each
  call of the `Create` method

```sql
SELECT * FROM
MSFT_WmiProvider_ExecMethodAsyncEvent_Pre
WHERE ObjectPath="Win32_Process" AND MethodName="Create"
```

* The full VQL query

```sql
SELECT Parse from wmi_events(
  query='''SELECT * FROM MSFT_WmiProvider_ExecMethodAsyncEvent_Pre
           WHERE ObjectPath="Win32_Process" AND MethodName="Create"''',
  namespace='ROOT/CIMV2', wait=50000000)
```

---

<!-- .slide: class="content" -->

## Exercise: Watch for new service creation

* Send the server an event for each new service installed into the system.
* Include the service binary authenticode signature information
* Test it with `winpmem.exe -l` and `winpmem.exe -u`

---

<!-- .slide: class="content" -->

## Sysmon support

* Sysmon is a great tool providing visibility into telemetry.

* In Velociraptor this is available through the artifact `Windows.Sysinternals.SysmonLogForward`

* It is a simple artifact that is used to just install sysmon and forward events.

<!-- .slide: class="content" -->

## Enable sysmon


---

<!-- .slide: class="title" -->

# Integration with external systems

## Interfacing with Elastic/Kibana

---

<!-- .slide: class="content" -->

## Connect to Elastic

* Velociraptor offers Elastic or Splunk bulk uploaders.
* Simple Bulk upload API clients - just push data upstream.
* Works by installing a server event monitor API:
    * Watch for System.Flow.Completion events
    * Get all the data for each artifact query and upload to elastic index

---

<!-- .slide: class="full_screen_diagram" -->


---

<!-- .slide: class="content" -->

## Installing ELK

* A lot of tutorials online about installing ELK
https://www.digitalocean.com/community/tutorials/how-to-install-elasticsearch-logstash-and-kibana-elastic-stack-on-ubuntu-18-04

* I will have elastic and kibana listening on localhost only, then use
  SSH to port forward (tunnel) to the ELK server.

* Forward both kibana and elastic ports.


---

<!-- .slide: class="content" -->
## Check for new indexes

* Velociraptor pushes data to indices named after the artifact name.

* By default Elastic automatically creates the index schema.
   * This is not idea because the index will be created based on the first row.
   * If the first row is different from the rest the index will be
     created with the wrong schema!
* This is why it is important to keep your artifacts consistent!
  always return the same rows even if null.

---

<!-- .slide: class="content" -->

## You can clear indices when you want

You can use the Elastic documentation to manipulate the Elastic server
https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-delete-index.html

---

<!-- .slide: class="content" -->

## Exercise: Create and maintain Elastic index

* We can maintain any kind of elastic index
* For this exercise, create and maintain an index of all known clients
    * Update this index every hour.

---

<!-- .slide: class="content" -->

## Integration with Slack/Discord

TODO

---

<!-- .slide: class="content" -->

## Exercise: Forwarding alerts to Discord

* Write an artifact to forward alerts to a discord channel.

---

<!-- .slide: class="content" -->

## Summary

* The hunting possibilities are only limited by your imagination!

* In this module we have adapted high level Tools Techniques and
  Procedures (TTPs) into Velociraptor artifacts. We can now hunt for
  them across the fleet.

* Two types of goals:
    * Establish a baseline
    * High value detection
