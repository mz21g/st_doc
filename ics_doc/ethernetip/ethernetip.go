package ethernetip

import (
	"errors"
	"unsafe"
	"time"
	"net"
	"github.com/zmap/zgrab2"
	"encoding/hex"
)

const MAX_FRAME_LEN = 4096

type ServiceInfo struct {
	ServiceVersion string //协议版本
	Status         bool   //协议识别结果
	Banner         string //Banner
}

//修改
type PortInfo struct {
	Protocol       string        `json:"trans_protocol"`
	Port           uint          `json:"port"`
	ServiceName    string        `json:"service_name"`
	ServiceVersion string        `json:"service_version"`
	Status         bool          `json:"status"`
	BannerList     []*BannerInfo `json:"banner_list"`
}

//修改
//BannerInfo is an struct of Banenr Info
type BannerInfo struct {
	ReqName  string `json:"req_name"`
	ReqData  string `json:"req_data"`
	Banner   string `json:"banner"`
	Encoding string `json:"encoding"`

	Timestamp  string `json:"timestamp"`
	PrivStatus bool   `json:"priv_status"`
	PrivError  string `json:"priv_error"`
}

//添加基本的端口信息
func addBasicPortInfo(scan *Connection, transProtocol string) {
	if len(scan.results) == 0 {
		scan.results = make([]*PortInfo, 0)
	}
	portInfo := new(PortInfo)
	portInfo.Protocol = transProtocol
	portInfo.ServiceName = scan.protocol
	portInfo.Port = scan.config.Port
	scan.results = append(scan.results, portInfo)
}

//添加协议基本信息
func addBasicBannerInfo(results []*PortInfo, reqData, reqName string) {
	portInfo := results[len(results)-1]
	bannerInfo := new(BannerInfo)
	bannerInfo.ReqData = reqData
	bannerInfo.ReqName = reqName
	bannerInfo.PrivError = ""
	bannerInfo.PrivStatus = true
	bannerInfo.Timestamp = time.Now().Format(time.RFC3339)
	portInfo.BannerList = append(portInfo.BannerList, bannerInfo)
}

//添加协议相关信息
func addServiceInfo(serviceVersion string, results []*PortInfo) {
	portInfo := results[len(results)-1]
	portInfo.ServiceVersion = serviceVersion
	portInfo.Status = true
}

//检查连接错误
func checkConnectionError(err error, results []*PortInfo) {
	portInfo := results[len(results)-1]
	bannerInfo := portInfo.BannerList[len(portInfo.BannerList)-1]
	if err != nil {
		bannerInfo.PrivError = bannerInfo.PrivError + "; " + err.Error()
		bannerInfo.PrivStatus = false
	}
}
func checkWriteError(err error, results []*PortInfo) error {
	portInfo := results[len(results)-1]
	bannerInfo := portInfo.BannerList[len(portInfo.BannerList)-1]
	if err != nil {
		bannerInfo.PrivError = bannerInfo.PrivError + "; " + err.Error()
		bannerInfo.PrivStatus = false
	}
	return err
}

//检查Read信息错误
func checkWriteReadErrorInfo(buf []byte, err error, results []*PortInfo) error {
	portInfo := results[len(results)-1]
	bannerInfo := portInfo.BannerList[len(portInfo.BannerList)-1]

	if err != nil {
		bannerInfo.PrivError = bannerInfo.PrivError + "; " + err.Error()
		bannerInfo.PrivStatus = false
	}
	if len(buf) == 0 {
		errZero := errors.New("packet error: result packet is zero")
		bannerInfo.PrivError = bannerInfo.PrivError + "; " + errZero.Error()
		bannerInfo.PrivStatus = false
		if err != nil {
			return err
		}
		return errZero
	}
	//只要不是结果长度为0，就添加标语信息
	addBannerInfo(buf, zgrab2.HEX_ENCODING, results)
	return nil
}

//添加标语信息
func addBannerInfo(bannerBytes []byte, encoding string, results []*PortInfo) {
	banner := ""
	switch encoding {
	case "hex":
		banner = B2Hex(bannerBytes)
	case "utf-8":
		banner = B2S(bannerBytes)
	default:
		banner = B2Hex(bannerBytes)
	}
	portInfo := results[len(results)-1]
	bannerInfo := portInfo.BannerList[len(portInfo.BannerList)-1]
	bannerInfo.Encoding = encoding
	bannerInfo.Banner = banner
}
func B2Hex(buf []byte) string {
	return hex.EncodeToString(buf)
}

//检查标语信息错误
func checkServiceRecognitionError(serviceInfo *ServiceInfo, err error, results []*PortInfo) error {
	if err != nil {
		return errors.New(zgrab2.SCAN_PROTOCOL_ERROR)
	}
	addServiceInfo(serviceInfo.ServiceVersion, results)
	return nil
}
func readResponse(conn net.Conn, data []byte) (buf []byte, err error) {
	w := 0
	for w < len(data) {
		writen, err := conn.Write(data[w:])
		w += writen
		if err != nil {
			return nil, err
		}
	}
	buf = make([]byte, MAX_FRAME_LEN)
	n, err := conn.Read(buf)
	if n == MAX_FRAME_LEN {
		err = errors.New("buffer size is lower than packet size")
	}
	return buf[0:n], err
}

//标语抓取
func getEthernetipPacket(scan *Connection) (err error) {
	//发送的数据包
	pkt := "\x63\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xc1\xde\xbe\xd1\x00\x00\x00\x00"
	addBasicBannerInfo(scan.results, pkt, "ethernetip_index")
	buf, err := readResponse(scan.conn, []byte(pkt))
	err = checkWriteReadErrorInfo(buf, err, scan.results)
	if err != nil {
		return err
	}
	serviceInfo, err := serviceRecognition(buf)
	err = checkServiceRecognitionError(serviceInfo, err, scan.results)
	return err

}

//协议识别算法
func serviceRecognition(buffer []byte) (serviceInfo *ServiceInfo, err error) {
	body := B2S(buffer)
	serviceInfo = &ServiceInfo{Banner: body}
	if buffer[0] != 0x63 || len(buffer) < 27 {
		serviceInfo.Status = false
		return serviceInfo, errors.New("protocol error:Invalid ethernetip packet")
	}
	serviceInfo.Status = true
	return serviceInfo, nil
}

//将byte数组转化为string，方便结果存储
func B2S(buf []byte) string {
	return *(*string)(unsafe.Pointer(&buf))
}