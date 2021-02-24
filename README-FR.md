# README les modifications

La base du code est accessible sur 
https://www.socketloop.com/tutorials/golang-get-current-time-from-the-internet-time-server-ntp-example

Le code a été légèrement modifié afin de changer :
 - le code de retour,
 - le serveur
 - le mode verbose de verification

## Exemple de sortie : 
```shell
# go run .\main.go -v -server time.windows.com
Getting Ntp time from time.windows.com
Ntp time :      Wed Nov 11 00:37:01 CET 2020
Local time :    Wed Nov 11 00:37:01 CET 2020
Delta is 0s
```


## Vérification quotidienne via Mmonit : 
```shell
# cat  /etc/monit.d/go-ntp-check

check program go-ntp-check with path "/local/sbin/go-ntp-check"
        every "30 * * * 1-5"
        if status > 0 then alert

```
A noter que si le code de retour via monit ne change pas (comprendre pas de corrections), l'alerte ne sera pas mise à jour. 