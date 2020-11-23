# EthernetIP 协议

## 协议简介

名称：EthernetIP

端口：44818

传输层协议：TCP

协议分类：工控协议

协议描述：工业以太网协议 (Ethernet/IP) 是由ODVA所开发并得到了罗克韦尔自动化的强大支持。它使用已用于ControlNet和DeviceNet的控制和信息协议 (CIP) 为应用层协议。

## 探测过程

向目标主机发送如下探测包：

> \x63\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xc1\xde\xbe\xd1\x00\x00\x00\x00

## 识别方式

响应报文首位是 “0x63” **或者**响应报文长度大于 27

## 协议解析

EthernetIP 协议按位解析，可参考上边链接中的文档，目前我们解析的方式如下：

```python
def decode_ethernet(self, banner, encoding) -> bool:
        """
        解析ethernetip协议，并返回解析成功与否
        :param banner:
        :param encoding:
        :return:返回协议的解析成功与否
        """
        if encoding != HEX_ENCODING:
            return False
        bytes_banner = binascii.unhexlify(banner)
        if len(bytes_banner) > 65:
            self.encapsulation_version = int.from_bytes(bytes_banner[30:32], "little")
            self.sin_family = int.from_bytes(bytes_banner[32:34], "big")
            self.sin_port = int.from_bytes(bytes_banner[34:36], "big")
            self.sin_addr = '.'.join(str(i) for i in struct.unpack('>BBBB', bytes_banner[36:40]))
            self.sin_zero = ''.join('{0:0x}'.format(i) for i in bytes_banner[40:48])
            vendor_start, vendor_end = 48, 50
            self.vendor_code = int.from_bytes(bytes_banner[vendor_start:vendor_end], byteorder="little")
            self.vendor = VENDOR_MAP.get(self.vendor_code, "")
            device_type_start, device_type_end = 50, 52
            self.device_type_code = int.from_bytes(bytes_banner[device_type_start:device_type_end], byteorder="little")
            self.device_type = DEVICE_TYPE_MAP.get(self.device_type_code, "")
            product_code_start, product_code_end = 52, 54
            self.product_code = int.from_bytes(bytes_banner[product_code_start:product_code_end], byteorder='little')
            revision_first_index, revision_second_index = 54, 55
            self.revision = '{0}.{1}'.format(bytes_banner[revision_first_index], bytes_banner[revision_second_index])
            #self.revision = '{0}.{1:02}'.format(bytes_banner[revision_first_index], bytes_banner[revision_second_index])
            status_hex_start, status_hex_end = 56, 58
            status_list = struct.unpack("<2B", bytes_banner[status_hex_start:status_hex_end])
            # 此处是反序的
            self.status = '0x' + ''.join("{0:02x}".format(s) for s in reversed(status_list))
            serial_number_start, serial_number_end = 58, 62
            serial_list = struct.unpack("<4B", bytes_banner[serial_number_start:serial_number_end])
            self.serial_number = "0x" + ''.join("{0:0x}".format(serial) for serial in reversed(serial_list))
            self.product_name_length = struct.unpack('B', bytes_banner[62:63])[0]
            self.product_name = bytes_banner[63:-1].decode("utf-8", "ignore")
            self.state = "0x" + bytes_banner[-1:].hex()
```

完整解析模块请查看 ```ethernertip.py```

## 附件

ethernertip.go：EthernetIP 协议探测脚本

ethernertip.py：EthernetIP 协议解析脚本