
<!-- .slide: class="content" -->
## NTFS Overview

* NTFS is the file system in all modern Windows operating systems.
* Feature packed with a design focused on storage optimization and resilience.
* NTFS implements Journalling to record metadata changes to track state and integrity of the file system.
* Allows for recovery after system crashes to avoid data loss
* File System objects referenced in a Master File Table (MFT)

---


<!-- .slide: class="content" -->
## New Technology File System

* In NTFS, the Master File Table (MFT) is at the heart of the file
  system. A structured database that stores metadata entries for every
  file and folder.
* Every object gets an entry within the MFT. Each entry is usually
  1024 bytes long.  Contains a series of attributes that fully
  describe the object.

---

<!-- .slide: class="content" -->

## MFT entries contain attributes

<div class="container small-font">
<div class="col">

## File entry examples
* $STANDARD_INFORMATION
* $FILE_NAME (Windows long name)
* $FILE_NAME (short name)
* $DATA
* $DATA  (alternate data stream sometimes)

</div>
<div class="col">

## Folder entry examples
* $STANDARD_INFORMATION
* $FILE_NAME (Windows long name)
* $FILE_NAME (short name)
* $INDEX_ROOT
* $INDEX_ALLOCATION (sometimes)

</div>

---

<!-- .slide: class="content small-font" -->

## NTFS Analysis

Velociraptor offers a number of plugins to access detailed information
about NTFS:
* `parse_mft()`: parses each MFT entry and returns high level metadata
  about the entry - including reconstruct the full path of the entry
  by traversing parent MFT entries.
* `parse_ntfs()`: Given an MFT ID this function will display
  information about the various streams (e.g. `$DATA`, `$Filename`
  etc)
* `parse_ntfs_i30()`: This scans the `$i30` stream in directories to
  recover potentially deleted entries.

---

<!-- .slide: class="content small-font" -->
## Finding suspicious files

Parse the MFT using `Windows.NTFS.MFT`

* Common DFIR use case is finding files
    * File name
    * Path
    * File type
    * Content
* Velociraptor plugins
    * glob
    * parse_mft
    * yara
    * other content based plugins

<img src="/modules/ntfs_forensics/MFT_artifact.png" style="bottom: 0px" class="inset" />

---


<!-- .slide: class="content" -->
## Windows.Forensics. FilenameSearch

* Apply yara on the MFT
    * fast yara
    * simple string based
    * filename / top level folder only
    * comma separated
* Crude and less control
* Verbose results

<img src="/modules/ntfs_forensics/Windows.Forensics.FilenameSearch.png" style="bottom: 0px" class="inset" />

---

<!-- .slide: class="content" -->
## Windows.NTFS.MFT

<div class="container small-font">
<div class="col">

* Parses MFT
* Easy to use
* Filters
    * Path
    * File name
    * Drive
    * Time bounds
    * Size
* Performance optimised

</div>
<div class="col">

<img src="/modules/ntfs_forensics/Windows.NTFS.MFT.png" style="bottom: inherit" class="inset" />

</div>
</div>

---


<!-- .slide: class="content" -->
## Exercise - Generate test data

To automatically prep your machine run this script:
```powershell
### NTFS exercise setup

## 1. download some files to test various content and add ADS to simulate manual download from a browser

$downloads = (
    "https://live.sysinternals.com/PsExec64.exe",
    "https://live.sysinternals.com/procdump64.exe",
    "https://live.sysinternals.com/sdelete64.exe"
)

foreach ( $url in $downloads){
    "Downloading " + $Url
    $file = Split-Path $Url -Leaf
    $dest = "C:\PerfLogs\" +$file
    $ads =  "[ZoneTransfer]`r`nZoneId=3`r`nReferrerUrl=https://18.220.58.123/yolo/`r`nHostUrl=https://18.220.58.123/yolo/" + $file + "`r`n"

    Remove-Item -Path $dest -force -ErrorAction SilentlyContinue
    Invoke-WebRequest -Uri $Url -OutFile $dest -UseBasicParsing
    Set-Content -Path $dest":Zone.Identifier" $ads
}
```

---


<!-- .slide: class="content" -->
## More setup

```powershell
## 2.Create a PS1 file in staging folder (any text will do but this is powershell extension)
echo "Write-Host ‘this is totally a resident file’" > C:\Perflogs\test.ps1

## 3.Modify shortname on a file
fsutil file setshortname C:\PerfLogs\psexec64.exe fake.exe

## 4. Create a process dumpOpen calculator (calc.exe)
calc.exe ; start-sleep 2
C:\PerfLogs\procdump64.exe -accepteula -ma win32calc C:\PerfLogs\calc.dmp
get-process | where-object { $_.Name -like "*win32calc*" } | Stop-Process

## 5. Create a zip file in staging folder
Compress-Archive -Path C:\PerfLogs\* -DestinationPath C:\PerfLogs\exfil.zip -CompressionLevel Fastest

## 6. Delete dmp,zip and ps1 files - deleted file discovery is important for later!
Remove-Item -Path C:\PerfLogs\*.zip, C:\PerfLogs\*.dmp, C:\PerfLogs\*.ps1
```

