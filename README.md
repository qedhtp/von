*_preface_*   
I use hexadecimal numbers to indicate the maximum title, decimal for the secondary title, octal for the third title, in order to make the article structure comparison cleaning, but also convenient for your own reference and your learning.
eg:





# Recon(Reconnaissance) # 
## 1. Passive Recon
### commmand: ###   
```
whois domain.com

nslookup domain.com
nslookup -type=A domain.com 1.1.1.1
nslookup -type=a domain.com 1.1.1.1 

dig domain.com 
dig domain.com MX  
dig @1.1.1.1 domain.com MX
```
>[public DNS servers](https://duckduckgo.com/?q=public+dns) 

| Query type | Result              |
| ---------- |-------------        | 
| A          | IPv4 Addresses      |
| AAAA       | IPv5 Addresses      |
| CNAME      | Canonical Name      |
| MX         | Mail Servers        |
| SOA        | Start of Authorirty |
| TXT        | TXT Records         |

### DNSDumpster ### 

* [https://dnsdumpster.com/](https://dnsdumpster.com/)  

### shodan ### 

* [https://www.shodan.io/](https://www.shodan.io/) 
* [Search Query Fundamentals](https://help.shodan.io/the-basics/search-query-fundamentals) 

## 2. Active Recon
### web browser 
  (1) Add a port to the address egg:https://127.0.0.1:8834/ will connect to 127.0.0.1 (localhost) at port 8834 via HTTPS protocol.   
  (2) Developer Tools  
  (3) add-ons    
* FoxyProxy
* [User-Agent Switcher and Manager](https://addons.mozilla.org/en-US/firefox/addon/user-agent-string-switcher) 
* Wappalyzer
### ping 
```
ping ipaddress
ping hostname
ping -n ipaddress //windows
ping -c ipaddress //linux
```
### traceroute
>The purpose of a traceroute is to find the IP addresses of the routers or hops that a packet traverses as it goes from your system to a target host
```
traceroute domain.com //linux
tracert domain.com //windows

```
### telnet 
```
telnet ipaddress port
GET / HTTP/1.1
host: telnet
```
### netcat 
```
nc ipaddress port //client
nc -lvnp port//server
```
*tip: Write a shell script and put them all together* 

# Network Security
## Nmap 
### Nmap Live Host Discovery
**Goal: knowing which hosts are up and avoid wasting our time port-scanning an offline host or an IP address not in use**  

basic command:
```
namp ipaddress domain.com example.com  //scan 3 IP address
nmap 10.11.12.15-20  //scan 6 IP address
10.10.0-255.101-125 //scan 6400 IP address
nmap 10.11.12.15/30  //scan 4 IP address
nmap -iL list_of_hosts.txt //scan form a file
nmap -sL targets //display detailed list of the hosts

```
#### Nmap host Discovery using APR
1. privileged user scan targets on local network(Ethernet)  
    * using ARP requests
2. privileged user scan targets outside the local network  
    * using ICMP echo requests
    * using TCP ACK to port 80
    * TCP SYN(Synchronize) to port 443
    * ICMP timestamp request
3. unprivileged user scan targets outside the local network
    * resorts to a TCP3-way handshake by sending SYN packets to ports 80 and 443  

  command:
  ```
  nmap -PR -sn ipaddress/bumber //-PR only ARP scan    -sn without port-scanning
  
  arp-scan -localnet or arp-scan -l   //scan all localnet
  more than one interface:
  arp-scan -I eth0 -l //-I sepcify interface
  apr-scan ipadress/number equal to nmap -PR -sn ipadress/number
  ```
#### Nmap Host Discovery Using ICMP
Rember, [ICMP](https://en.wikipedia.org/wiki/Internet_Control_Message_Protocol) recho requests tend to be blocked  
command:  
```
nmap -PE -sn ip_address/number //-PE ICMP echo request
nmap -PP -sn ip_address/number //-pp use ICMP timestamp request. if ICMP echo request be blocked  
nmap -Pm -sn ip_address/number

tip: try another if one is blocked
```
#### Nmap Host Discovery Using TCP and UDP
##### TCP
command:
```
TCP SYN Ping
nmap -PS -sn ip_address/number   //privileged user can choice no TCP 3-way(TCP ACK ping), vice versa. use -PS{port}  to specific port eg: -Ps21 -PS21-25 -PS22,80,443,8080

TCP ACK Ping
nmap -PA -sn ip_address/number //if is unprivileged user, Nmap attempt 3-way handshake. specific port is similar to TCP SYN ping   80 by default
   
```
##### UDP
send a UDP packet to open port is not expected to lead to any reply. But if send a UDP packet to a closed UDP port will get an ICMP port unreachable packet. this indicates that the target system is up and available.
command:
```
nmap -PM -sn ip_address/number
```
#### Using Reverse-DNS lookup

```
-n  //using reverse-DNS by default. using -n to skip this step
-R  //reverse-DNS look up for all hosts(include offline)
```
### Nmap basic port Scans
#### TCP Connect Scan
command:
```
nmap -sT ip_address/number //TCP scan -F only 100 most ports  -r consecutive scan instead of random scan. This is a unprivileged user only possible option to discover open TCP ports. 
```
#### TCP SYN Scan
No tcp connection was fully established with the target.   
command:
```
nmap -sS ip_address/number //default TCP scan mode, only privileged user can do.
```
#### UDP scan 
no handhake  
command:
```
nmap -sU ip_address/number
```
#### Fine-Tuning Scope and Performance 
basic option:
```
port:
-p22,80,443
-p1-1023
-p- //scan all
-F //most common port 100
--top-ports 10 //ten most common ports

control scan timing:
-T0  //paranoid(slowest)
-T1  //real engagements
-T2
-T3 //default,normal
-T4 //CTF
-T5

control packet rate:
--max-rate 15 //rate <=50packets/sec
--min-rate number or --min-rate=number  //eg:number=10 packets less than ten packets per second

parallel:
--max-parallelism number
--min-parallelism number  //egg:number=512  set at least 512 probes in parallel.
```
### Nmap Advanced Ports Scans  
#### TCP Null Scan, FIN Scan, Xmas Scan 
command:  
```
Null Scan:
nmap -sN ip_address    //no TCP flag. RST,ACK-->close, no reply-->open or firewall. Please rember-->many Nmap need root

FIN Scan:
nmap -sF ip_address //set FIN flag. RST,ACK-->close, noreply-->open or firewall. Nmap need no root

Xmas Scan:
mmap -sX ip_address //set the FIN,PSH and URG flags simultaneously. RST,ACK-->close, noreply-->open or firewall. Nmap need no root.
```
#### TCP Maimon 
Rember, this scan could not infer the port is close or open, but benefit for hacking mindset.  
command:
```
nmap -sM ip_address //set FIN/ACK-->most target reply RST, whether open or close
```
#### TCP ACK,Window,and Custom Scan
command:
```
TCP ACK:
nmap -sA ip_address //infer which port are not being blocked. discover firewall rulesets and configuration.

Window Scan:
nmap -sW ip_address //similar to TCP ACK scan, but results might be different. 

Custom Scan:
nmap --scanflags RSTSYNFIN //set RST,SYN,FIN scan. note: firewall not block does't mean port open. Maybe firewall need update to reflect recent service changes.
```
#### Spoofing and Decoys
command:   
```
Spoofing:
nmap -s spoofed_ip ip_address //basic command
mmap -e net_interface -Pn -S spoofed_ip ip_addresss //-e specify the network interface  -Pn disable ping scan. ???ping is active host discover?
nmap --spoof-mac //same subnet(same Ethernet802.3   same WiFi 802.11)
rember: Spoofing works need a certain condition.


Decoys:
nmap -D 10.10.0.1,10.10.0.2,me ip_address //easy to infer attack's ip
nmap -D 10.10.0.1,10.10.0.2,random,random,me ip_address
```
#### Fragmented Packets  

for evade firewall or IDS, divide IP header data 
```
-f //diveded into 8 bytes or less    always choice
-ff //16 bytes
--mtu //default bytes
--data-length num //increase packet size, num is a append packets bytes.
```
#### idle/Zombie Scan 
command: 
```
nmap -sI zombie_ip ip_address //difference=1, close or filter. difference=2, open
```
#### Getting More Details 
command:
```
--reason //display reason
-v  //display verbose output
-vv //even more verbosity
-d //dubugging
-dd //more dubugging
```
### Nmap Post Port Scans  
#### Service Detection 
command:  
```
nmap -sV ip_address  // -sV collect and determine service and version for the open ports. note: this force TCP 3-way. So the default stealth SYN scan not work.

--version-intensity level  //range between 0 to 9
--version-light  //intensity 2 
--version-all  //intensity 9

```
####  OS Detection and Traceroute
command:  
```
nmap -sS -O ip_address //note: Less precise

nmap -sS --tracroute ip_address //note: TTL reduce. different form traceroute or tracert command
```
#### Nmap Scripting Engine (NSE)
command:
```
nmap -sS -n --script "script" ip_address //--script=default or -sC is default category. there is another categories,eg:auth,broadcast,brute,vuln......
```
#### Saving the output 
command:
```
normal:
nmap -sS -sV -O -oN 10.10.178.34_scan ip_address 

Greable:
nmap -sS -sV -O -oG 10.10.178.34_scan ip_address

XML:
nmap -sS -sV -O -oX 10.10.178.34_scan ip_address

-oA for all
```
### Protocols and Servers
| Protocol | TCP Port |Application(s)|Data Security|
| ---------|----------| ---------    |--------     |
| FTP      | 21       | File Transfer| Cleartext   |
| HTTP     | 80       | Worldwide Web| Cleartext   |
| IMAP     | 143      | Email (MDA)  | Cleartext   |
| POP3     | 110      | Email (MDA)  | Cleartext   |
| SMTP     | 25       |Email (MTA)   |Cleartext    |
| Telnet   | 23       |Remote Access | Cleartext   |
#### Telnet
remote login, not encrypted
#### 
http
#### 
ftp
command:
```
ftp ip_address  //anonymous ascii get filename
```
#### SMTP 
command:  
```
telnet ip_address 25
```
#### POP3 


#### IMAP



















