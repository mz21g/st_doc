#### TCP 收发数据

```python
import socket

target_ip = ''
target_port = 80

payload = 'A' * 1024
rtsp_request = (f"SETUP rtsp://{target_ip}:{target_port}/play1.sdp/trackID=1 RTSP/1.0\r\n"
                "CSeq: 8\r\n"
                F"Authorization: Basic {payload}\r\n"
                "User-Agent: LibVLC/3.0.11 (LIVE555 Streaming Media v2016.11.28)\r\n"
                "Transport: RTP/AVP;unicast;client_port=61444-61445\r\n\r\n")

try:
    client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    client.connect((target_ip, target_port))

    client.send(bytes(rtsp_request, encoding='utf-8'))
    response = client.recv(4096)
    print(response)
    client.close()
except Exception as e:
    print(f'error: {e}')
```

#### UDP 收发数据

```python
import socket

target_ip = ''
target_port = 80

try:
    client = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    client.sendto(b"AAABBBCCC", (target_ip, target_port))
    data, addr = client.recvfrom(4096)
    print(data)
except Exception as e:
    pass
```

#### TCP 服务器

```python
bind_ip = ""
bind_port = 9999

server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
server.bind((bind_ip, bind_port))

server.listen(5)

print(f'Listening on {bind_ip}:{bind_port}')


def handle_client(client_socket):

    request = client_socket.recv(4096)    
    print(f'Received: {request}')

    client_socket.send(b"ACK!")
    client_socket.close()


while True:
    client, addr = server.accept()
    print(f'Accepted connection from:{addr[0]}:{addr[1]}')

    client_handler = threading.Thread(target=handle_client, args=(client,))
    client_handler.start()
```

#### 在数据链路层收发数据

```python
#!/usr/bin/env python
# -*- coding:utf-8 -*-
# author:chenyanzhi
# datetime:2020/11/20 17:16
# software: PyCharm
import binascii
import socket
import uuid


def get_mac_address():
    """获取本机MAC地址"""
    mac = uuid.UUID(int=uuid.getnode()).hex[-12:]
    return ''.join([mac[e:e + 2] for e in range(0, 11, 2)])


def get_resp(src_mac, raw_socket):
    """获取目标设备响应。

    数据链路层收到的是整个网卡的数据报文，需要过滤一下才能获取到发送的数据帧所对应的响应报文。
    """
    resp = raw_socket.recv(4096).hex()
    flag = 100
    while True:
        if resp[:12] == src_mac or resp[:12] == dst_mac:
            return resp
        elif flag == 0:
            raise Exception("没有收到正确响应")
        else:
            flag -= 1
            resp = raw_socket.recv(4096).hex()

ETH_P_ALL = 3
dst_mac = '00:1b:1b:e6:10:ff'.replace(':', '')
src_mac = get_mac_address()
nameOfStation = '53656375726974792e636865636b'  # Security.check

raw_socket = socket.socket(socket.AF_PACKET, socket.SOCK_RAW)
raw_socket.bind(('eth0', ETH_P_ALL))

pn_dcp = (f'{dst_mac}{src_mac}810000008892fefd0400040000040000001a020200100001'
          f'{nameOfStation}05020002000000000000')
payload = binascii.unhexlify(pn_dcp)

print(f'send packet: {pn_dcp}')
raw_socket.send(payload)
print('ready to receive')
resp = get_resp(src_mac, raw_socket)
print(f'resp is: {resp}')

raw_socket.close()

```

