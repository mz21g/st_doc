#### sqlmap 用法

查看是否存在SQL注入漏洞

```bash
sqlmap -u 'http://TARGET_IP?id=XXX' --cookie='TARGET_COOKIE'
```

探测存放Web应用数据的数据库名称

```bash
sqlmap -u "http://TARGET_IP?id=XXX" --cookie="TARGET_COOKIE" --dbs -v 0
```

