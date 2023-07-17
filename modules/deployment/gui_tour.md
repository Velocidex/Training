<!-- .slide: class="title" -->
# A Velociraptor GUI tour

<img src="/modules/gui_tour/tour-bus.png" class="title-inset">

---

<!-- .slide: class="content" -->

## The Dashboard

* The Dashboard shows the current state of the installation:
    * How many clients are connected
    * Current CPU load and memory footprint on the server.
    * When running hunts or intensive processing, memory and CPU requirements will increase but not too much.
    * You can customize the dashboard - it’s also just an artifact.

---

<!-- .slide: class="full_screen_diagram" -->

## The Dashboard

<div style="text-align: center;">
    <img src="/modules/gui_tour/dashboard.png" style="height: 500px">
</div>

---

<!-- .slide: class="content" -->
## User Preferences

You can customize the interface to your liking

<div class="container small-font">
<div class="col">

* Themes

</div>
<div class="col">

* Languages
* Timezones

</div>
</div>

<img src="/modules/gui_tour/user_preferences.png" style="height: 300px" class="" />

---

<!-- .slide: class="title" -->

# Interactively investigate individual clients

---


<!-- .slide: class="content small-font" -->
## Searching for a client

To work with a specific client we need to search for it.  Press the
**Search** or **Show All** button to see some clients. You can also
use the **Show recent hosts** to see your own clients.

<div style="text-align: center;">
    <img src="/modules/gui_tour/search_clients.png" class="mid-height"
        style="height: 300px">
</div>

---

<!-- .slide: class="content" -->
## Search for clients

### hostname, label, or client ID.

* You can start typing the hostname to auto-complete
* Some common terms:
   * `host`: search by hostnames
   * `mac`: Mac addresses
   * `ip`: last seen IP address
   * `label`: Search by labels
   * `recent`: Show all clients recently interacted with

---


<!-- .slide: class="content small-font" -->
## Client Overview

* Internally the client id is considered the most accurate source of
endpoint identity

<div style="text-align: center;">
    <img src="/modules/gui_tour/client_overview.png" class="mid-height"
        style="height: 400px">
</div>


---

<!-- .slide: class="content small-font" -->
## Shell commands

* Velociraptor allows running shell commands on the endpoint using
  `Powershell`/`Cmd`/`Bash`
    * Only Velociraptor users with the administrator role are allowed to
  do this!
    * Actions are logged and audited

```powershell
Get-LocalGroupMember -Group "Administrators"
```

---

<!-- .slide: class="content" -->

## Running Shell Commands

![](/modules/gui_tour/shell_commands.png)

---

<!-- .slide: class="title" -->
# Interactively fetching files from the endpoint

<img src="/modules/gui_tour/fetch.png" class="title-inset">

---

<!-- .slide: class="full_screen_diagram small-font" -->
### The VFS View

Remember that the VFS view is simply a server side cache of
information we know about the endpoint - it is usually out of date!

![](/modules/gui_tour/vfs_view.png)

---


<!-- .slide: class="content small-font" -->
## Navigating the interface

* Click the “Refresh this directory” will schedule a directory listing
  artifact and wait for the results (usually very quick if the
  endpoint is online).
* The “Recursively refresh this directory” will schedule a recursive
  refresh - this may take some time! After this operation a lot of the
  VFS will be pre-populated already.
* “Collect from client” will retrieve the file data to the
  server. After which, the floppy disk sign indicates that we have
  file data available and you can click the “Download” link to get a
  copy of the file.


---

<!-- .slide: class="content small-font" -->
## The VFS interface

Previewing a file after download.

![](/modules/gui_tour/vfs_view_2.png)

---

<!-- .slide: class="content small-font" -->

## Previewing files

The GUI allows close inspection of binary files
* Viewing in hex or text
* Paging - skipping to offset
* Searching using regex or hex strings

<img src="/modules/gui_tour/vfs_view_3.png" class="title-inset">

---

<!-- .slide: class="content" -->
## Velociraptor artifacts

Velociraptor is just a VQL engine!

* We package VQL queries in Artifacts:
    * YAML files
    * Include human description
    * Package related VQL queries into “Sources”
    * Take parameters for customization
    * Can in turn be used in VQL as well...

---

<!-- .slide: class="full_screen_diagram small-font" -->

### What does the VFS view do under the cover?

* Refreshing the VFS simply schedules new artifacts to be collected - it is just a GUI convenience.

![](/modules/artifacts_introduction/vfs_collections.png)

---

<!-- .slide: class="content" -->

## Velociraptor uses expert knowledge to find the evidence

A key objective of Velociraptor is encapsulating DFIR knowledge into
the platform, so you don’t need to be a DFIR expert.  We have high
level questions to answer We know where to look for evidence of user /
system activities

We build artifacts to collect and analyze the evidence in order to answer our investigative questions.
