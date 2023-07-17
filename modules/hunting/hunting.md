<!-- .slide: class="content" -->

## Typical hunting workflow

1. Get an idea for a new artifact by reading blogs, articles and doing research!
2. Explore the VQL in the notebook
3. Convert VQL into an artifact
4. Go hunting!
5. Back in the notebook - post process and analyze

---

<!-- .slide: class="content" -->

## Mitre Att&ck framework

* Mitre framework is the industry standard in documenting and
  identifying attacker Tactics, Techniques and Procedures (TTP)

---

<!-- .slide: class="content" -->

## Atomic Red Team

* We will use Atomic Red Team to help develop our artifacts!

---

<!-- .slide: class="content" -->
## Exercise: Detect Att&ck Techniques

https://attack.mitre.org/techniques/T1183/

---

<!-- .slide: class="content" -->
## First plant a signal on your machine

```
REG ADD "HKLM\SOFTWARE\Microsoft\Windows NT\CurrentVersion\Image File Execution Options\notepad.exe" /v Debugger /t REG_SZ /d """"C:\Program Files\Notepad++\notepad++.exe""" -notepadStyleCmdline -z" /f
```

* Type notepad - you get notepad++ (useful but….)

---

<!-- .slide: class="content" -->
## Windows.Persistence.Debug

* Write an artifact to detect this modification.
* Hash and upload a copy of each binary specified by this key.
* Create a whitelist of OK binaries to have here - Use digital
  signatures to verify the authenticity of the binary.

* The exercise demonstrate extending the basic detection capability
  with enrichment and customization

---

<!-- .slide: class="content" -->
## Hunting - mass collections

Hunting is Velociraptor's strength - collect the same artifact from
thousands of endpoints in minutes!

* Two types of hunts:
   * Detection hunts are very targeted aimed at yes/no answer
   * Collection hunts collect a lot more data and can be used to
     build a baseline.

---

<!-- .slide: class="content" -->
## Exercise - baseline event logs

For this exercise we start a few more clients.

```text
c:\Users\test>cd c:\Users\test\AppData\Local\Temp\

c:\Users\test\AppData\Local\Temp>Velociraptor.exe
   --config client.config.yaml pool_client --number 100
```

This starts 100 virtual clients so we can hunt them
* We use pool clients to simulate load on the server

---


<!-- .slide: class="full_screen_diagram" -->
## Pool clients
Simply multiple instances of the same client

![](/modules/bit_log_disable_hunting/pool_clients.png)

---


<!-- .slide: class="full_screen_diagram" -->
## Create a hunt

![](/modules/bit_log_disable_hunting/create-hunt_2.png)

---


<!-- .slide: class="full_screen_diagram" -->
## Select hunt artifacts

![](/modules/bit_log_disable_hunting/create-hunt_3.png)
---


<!-- .slide: class="full_screen_diagram" -->
## Collect results

![](/modules/bit_log_disable_hunting/create-hunt.png)

---


<!-- .slide: class="content" -->
## Exercise - Stacking

* The previous collection may be considered the baseline
* For this exercise we want to create a few different clients.
    * Stop the pool client
    * Disable a log channel
    * Start the pool client with an additional number of clients

```
Velociraptor.exe --config client.config.yaml pool_client --number 110
```

---


<!-- .slide: class="full_screen_diagram" -->
## Stacking can reveal results that stand out

![](/modules/bit_log_disable_hunting/stacking-a-hunt.png)

---

<!-- .slide: class="content" -->

## Exercise: Labeling suspicious hits

* After stacking it becomes obvious which machines are out of place.
* We can label those machines in order to narrow further hunting to them.
* Use the `label()` function to add a label to all machines with the
  disabled log sources.

---

<!-- .slide: class="content" -->

## Exercise: Post processing a large hunt

* For this exercise collect all scheduled tasks from all systems.


---

<!-- .slide: class="content" -->

## Create a malicious service
* Lets create a malicious service
```
sc.exe create backdoor binpath="c:\Windows\Notepad.exe"
```

* And start a couple more clients

```
velociraptor.exe --config client.config.yaml pool_client --number 102 --writeback_dir filestore
```

---

<!-- .slide: class="content" -->

## Optimizing filtering

* By default VQL run each query on one core, examining a row at a
  time.
* You can speed up filtering by using the `parallelize()` plugin
    * Same parameters as the `source()` plugin with the addition of a
      query.
    * The specified query will run on multiple workers and receive
      rows from the `source()` plugin.
    * Faster than `foreach(workers=30)` because the result set parsing
      is also parallelized.

---

<!-- .slide: class="content" -->

## Recollecting failed hunts

* Sometimes a collection may have failed (e.g. timeout exceeded)
* We might want to redo the same collection in that hunt:
   * Find the failed collection
   * Press the "Copy Collection" button in the toolbar
   * Modify the collection parameters (e.g. timeout)
   * Relaunch the new collection
   * When satisfied, simply add the new collection to the hunt
     manually.

---


<!-- .slide: class="content" -->

## Exercise: Automating hunting

* Sometimes we want to run the same hunt periodically
* Automate scheduling a hunt collecting Scheduled Tasks every day at
  midnight.
