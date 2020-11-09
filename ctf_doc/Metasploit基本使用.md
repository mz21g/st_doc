#### Metasploit 常见网络服务扫描

```bash
# Telnet 服务查点
use auxiliary/scanner/telnet/telnet_version
# SSH 服务查点
use auxiliary/scanner/ssh/ssh_version
# Oracle 服务查点
use auxiliary/scanner/oracle/tnslsnr_version
```

#### SSH 服务弱口令爆破

```bash
use auxiliary/scanner/ssh/ssh_login
```

#### psnuffle 口令嗅探

*截获常见协议的身份认证过程，并将用户名和口令信息记录下来*

```bash
use auxiliary/sniffer/psnuffle
```

