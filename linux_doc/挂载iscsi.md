#### iscsi_挂载

```bash
# 目标发现
sudo iscsiadm -m discovery -t sendtargets -p 192.168.12.20:3260
# 登录
sudo iscsiadm -m node -T xxxx -p 192.168.12.20:3260 --login
# 查看可挂载分区
sudo fdisk -l
# 挂载目录
sudo mount /dev/sdb1 ./iscsi_test/
# 卸载目录
sudo umount iscsi_test/
# 登出
sudo iscsiadm -m node -T xxxxx -p 192.168.12.20:3260 --logout
```

