# sect

Get indented section of text from stdin by regular expression
TAB counted as one space
If matching line ends with opening braces, first non-section line with corresponding closing brace is also printed

Example:
```
# ip addr | sect ens160\.75
14: ens160.75@ens160: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
    link/ether 00:0c:29:10:a5:f9 brd ff:ff:ff:ff:ff:ff
    inet 172.24.56.120/25 brd 172.24.56.127 scope global ens160.75
       valid_lft forever preferred_lft forever
    inet6 fe80::20c:29ff:fe10:a5f9/64 scope link 
       valid_lft forever preferred_lft forever
```

Example (JSON):
```
# curl -s 'http://localhost:8181/debug' | sect "dev_list"
    "dev_list": {
      "10.0.11.248": {
        "id": "serial:76310777A09A",
        "proc_error": "",
        "proc_result": "done in 6 ms",
        "state": "run",
        "time": 1576229006
      },
      "10.0.6.48": {
        "id": "lldp:f8f0827867be",
        "proc_error": "",
        "proc_result": "done in 8 ms",
        "state": "run",
        "time": 1576229041
      },
      "10.0.6.49": {
        "id": "lldp:f8f082736bf5",
        "proc_error": "",
        "proc_result": "done in 24 ms",
        "state": "run",
        "time": 1576229022
      }
    },
```

Building:

setup GOPATH variable first

```
go get github.com/ShyLionTjmn/sect
cd $GOPATH/src/github.com/ShyLionTjmn/sect
go build
cp sect /usr/local/bin/
```
