# scp-sl-whitelist
 A website ip whitelist for linux scp-sl server using ipset & iptables

# Dependencies

ipset & iptables

# How to use

1. Download code from git and build it

```bash
go build
```

2. Run the shell

```bash
sudo ./start.sh
```

3. Visit the website to join whitelist ` http://ip:2333/whitelist `