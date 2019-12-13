# sect

Get indented section of text from stdin by regular expression
TAB counted as one space
If matching line ends with opening braces, first non-section line with corresponding closing brace is also printed

Example:
```# ip addr | sect ens160\.75
14: ens160.75@ens160: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
    link/ether 00:0c:29:10:a5:f9 brd ff:ff:ff:ff:ff:ff
    inet 172.24.56.120/25 brd 172.24.56.127 scope global ens160.75
       valid_lft forever preferred_lft forever
    inet6 fe80::20c:29ff:fe10:a5f9/64 scope link 
       valid_lft forever preferred_lft forever
```
