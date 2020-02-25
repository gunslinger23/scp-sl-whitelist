#!/bin/bash

#firewall
iptables -D INPUT -p udp -m set --match-set scp-whitelist src --dport 7777 -j ACCEPT
iptables -D INPUT -p tcp -m set --match-set scp-whitelist src --dport 7777 -j ACCEPT
iptables -D INPUT -p udp --dport 7777 -j DROP
iptables -D INPUT -p tcp --dport 7777 -j DROP
iptables -A INPUT -p udp -m set --match-set scp-whitelist src --dport 7777 -j ACCEPT
iptables -A INPUT -p tcp -m set --match-set scp-whitelist src --dport 7777 -j ACCEPT
iptables -A INPUT -p udp --dport 7777 -j DROP
iptables -A INPUT -p tcp --dport 7777 -j DROP

#start web whitelist
./whitelist