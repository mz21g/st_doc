package modbus

import (
	"bytes"
	"errors"
	"github.com/zmap/zgrab2"
	"net"
	"unsafe"
)

//检查读写错误错误
func checkWriteReadErrorInfo(buf []byte, err error) error {
	errStr := ""
	if err != nil {
		errStr += err.Error()
	}
	if buf == nil || len(buf) == 0 {
		errStr += ";packet error: result packet is zero"
		if err != nil {
			return errors.New(errStr)
		}
	}
	return nil
}

//检查标语信息错误
func checkServiceRecognitionError(serviceInfo *zgrab2.ServiceInfo, err error) error {
	if err != nil {
		return errors.New(zgrab2.SCAN_PROTOCOL_ERROR)
	}
	return nil
}

func socketWriteAndRead(conn net.Conn, data []byte, maxBytes int) (buf []byte, err error) {
	w := 0
	for w < len(data) {
		writen, err := conn.Write(data[w:])
		w += writen
		if err != nil {
			return nil, err
		}
	}
	buf = make([]byte, maxBytes)
	n, err := conn.Read(buf)
	if n == maxBytes {
		err = errors.New("buffer size is lower than packet size")
	}
	return buf[0:n], err
}

func readBanner(conn net.Conn, bannerInfo *zgrab2.BannerInfo, maxBytes int) (err error, serviceInfo *zgrab2.ServiceInfo) {
	// get package to send request
	bindQuery := "\x00\x00\x00\x00\x00\x05\x01\x2b\x0e\x01\x00"
	byteQuery := []byte(bindQuery)
	buf, err := socketWriteAndRead(conn, byteQuery, maxBytes)
	err = checkWriteReadErrorInfo(buf, err)
	if err != nil {
		bannerInfo.SetError(err)
		return err, nil
	}
	bannerInfo.SetBanner(buf, zgrab2.HEX_ENCODING)
	//识别协议并存Banner
	serviceInfo, err = serviceRecognition(buf)
	return checkServiceRecognitionError(serviceInfo, err), serviceInfo
}

//协议识别
func serviceRecognition(buffer []byte) (serviceInfo *zgrab2.ServiceInfo, err error) {
	serviceInfo = &zgrab2.ServiceInfo{}
	value := []byte{0x00, 0x00, 0x00, 0x00, 0x00}
	if bytes.HasPrefix(buffer, value) {
		serviceInfo.Status = true
		return serviceInfo, nil
	}
	serviceInfo.Status = false
	return serviceInfo, errors.New(zgrab2.SCAN_PROTOCOL_ERROR)
}

//将byte数组转化为string，方便结果存储
func B2S(buf []byte) string {
	return *(*string)(unsafe.Pointer(&buf))
}