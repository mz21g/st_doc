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