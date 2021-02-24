# go-ntp-check

go-ntp-check is a ntp cli used to automatically check ntp deviation between local host and a ntp server.

## Code base 
Code based on 
https://www.socketloop.com/tutorials/golang-get-current-time-from-the-internet-time-server-ntp-example

Code was modified in order to change 
- code return
- server
- verbose mode 

## Build
```shell
# go build # as usual for golang 
```

## Basic output example : 
```
PS C:\dev\src\projects\ntpcheck> go run .\main.go -v -server time.windows.com
Getting Ntp time from time.windows.com
Ntp time :      Wed Nov 11 00:37:01 CET 2020
Local time :    Wed Nov 11 00:37:01 CET 2020
Delta is 0s
```


## Automating checks via via Monit : 
```
# cat  /etc/monit.d/ntpcheck

check program ntpcheck with path "/local/sbin/ntpcheck"
        every "30 * * * 1-5"
        if status > 0 then alert

```

If you don't correct  (return code), monit alerts won't change !