Note:

* Download and copy to staging folder C:\PerfLogs\
    * https://live.sysinternals.com/procdump64.exe
    * https://live.sysinternals.com/sdelete64.exe
    * https://live.sysinternals.com/psexec64.exe
* Add ADS to simulate Mark of the Web

Create a PS1 file in staging folder (any text will do but this is powershell extension)
```
echo "Write-Host ‘this is totally a resident file’" > C:\Perflogs\test.ps1
```

Modify short name on a file
```
fsutil file setshortname C:\PerfLogs\psexec64.exe fake.exe
```

Create a process dump Open calculator (`calc.exe`)
```
C:\PerfLogs\procdump64.exe -accepteula -ma calc C:\PerfLogs\calc.dmp
```

Create a zip file in staging folder - open `C:\Perflogs in Explorer`
highlight and select: Send to > Compressed (zipped) folder.
Delete `dmp.zip` and `ps1` files - deleted file discovery is important for later!
```
Remove-Item -Path C:\PerfLogs\*.zip, C:\PerfLogs\*.dmp, C:\PerfLogs\*.ps1
```


---

<!-- .slide: class="content" -->
## Exercise

<div class="container small-font">
<div class="col">

* Find contents of `C:\Perflogs`
* Review metadata of objects
* Explore leveraging filters
    * to target specific files or file types
    * to find files limited to a time frame

* Can you find the deleted files?
    * You may get lucky and have an unallocated file show.
    * Try `Windows.Forensics.Usn` with filters looking for suspicious
      extensions in our staging location!

</div>
<div class="col">
  <img src="/modules/ntfs_forensics/MFT_exercise1.png" style="bottom: inherit" class="inset" />
</div>
</div>

---

<!-- .slide: class="content small-font" -->

## The USN journal

* Update Sequence Number Journal or Change journal is maintained by
  NTFS to record filesystem changes.
* Records metadata about filesystem changes.
* Resides in the path $Extend\$UsnJrnl:$J

![](/modules/ntfs_forensics/usnj.png)

---

<!-- .slide: class="content" -->

## USN Journal
* Records are appended to the file at the end
* The file is sparse - periodically NTFS will remove the range at the start of the file to make it sparse
* Therefore the file will report a huge size but will actually only take about 30-40mb on disk.
* When collecting the journal file, Velociraptor will collect the sparse file.

---

<!-- .slide: class="content" -->
## Exercise - Windows.Forensics.Usn

<div class="container small-font">
<div class="col">

Target `C:\PerfLogs` with the `PathRegex` field.

* typically the USN journal only records filename and MFTId and
  ParentMFTId record. Velociraptor automatically reconstructs the
  expected path so the user can filter on path.
* This artifact uses FullPath results with “/”.

</div>
<div class="col">
  <img src="/modules/ntfs_forensics/Windows.Forensics.USN.png" style="bottom: inherit" class="inset" />
</div>
</div>

---

<!-- .slide: class="full_screen_diagram" -->
## Exercise - UsnJ solution

* There are many entries even for a simple file action like download to disk.

![](/modules/ntfs_forensics/USN_results.png)

---


<!-- .slide: class="content" -->
## Exercise - UsnJ solution

<div class="small-font">

* But these are simple to detect when you know what to look for!

<div class="container">
<div class="col">

![](/modules/ntfs_forensics/USN_groupby.png)

</div>
<div class="col">

![](/modules/ntfs_forensics/USN_delete.png)

</div>
</div>
</div>

---

<!-- .slide: class="content" -->
## Advanced NTFS: Alternate Data Stream

<div class="container small-font">
<div class="col">

* Most browsers attach an ADS to files downloaded from the internet.
* Use the VFS viewer to view the ADS of downloaded files.
* Use ADS Hunter to discover more interesting ADS
* Use `Windows.Analysis. EvidenceOfDownload` to identify downloaded
  files and unpacked ZIP files.

</div>
<div class="col">

<img src="/modules/ntfs_forensics/ADS_grupby.png" style="bottom: inherit" class="inset" />

</div>
</div>

Note:
 The inset shows typical frequency analysis of ADS naturally occurring

 What is the `Wof` stuff? https://devblogs.microsoft.com/oldnewthing/20190618-00/?p=102597
