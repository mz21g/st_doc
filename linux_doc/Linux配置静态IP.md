#### Linux 配置静态IP

```bash
# 1. vim /etc/network/interfaces 
auto eth0
iface eth0 inet static
address 10.10.10.252
netmask 255.255.255.0
gateway 10.10.10.1

# 2. vim /etc/resolv.conf
nameserver 114.114.114.114
nameserver 8.8.8.8

# 3. 重启networking服务
service networking restart
systemctl restart networking
/etc/init.d/networking restart

ps: dhclient eth0: 发出DHCP请求动态获取IP地址，如果是DHCP模式且开机后发现未获取到IP地址，可以用这个命令
```

