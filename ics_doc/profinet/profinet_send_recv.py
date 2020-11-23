#!/usr/bin/env python
# -*- coding:utf-8 -*-
# author:chenyanzhi
# datetime:2020/11/20 17:16
# software: PyCharm
import binascii
import socket
import uuid

ETH_P_ALL = 3


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


dst_mac = '00:1b:1b:e6:10:ff'.replace(':', '')
src_mac = get_mac_address()
nameOfStation = '53656375726974792e636865636b'  # Security.check

raw_socket = socket.socket(socket.AF_PACKET, socket.SOCK_RAW)
raw_socket.bind(('eth0', ETH_P_ALL))

pn_dcp = (f'{dst_mac}{src_mac}810000008892fefd0400040000040000001a020200100001'
          f'{nameOfStation}05020002000000000000')
req_device_name = (f'{dst_mac}{src_mac}810000008892fefe05000400000f00800004ffff000'
                   f'00000000000000000000000000000000000000000000000000000')
payload = binascii.unhexlify(pn_dcp)

print(f'send packet: {pn_dcp}')
raw_socket.send(payload)
print('ready to receive')
resp = get_resp(src_mac, raw_socket)
print(f'resp is: {resp}')

raw_socket.close()
