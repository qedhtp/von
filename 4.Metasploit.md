# Metasploit
## Metasploit Introduction 
command:  
```
msfconsole   //enter msf  

use <modules> or <number>  //enter modules' context

show <modules>  //show relevant modules 

back   //quie context 
info //info command relevant modules  

search <keyword>
search type:<modules> <type>

show options //list the required parammeters 

set PARAMETER_NAME VALUE

unset  //unset specific context 
unset all //unset all var

setg <modules>//set values for all modules
unsetg //clear any value set with setg

exploit or run  // in context
exploit -z //exploit in the background

check  //check if the target system is vulnerable without exploiting it 

background //background session
sessions //see the existing sessions
sessions -i //open specific session
```
## Metasploit: Exploitation 
### Scanning

command: 
```
search portscan //search portscan modules

search udp_sweep //UDP service identification  DNS NetBIOS

search smb_version //SMB scans 

```
### The Metasploit Database 
command:  
```
systemctl start postgresql //start postgreSQL 
msfdb init //initialize the Metasploit datebase

db_status //launch msfconsole check database status

workspace //list available workspaces
workspace -a //add a workspace 
workspace -d // delete workspace 
workspace <name> //switch workspace 
workspace -h //list available options 

help  //once launched database, show database backends commands

db_namp //sun nmap using the db_nmap, all esults will be saved to the database 

hosts //reach information relevant to hosts
services //reach information relevant to services

hosts -h //
services -h //
services -S <service> //search specific services

hosts -R //once the host information is stored in th database, using this to add host to the RHOSTS param.

real world:
finding available(live) hosts using the db_nmap command 
scanning these for further vulnerabilities or open ports

```
### Vulnerability Scanning 
command:  
```
step 1 fingerprint and recon the target then use search moduls
```
### Exploitation 
command: 
```
show payload //list other commands can use with that specific exploit 
set payload <number>  //set specific payload

rember:some payload will need new param. eg:reverse payload need Lhost

Ctrl+z //background
Ctrl+c //abort

sessions -h
sessions -i //follow by the session ID

hashdump
```
### Msfvenom 

practice again *   
command: 
```

```

### Metasploit: Meterpreter

#### Meterpreter Commands 
```
help //list all available commands

getuid //display the user with which metepreter is currently running

ps //list all process

migrate PID //migrate to another process

hashdump //list the content of the SAM database

search -f flag.txt  //search specific file

shell //launch a regular command-line shell
```
