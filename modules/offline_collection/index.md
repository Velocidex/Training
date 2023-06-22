<!-- .slide: class="title" -->

# Interactive triage collections

## Preserving and collecting data

---

<!-- .slide: class="content small-font" -->

## Module overview

* We previously saw how Velociraptor can be used to triage, collect
  indicators and remotely analyze a system.
* Sometimes we can not deploy the Velociraptor client/server model but
  we still want to be able to collect artifacts.
* Sometimes we need to rely on another agent to actually do the
  collection (either a human agent or another software).
* This module shows how to prepare Velociraptor for local, interactive
  triage collection - simply collect the relevant artifacts and ship
  them out.

---

<!-- .slide: class="content small-font" -->

## Velociraptor Artifacts

* Velociraptor is essentially a VQL engine.
* VQL queries are encapsulated inside an Artifact - a structured yaml
  file describing the query’s purpose and parameters.
* We have seen how to collect artifacts remotely using the
  client/server model - however we can also collect artifacts locally

---

<!-- .slide: class="content " -->

## Collecting files

* Being able to efficiently and quickly collect and preserve evidence
  is important:

    * Capture machine state at a point in time.
    * Collect files for further analysis by forensic tools.

---

<!-- .slide: class="content " -->

## Windows.KapeFiles.Targets

* This is a popular artifact for mass file collection.
* It does no analysis but just collects a bunch of files.
* Uses low level NTFS accessor

---

<!-- .slide: class="content " -->

## Windows.KapeFiles.Targets
* Simply select the target to collect.
* Many targets automatically include sub-targets.

---

<!-- .slide: class="content " -->

## Resource control
* Collecting artifacts can generate huge amount of data
* Because Velociraptor is so fast and efficient it is easy to
  accidentally overwhelm networks

* Math is a harsh mistress:
   * Collecting 100Mb  from 10,000 endpoints = 1Tb
   * e.g. $MFT is usually around 300-400Mb

---

<!-- .slide: class="full_screen_diagram" -->

### Velociraptor has your back!


---

<!-- .slide: class="title" -->

# Offline collections

## Digging deeper without a server

---

<!-- .slide: class="content " -->

## Why Offline collection?

* I want to collect artifacts from an endpoint
* But Velociraptor is not installed on the endpoint!
* Or the endpoint is inaccessible to the Velociraptor server (no
  egress, firewalls etc).

But Velociraptor is just a VQL engine!  It does not really need a
server anyway

---

<!-- .slide: class="content " -->

## Create an offline collector


* Creating an offline collector looks very similar to collecting
  client artifacts

* Only difference is that results are delivered over sneakernet!


---

<!-- .slide: class="content " -->

## Prepare a special executable

---

<!-- .slide: class="full_screen_diagram" -->

### Collector binary automatically starts collection as soon as it is run… No need for user to enter command line parameters.


---

<!-- .slide: class="content " -->

## Offline collection

* Collector creates a container with all the files and query results

---

<!-- .slide: class="content " -->

## Exercise: Collect triage data and upload to a cloud bucket

Configure an offline collector for cloud upload.

---

<!-- .slide: class="content " -->

## Protecting the collection file: Encrypion
* For added protection, add a password to the zip file
* If we used a simple password it would be embedded in the collector
* Use an X509 scheme to generate a random password.

* Zip files do not password protect the directory listing - So
  Velociraptor creates a second zip file inside the password protected
  zip.

---

<!-- .slide: class="content " -->

## Including third party binaries

* Sometimes we want to collect the output from other third party
  executables.
* Velociraptor can package dependent binaries together with the
  offline collector.
* Velociraptor can append a zip file to the end of the binary and
  adjust PE headers to ensure it can be properly signed.

---

<!-- .slide: class="content " -->

## Take a memory image with winpmem

* We will shell out to winpmem to acquire the image. We will bring
  winpmem embedded in the collector binary.

---

<!-- .slide: class="content " -->
## Preparing an SMB dropbox

* Sometimes it is easiest to configure an SMB directory to receive the
  offline collector.

* TODO


---

<!-- .slide: class="content " -->

## Importing collections into the GUI

---

<!-- .slide: class="content " -->

## Local collection considerations
* Local collection can be done well without a server and permanent
  agent installed.
* A disadvantage is that we do not get feedback of how the collection
  is going and how many resources are consumed.
* We really need to plan ahead what we want to collect and it is more
  difficult to pivot and dig deeper in response to findings.
