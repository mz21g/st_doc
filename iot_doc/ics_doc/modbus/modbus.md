# Modbus 协议

## 协议简介

名称：Modbus

端口：502

传输层协议：TCP

协议分类：工控协议

协议描述：Modbus是一种串行通信协议，是Modicon公司（现在的施耐德电气 Schneider Electric）于1979年为使用可编程逻辑控制器（PLC）通信而发表。Modbus已经成为工业领域通信协议的业界标准（De facto），并且现在是工业电子设备之间常用的连接方式。

## 探测过程

向目标主机发送如下探测包：

> \x00\x00\x00\x00\x00\x05\x01\x2b\x0e\x01\x00

*以上探测包含义参考：[Modbus 协议中文介绍](http://10.10.2.22/iot_sat/iot_doc/tree/master/protocol/modbus)*

## 识别方式

响应报文以 “0x00, 0x00, 0x00, 0x00, 0x00” 开头

## 协议解析

Modbus 协议按位解析，可参考上边链接中的文档，目前我们解析的方式如下：

```python
MODUBS_EXCEPTION_CODES = {
    1: "ILLEGAL FUNCTION",
    2: "ILLEGAL DATA ADDRESS",
    3: "ILLEGAL DATA VALUE",
    4: "SLAVE DEVICE FAILURE",
    5: "ACKNOWLEDGE",
    6: "SLAVE DEVICE BUSY",
    8: "MEMORY PARITY ERROR",
    10: "GATEWAY PATH UNAVAILABLE",
    11: "GATEWAY TARGET DEVICE FAILED TO RESPOND"
}
. . .

def _decode_banner_inner(req_name, encoding, banner):
    banner_decode_info = ""
    try:
        banner_bytes = _get_banner_bytes(encoding, banner)
    except EncodingNotSupported as e:
        logging.warning(e.error_msg)
        return ""
    try:
        transaction_id, protocol_id, length, unit_id, function_id = struct.unpack('!HHHBB', banner_bytes[:8])
        if unit_id != 0:
            banner_decode_info = "Unexpected unit ID"
        elif function_id != 0x2b:
            ret_code = banner_bytes[8]
            banner_decode_info = MODUBS_EXCEPTION_CODES.get(ret_code, "Unknown exception {0}".format(ret_code))
        else:
            # 解析设备信息
            data = banner_bytes[8:]
            data_len = len(data)
            if data_len > 6:
                objects_count = data[5]
                start_index, end_index = 6, data_len
                for i in range(objects_count):
                    object_id = data[start_index]
                    if object_id == 0:  ##代表vendor_name
                        object_length = data[start_index + 1]
                        end_index = start_index + 2 + object_length
                        banner_decode_info += "VendorName:{0}\r\n".format(
                            data[start_index + 2:end_index].decode("utf-8", errors='ignore'))
                        start_index = end_index
                    elif object_id == 1:  ##代表ProductCode
                        object_length = data[start_index + 1]
                        end_index = start_index + 2 + object_length
                        banner_decode_info += "ProductCode:{0}\r\n".format(
                            data[start_index + 2:end_index].decode("utf-8", errors='ignore'))
                        start_index = end_index
                    elif object_id == 2:  ##代表Version
                        object_length = data[start_index + 1]
                        end_index = start_index + 2 + object_length
                        banner_decode_info += "Version:{0}\r\n".format(
                            data[start_index + 2:end_index].decode("utf-8", errors='ignore'))
                        start_index = end_index
        return banner_decode_info
    except Exception as e:
        logging.warning("解析MODBUS协议错误,原因:%s", repr(e))
    return ""
```

完整解析模块请查看 ```modbus.py```

## 附件

modbus.go：Modbus 协议探测脚本

modbus.py：Modbus 协议解析脚本