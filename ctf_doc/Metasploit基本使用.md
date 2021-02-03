## 基本使用

```bash
search - 根据关键字搜索某模块
show payloads – 查看该模块适用的所有载荷代码 
show targets – 查看该模块适用的攻击目标类型
use – 进入使用某渗透攻击模块 
```

## Metasploit 常见网络服务扫描

```bash
# Telnet 服务查点
use auxiliary/scanner/telnet/telnet_version
# SSH 服务查点
use auxiliary/scanner/ssh/ssh_version
# Oracle 服务查点
use auxiliary/scanner/oracle/tnslsnr_version
```

## SSH 服务弱口令爆破

```bash
use auxiliary/scanner/ssh/ssh_login
```

## psnuffle 口令嗅探

*截获常见协议的身份认证过程，并将用户名和口令信息记录下来*

```bash
use auxiliary/sniffer/psnuffle
```

