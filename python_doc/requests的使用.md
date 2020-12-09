#### GET 请求

```python
import requests
headers = {"User-Agent": "Windows"}
params = {'WEPEncryption': 'AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABBBB'}
r = requests.get("http://httpbin.org/get", headers=headers, params=params)
```

#### POST 请求

```python
data = {'name': 'germey', 'age':'22'}
r = requests.post("http://httpbin.org/post", data=data)
```

#### 其他请求

```python
r = requests.put("http://httpbin.org/put")
r = requests.delete("http://httpbin.org/delete")
r = requests.head("http://httpbin.org/get")
r = requests.options("http://httpbin.org/get")

print(r.headers)
print(r.cookies)
print(r.status_code)
print(r.url)
print(r.history)
```

#### 关闭SSL证书验证

```python
# 关闭SSL证书验证
r = requests.get('https://www.12306.cn', verify=False)
```

#### 身份认证

````python
# basic64 认证
r = requests.get("http://localhost:5000", auth=('username', "password"))
````

#### 抓取二进制数据

```python
r = requessts.get("http://github.com/favicon.ico")
with open('favicon.ico', 'wb') as f:
    f.write(r.content)
```

#### 会话维持

```python
import requests
s = requests.Session
s.get('http://httpbin.org/cookies/set/number/123456789')
r = s.get('http://httpbin.org/cookies')
print(r.text)
```

#### 超时设置

```python
# 将超时时间设置为1秒，1秒内无响应则抛出异常
r = requests.get("https://www.taobao.com", timeout = 1)
# 实际上，请求分为两个阶段，即连接和读取,上面设置的timeout用作这两者之和，分别指定可传入一个元组
r = requets.get("https://www.taobao.com", timeout=(5.11, 30))
# 如果想永久等待则将timeout设置为 None，或者不设置直接留空
r = requets.get("https://www.taobao.com", timeout=None)
```

#### 文件上传

```python
import requests
files = {'file': open('favicon.ico', 'rb')}
r = requests.post('http://httpbin.org/post', files=files)
print(r.text)
```

#### 禁止重定向

```python
r = requests.get('http://github.com', allow_redirects=False)
```

