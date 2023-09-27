## Scoping the environment

What users are logging into machines?

`Windows.Sys.AllUsers`

```vql
SELECT Name, UUID, Mtime, count()
FROM source(artifact="Windows.Sys.AllUsers")
WHERE Name
GROUP BY Name
```

May or may not be suspicious.

## Are any of the users recently created?

Use the VFS to navigate to the user's home directory - note the birth
time.

User home directory creation time is a good proxy for when the user
first logged in.

Hunt for `FileFinder` - post process by sorting the birth time of the
home directory.

Let's talk about file finder as a general purpose tool for fetching file
metadata and data.


## Let's look at local users created in the SAM

Get users who logged in recently.

```vql
SELECT ParsedF.LastLoginDate AS LastLoginDate, ParsedV, ClientId, Fqdn
FROM source(artifact="Windows.Forensics.SAM")
WHERE LastLoginDate > "2023-01-01"
```

The `winsupport` user seems suspicious... No one knows about it....

## RDP Auth

Collect RDP authentications from the event logs `Windows.EventLogs.RDPAuth`

```vql
SELECT EventTime, Computer, SourceIP, UserName, Description, ClientId , count() AS Count
FROM source(artifact="Windows.EventLogs.RDPAuth")
WHERE Description =~ "LOGON_SUCCESSFUL"
GROUP BY UserName, Description, ClientId
```

Get timeline of login - what is the blast radius?

Which machines are affected?

Get earliest use of `winsupport`:

```vql
SELECT * FROM source(artifact="Windows.EventLogs.RDPAuth")
WHERE UserName =~ "winsupport" and Description =~ "SUCCESS"
ORDER BY EventTime
```

## How is RDP allowed?

Check local firewall rules `Windows.Sys.FirewallRules` for RDP access?

## Lets look at created services

Use the artifact `Windows.EventLogs.ServiceCreationComspec` to search for created services - update the service regex to `.`

## Check other methods of logging in

Look for all login sessions `Exchange.Windows.EventLogs.LogonSessions`

See this for logon types:

https://www.ultimatewindowssecurity.com/securitylog/encyclopedia/event.aspx?eventid=4624

Type 3: *Network (i.e. connection to shared folder on this computer from elsewhere on network)*

## What happened on the machine around that time?
### Hayabusa + Sigma

Hayabusa is a "SIEM in a box" - tool for running many Sigma rules over
the event logs on the end point.

A lot of false positives so it is useful for a quick overview before
digging deeper.

```vql
SELECT *, count()
FROM source(artifact="Exchange.Windows.EventLogs.Hayabusa/Results")
GROUP BY RuleTitle
```

Order by level to show critical first.

Lots of interesting activities!

- Look for `winsupport` login events
- Account creation alerts
- Service creation - `psexec`

## Search for ADS

Mark of the web can sometimes give us a hint of where a file came from `Windows.NTFS.ADSHunter`

In this demo we use C:\Users\ to limit the time taken.

## What files appeared on the endpoint? USN Journal

The USN Journal records file activity on the endpoint.

Limit by the earliest time

- Look for interaction with powershell files - see new powershell file created
- Look for `psexec` files...
- Look for `prefetch` file

- Look for executable files being created - find `notsuspicious.exe`
  created in Windows directory - very suspicious!
- search for file with a .key extension - typical tool mark of
  `psexec`. This also tells us where the attacker came from.

The USN Journal allows us to look back in time
```vql
SELECT * FROM source(artifact="Windows.Forensics.Usn")
WHERE OSPath =~ "\\.exe$" AND Reason =~ "DELETE"
```

What executables were deleted? In the windows directory?

## Look for powershell artifacts:

- `ISEAutosave`
- `Powershell ReadLine`

Examine the powershell activity - disabling firewall

## Process execution

Prefetch timeline - see activity in prefetch

## SQLite Hunting

We still don't know exactly what the `winsupport` user did?

`SQLiteHunter` parses many artifacts
- browser artifacts - History downloads etc. Reveal the watering hole.

## What else did the attacker do on the system?

Lets find evidence of the attacker interacting with the system.

- `RecentDocs`
- `Lnk analysis`

This confirms the attacker opened the documents to view them and
potentially ex-filtrated them.

## Look for new services T1543.003

`Windows.System.Services`

Closely inspect unsigned services.

```vql
SELECT Name, PathName, HashServiceExe, CertinfoServiceExe
FROM source(artifact="Windows.System.Services")
WHERE NOT CertinfoServiceExe.Trusted
```

Services with low frequency

```vql
SELECT Name, PathName, HashServiceExe, CertinfoServiceExe, count() AS Count
FROM source(artifact="Windows.System.Services")
GROUP BY HashServiceExe
```

## Scheduled tasks

```vql
SELECT *, count() AS Count
FROM source(artifact="Windows.System.TaskScheduler/Analysis")
GROUP BY Command
```
