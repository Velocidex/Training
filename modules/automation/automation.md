<!-- .slide: class="title " -->


# Automating the Velociraptor Server


---

<!-- .slide: class="content " -->

## Server artifacts

* Server automation is performed by exporting server administration
  functions as VQL plugins and functions
* This allows the server to be controlled and automated using VQL queries
* Server artifacts encapsulate VQL queries to performs certain actions
* Server monitoring artifacts watch for events on the server and respond.

---

<!-- .slide: class="content " -->

## Example: Client version distribution

* 30 Day active client count grouped by version

```
SELECT count() AS Count, agent_information.version AS Version
FROM clients()
WHERE timestamp(epoch=last_seen_at) > now() - 60 * 60 * 24 * 30
GROUP BY Version
```

---

<!-- .slide: class="content " -->
## Server concepts

* `Client`: A Velociraptor instance running on an endpoint. This is
  denoted by client_id and indexed in the client index.
* `Flow`: A single artifact collection instance. Can contain multiple
  artifacts with many sources each uploading multiple files.
* `Hunt`: A collection of flows from different clients. Hunt results
consist of the results from all the hunt's flows

---

<!-- .slide: class="content small-font" -->

## Exercise - label clients

* Label all windows machines with a certain local username.

1. Launch a hunt to gather all usernames from all endpoints
2. Write VQL to label all the clients with user "mike"

This can be used to label hosts based on any property of grouping that makes sense.
Now we can focus our hunts on only these machines.

---

<!-- .slide: class="content small-font" -->

## Exercise - label clients with event query

* The previous method requires frequent hunts to update the labels -
  what if a new machine is provisioned?

* Label all windows machines with a certain local username using an
event query.

---

<!-- .slide: class="content small-font" -->

## Exercise: Server management with VQL

* We use the offline collector frequently to facilitate collections on
  systems we have no access to.

* Write a server event query to automatically import new collections
  uploaded to:

  1. A Windows Share
  2. An S3 bucket.

---

<!-- content optional -->

## Exercise: Automating hunting

* Sometimes we want to run the same hunt periodically
* Automate scheduling a hunt collecting Scheduled Tasks every day at
  midnight.

---

<!-- .slide: class="content " -->

## Event Queries and Server Monitoring

* We have previously seen that event queries can monitor for new
  events in real time

* We can use this to monitor the server via the API using the
  `watch_monitoring()` VQL plugin.

* The Velociraptor API is asynchronous. When running event queries the
  `gRPC` call will block and stream results in real time.

---

<!-- .slide: class="content " -->
## Exercise - Watch for flow completions

* We can watch for any flow completion events via the API
* This allows our API program to respond whenever someone collects a
  certain artifact e.g.
     * Post process it and relay the results to another system).
     * Automatically collect another artifact after examining the
  collected data.

```sql
SELECT * FROM
    watch_monitoring(artifact=’System.Flow.Completion’)
```

---

<!-- .slide: class="content " -->

## Server Event Artifacts

* The Velociraptor server also offers a permanent Event Artifact
  service - this will run all event artifacts server side.

* We can use this to refine and post process events only using
  artifacts. We can also react on client events in the server.

---

<!-- .slide: class="content small-font" -->

## Exercise: Powershell encoded cmdline

* Powershell may accept a script on the command line which is base64
  encoded. This makes it harder to see what the script does, therefore
  many attackers launch powershell with this option
* We would like to keep a log on the server with the decoded
  powershell scripts.
* Our strategy will be:
   1. Watch the client’s process execution logs as an event stream on
      the server.
   2. Detect execution of powershell with encoded parameters
   3. Decode the parameter and report the decoded script.
   4. Use some regex to generate an escalation alert.

---

<!-- .slide: class="content " -->

## Exercise: Powershell encoded cmdline

* Generate an encoded powershell command using

```
powershell -encodedCommand ZABpAHIAIAAiAGMAOgBcAHAAcgBvAGcAcgBhAG0AIABmAGkAbABlAHMAIgAgAA==
```

Wait a few minutes for events to be delivered.

---

<!-- .slide: class="content small-font" -->

## Alerting and escalation.

* The `alert()` VQL function will generate an event on the
  `Server.Internal.Alerts` artifact.

* Alerts are collected from **all** clients or from the server.
* Alerts have a name and arbitrary key/value pairs.
* Alerts are deduplicated on the source.

* Your server can monitor that queue and issue an escalation to an
  external system:
    * Discord
    * Slack
    * Email

---

<!-- .slide: class="content " -->

## Exercise: Escalate alerts to slack/discord.

* Your instructor will share the API key for discord channel access.
* Write an artifact that forwards escalations to the discord channel.
