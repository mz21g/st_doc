#### Scapy 的使用

```bash
# 从pcap文件中读取数据包
pcap = rdpcap("C:\\Users\\CLAY\\Desktop\\pcap\\changeDeviceName.pcapng")
# 制作可视化数据包图像
pcap[13].pdfdump(layer_shift=1)
```

