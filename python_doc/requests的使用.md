#### python Requests的使用

```python
import requests

headers = {
    "User-Agent": "Windows"
}

params = {
    'WEPEncryption': 'AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABBBB'
}

r = requests.get("http://httpbin.org/get", headers=headers, params=params)

data = {'name': 'germey', 'age':'22'}
r = requests.post("http://httpbin.org/post", data=data)
r = requests.put("http://httpbin.org/put")
r = requests.delete("http://httpbin.org/delete")
r = requests.head("http://httpbin.org/get")
r = requests.options("http://httpbin.org/get")

print(r.headers)
print(r.cookies)
print(r.status_code)
print(r.url)
print(r.history)

# 关闭SSL证书验证
r = requests.get('https://www.12306.cn', verify=False)
# 身份认证
r = requests.get("http://localhost:5000", auth=('username', "password"))
```

