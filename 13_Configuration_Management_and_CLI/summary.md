# Basic Command Line Interface
Command line is adalah antarmuka cepat dan berbasis teks yang digunakan pengembang untuk berkomunikasi secara lebih efektif dan efisien dengan komputer guna menyelesaikan serangkaian tugas yang lebih luas.

## Mengapa menggunakan Command Line
* Kontrol granular dari suatu OS atau aplikasi
* Manajemen lebih cepat dari sejumlah besar sistem operasi
* Kemampuan untuk menyimpan skrip untuk mengotomatiskan tugas-tugas rutin
* Menghubungkan pengetahuan untuk membantu pemecahan masalah, seperti masalah koneksi jaringan

## Command Line Interface
Shell = CLI for OS'Services
* UNIX Shell
* Command Prompt (MDSOS)

Other app CLI 
* Python REPL
* MySQL CLI Client
* Mongo Shell
* redis-CLI
* etc 

## Intro to UNIX Shell
Example Shell : 
```
me@linuxbox:~$

```
* (me) -> your username
* (linuxbox) -> your hostname
* (~) -> your current path (home)
* ($) -> Shell for normal username
* (#)-> Shell for root user


#### Normal User vs Root User
* Normal User = Allowed to manupulating /home/username dir only
* Root User = Allowed to manupulating /dir (all directory)
* Normal User + Sudoers = Allowed to act as root temporary


## Unix Shell Most Popular Command
#### Directory 
* pwd
* ls
* mkdir
* cd 
* rm 
* cp 
* mv
* ln

#### Files
* create : touch
* view : head, cat, tail,less
* editor : vim, nano
* permission : chown, chmod
* Different : diff 

#### Network 
* ping
* ssh
* netstat
* nmap
* ip addr (ifconfig)
* wget
* curl
* telnet
* netcat

#### Ubility 
* man 
* env
* echo
* date
* which
* watch
* sudo 
* history
* grep 
* locate
