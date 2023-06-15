<!-- .slide: class="title" -->

# Interactive triage collections

## Digging deeper without a server

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

5
Collecting Artifacts
Capturing endpoint state

Collecting files
Being able to efficiently and quickly collect and preserve evidence is important:

Capture machine state at a point in time.
Collect files for further analysis by forensic tools.
6

Windows.KapeFiles.Targets
This is the most popular artifact for mass file collection.
It does no analysis but just collects a bunch of files.
Uses low level NTFS accessor
7

Windows.KapeFiles.Targets
Simply select the target to collect.
Many targets automatically include sub-targets.
8

9

Resource control
Collecting artifacts can generate huge amount of data
Because Velociraptor is so fast and efficient it is easy to accidentally overwhelm networks

Math is a harsh mistress:

Collecting 100Mb  from 10,000 endpoints = 1Tb
(e.g. $MFT is usually around 300-400Mb)
10

Resource control
Velociraptor has your back!

Every artifact collection is limited automatically:
Op/Sec - how fast to run on the endpoint
Max Time - Cancel if collection takes too long
Max Rows - Cancel if query returns crazy rows
Max Mbytes - Cancel if too much data is collected
11

Limit the collection size
12

13

14
Offline collections

Why Offline collection?
I want to collect artifacts from an endpoint

But Velociraptor is not installed on the endpoint!

Or the endpoint is inaccessible to the Velociraptor server (no egress, firewalls etc).

But Velociraptor is just a VQL engine!  It does not really need a server anyway
15

Create an offline collector
16

Creating an offline collector looks very similar to collecting client artifacts

Only difference is that results are delivered over sneakernet!
17

18

Prepare a special executable
19

20
Collector binary automatically starts collection as soon as it is run… No need for user to enter command line parameters.

21
Collector creates a container with all the files and query results

Also generates a html report as a summary of what was collected

22
The report is generated from the artifacts collected.

It is possible for you to create your own HTML report in your own custom artifact!

Local artifact collection
23
The “artifact collect” prints the output of the artifact to the console
Artifact Name and optionally a source name

Collecting to a zip file
24
Using the output flag we can redirect all output to a zip file. Artifact result sets are written as CSV files
This also redirects the upload() plugin into the output zip file - so we can capture files.

Collecting files
25
Sometimes we need to collect files for triage. Use the Windows.KapeFiles.Targets artifact to collect bulk data. This artifact uses preset sets of glob expressions to collect forensically relevant files using low level NTFS parsing (to avoid locked files).

Embedding configurations
26
Normally when we run Velociraptor we need to provide a configuration file.
Sometimes it is easier to embed the configuration file inside the binary - this way the user does not need to remember another file.
When Velociraptor starts up without the --config flag supplied - it will check its own binary for embedded configuration.

Autorun configuration
Velociraptor has many command line options and flags
When we need a human agent to help, this can be confusing - many users find the command line difficult and there can be a lot of options to type.

By embedding an autoexec command in the config, the binary will begin executing the command line when run with no options
27

Exercise
Create a self-executing package which collects event log files into a zip file.

Collect a process listing and dump out the wmi host service memory as well.
28

Whats going on under the covers?
29

30

31
Scenario: Collect triage data and upload to a cloud bucket

Artifact collection VQL
Artifact collection is exposed via VQL
You can design your collection requirements in advance using VQL artifacts
32

Uploading files to the cloud
Velociraptor can push files to a cloud bucket. We will use AWS for this example, but you can also use Google Cloud Storage.
Credentials for writing to the bucket will be embedded in the binary
If the binary is compromised we do not want to allow the attacker to download or delete any evidence! Therefore set up careful bucket permissions.
33

34

35

36

37
Creating an AWS user
An AWS User is identified by a key pair - this secret will be embedded in the binary.
A user belongs in a group
A group has a policy applied to it.
A policy is a list of permissions that are allowed.

In our case we need to only write to the bucket (not even read or delete).

38
Create a policy which applies to all objects inside this bucket

39

40
Create a new group which has the policy attached

41
Now add a user into this group

42

43

44

45

46
How to upload to S3?

47

Protecting the collection file
For added protection, add a password to the zip file
48

49

50
Zip files do not password protect the directory listing - So Velociraptor creates a second zip file inside the password protected zip.

51

52
Include third party binaries
Sometimes we want to collect the output from other third party executables.

It would be nice to be able to package them together with Velociraptor and include their output in the collection file.

53
Velociraptor can append a zip file to the end of the binary and adjust PE headers to ensure it can be properly signed.

You can just add any file to the zip file and access it using VQL from within Velociraptor Artifacts

54

55
Take a memory image with winpmem
We will shell out to winpmem to acquire the image. We will bring winpmem embedded in the collector binary.
So the strategy is
Locate Velociraptor’s own binary path
Open the file using the “zip” accessor to recover the payload
Copy winpmem from the archive into a temporary file.
Shell out to winpmem to produce a memory image.
When the image is complete, upload the image to S3

56

57

58
Collecting evidence across the network.

59
Collecting across the network
By having a single executable all we need is to run it remotely. We can use another EDR solution that allows remote execution if available.
We can use windows own remote management to deploy our binary across the network.
Copy our collector binary across the network to C$ share
Use wmic to launch our binary on the remote host.

60
Be aware that passing creds to a compromised host might facilitate a pass the hash attack

61

62
Importing third party artifacts
Velociraptor is an open source project

Many contributors may share their artifacts.

Recently launched CCXLabs as an example of an external artifact provider.

63
Artifacts can be imported in bulk from external providers.

Exercise - import CCXLabs
Download the CCXLabs project as a zip file.

Simply import CCXLabs repository dump into your Velociraptor
64

65

Create a CCX Digger collector
66

67

CCX Digger makes a custom report
68

Importing collections into the GUI
It is possible to import the offline collection back into the GUI
This allows:
Keeping related information from the same host together
Using a notebook to post process the results

Offline collection + Import is very similar to client/server except that instead of the client connecting over the internet, the data is delivered via sneakernet!
69

Exercise:
Collect MFT using offline collection
Import into the GUI
Post process - filter only the files modified in the last week
Create a timeline
Can you retrace your steps?
70

Importing collections into the GUI
71

72

73

74

75
Local collection considerations
Local collection can be done well without a server and permanent agent installed.
One way can be done via windows remote procedure call mechanism

But this is limited to machines with network browsing accessible (e.g. on LAN)


76
Conclusions
A disadvantage is that we do not get feedback of how the collection is going and how many resources are consumed.

We really need to plan ahead what we want to collect and it is more difficult to pivot and dig deeper in response to findings.
