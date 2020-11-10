
# README concernant l'implémentation locale à l'UVSQ

L'original du code est accessible sur 
https://www.socketloop.com/tutorials/golang-get-current-time-from-the-internet-time-server-ntp-example

Le code a été légèrement modifié afin de changer :
 - le code de retour,
 - le erveur
 - le mode verbose de verification

- ntpcheck : le binaire dans https://git.dsi.uvsq.fr/thiecail/ntpcheck

## Exemple de sortie : 
```
PS C:\dev\src\projects\ntpcheck> go run .\main.go -v -server time.windows.com
Getting Ntp time from time.windows.com
Ntp time :      Wed Nov 11 00:37:01 CET 2020
Local time :    Wed Nov 11 00:37:01 CET 2020
Delta is 0s
```


## Vérification quotidienne via Monit : 
```
# cat  /etc/monit.d/go-check-certs

check program ntpcheck with path "/local/sbin/ntpcheck"
        every "30 * * * 1-5"
        if status > 0 then alert

```
A noter que si le code de retour via monit ne change pas (comprendre pas de corrections), l'alerte ne sera pas mise à jour. Donc, pour un problème persistant, il vaut mieux déplacer le service dans le fichier exclude.txt avec un commentaire.

