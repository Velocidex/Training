<!-- .slide: class="title" -->
# Syslog logs

## Linux/Unix line based event logs

---

<!-- .slide: class="content" -->

## Line based logging

* On Linux line based logging is very common.

* Logs typically are **unstructurd**
    * Each application emits logs in free form text.
    * Makes it very difficult to accurately extract data

Example - Use Grok to detect SSH login events.

* Common compromise sequence:
    * Attackers compromise one machine through a vulnerability, or password guessing
    * Due to unsecured ssh keys, they can laterally move to other machines in the network.

---

<!-- .slide: class="content" -->

## Parsing SSH login events

* Linux systems typically use syslog for logging
   * Line based unstructured logs
   * Difficult to query across systems.
   * events are stored in /var/log/auth.log

* Looks similar to

<img src="ssh_log_sample.png" style="bottom: 0px"  />

---

<!-- .slide: class="content" -->

## Grok for parsing syslogs

* Grok is a way of applying regular expressions to extract structured information from log files.
* Used by many log forwarding platforms such as Elastic for example:

```
%{SYSLOGTIMESTAMP:Timestamp} %{SYSLOGHOST:logsource} \w+\[\d+\]: %{DATA:event}
%{DATA:method} for (invalid user )?%{DATA:user}
from %{IPORHOST:ip} port %{NUMBER:port} ssh2: %{GREEDYDATA:Key}
```

---

<!-- .slide: class="content small-font" -->

## Let's use VQL to parse ssh events

Read the first 50 lines from the auth log

```
Jun 25 18:56:08 devbox sshd[31872]: Accepted publickey for mic from 192.168.0.112 port 52323 ssh2: RSA SHA256:B4123453463443566345
```

![](parse_syslog_lines.png)

---

<!-- .slide: class="content" -->

## Filter lines and apply Grok


* Grok expressions for common applications are well published.
* You can figure out expressions for new log sources using online tools.

https://grokdebugger.com/

![](develop_grok.png)


---

<!-- .slide: class="content" -->

## Parsing log lines with Grok

* Applying the grok expression will match a dict
* Use foreach to expand the dict into columns

![](parse_syslog_grok.png)

---

<!-- .slide: class="content" -->

## Carving SSH auth logs

* SSH auth logs are often deleted from the system (either expired or
  maliciously)
* It is possible to carve for auth logs from the raw device.
* Use a fast but loose regular expression to box the syslog line
* Then apply the more accurate Grok parser to extract the line.
* Use the `raw_file` accessor in Linux to carve the raw disk device.

---

<!-- .slide: class="content" -->

## Exercise: Carving SSH auth logs

* Develop an artifact to carve SSH auth logs
* `Tip`: Create a very small sample for development by appending the
  read file to some junk data:

  ```
  cat /bin/ls auth.log > test.dd
  ```

* Apply the artifact on your Linux system to recover authentication
  events.
