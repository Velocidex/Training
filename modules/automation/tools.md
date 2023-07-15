<!-- .slide: class="title " -->

# Using external tools

<img src="tools.png" class="title-inset" />

---

<!-- .slide: class="content " -->

## Why use external tools?

* Velociraptor has a lot of built in functionality, but we can not
  cover all use cases!
* Velociraptor can automatically use external tools:
    * Velociraptor will ensure the tool is delivered to the endpoint
    * The tool can be called from within VQL
    * VQL can parse the output of the tool - thereby presenting the
      output in a structured way
    * VQL can then further process the data

---

<!-- .slide: class="content " -->

## Velociraptor Tools

* Tools are cached locally on the endpoint and are only re-downloaded
  when hashes change.
* Admin can control:
    * Where the tool is served from - Serve Locally or from Upstream
    * Admin can override the tool with their own
* Artifact writers can specify
    * The usual download location for the tool.
    * The name of the tool
* Velociraptor is essentially an orchestration agent

---

<!-- .slide: class="content " -->

## Autoruns
* Autoruns is a Sysinternals tool
* Searches for many different types of persistence mechanisms
* We could reimplement all its logic

OR

* We could just use the tool

---

<!-- .slide: class="content small-font" -->

## Autoruns

https://learn.microsoft.com/en-us/sysinternals/downloads/autoruns

![](autoruns_site.png)

---

<!-- .slide: class="content " -->

## Let's add a persistence

Using powershell run the following:
```
powershell> sc.exe create owned binpath="%COMSPEC% /c powershell.exe
   -nop -w hidden -command New-Item -ItemType File C:\art-marker.txt"
```

https://github.com/redcanaryco/atomic-red-team/blob/master/atomics/T1569.002/T1569.002.yaml

---

<!-- .slide: class="full_screen_diagram" -->

## Launching the Autoruns artifact

![](autoruns_artifact.png)

---

<!-- .slide: class="full_screen_diagram" -->

## Configuring the Autoruns tool

![](autoruns_tool_setup.png)


---

<!-- .slide: class="full_screen_diagram" -->

## Detect the malicious scheduled task

![](autoruns_detection.png)

---

<!-- .slide: class="content " -->

## Artifact encapsulation
* Use autoruns to specifically hunt for services that use %COMSPEC%
* Artifact writers just reuse other artifacts without needing to worry about tools.

```sql
name: Custom.Windows.Autoruns.Comspec
sources:
  - query: |
      SELECT * FROM Artifact.Windows.Sysinternals.Autoruns()
      WHERE Category =~ "Services" AND `Launch String` =~ "COMSPEC"
```
---

<!-- .slide: class="content " -->

## Exercise - Use Sysinternal DU

* Write an artifact to implement Sysinternal's `du64.exe` to calculate
  storage used by directories recursively.

---

<!-- .slide: class="content " -->

## Third party binaries summary

* `Generic.Utils.FetchBinary` on the client side delivers files to the client on demand.
* Automatically maintains a local cache of binaries.
* Declaring a new Tool is easy
* Admins can override tool behaviour
* Same artifact can be used online and offline
