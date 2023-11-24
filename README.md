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
  nmap -PR -sn ipaddress/bumber //-PR only ARP scan -sn withoutport-scanning
  
  arp-scan -localnet or arp-scan -l   //scan all localnet
  more than one interface:
  arp-scan -I eth0 -l //-I sepcify interface
  apr-scan ipadress/number equal to nmap -PR -sn ipadress/number
  ```
  #### Nmap Host Discovery Using ICMP
  
    