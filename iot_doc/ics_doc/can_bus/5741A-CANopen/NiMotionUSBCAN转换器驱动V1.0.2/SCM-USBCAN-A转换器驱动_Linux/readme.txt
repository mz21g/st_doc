1��ִ������������linux�°�װlibusb1.0
   sudo apt-get install libusb
   sudo apt-get install libusb-1.0
2��ִ�������������̬�⵽/usr/lib
   sudo cp libNiMotionUSBCAN.so /usr/lib
3��������ͨ�û�����USB�豸
   (1) ִ�������������usbfs�û���
       sudo groupadd usbfs
   (2) ִ����������û���ӵ�usbfs�û���
       sudo usermod -a -G usbfs qihuanming
   (3) ��61-NiMotionUSBCAN.rules�ļ�������/etc/udev/rules.dĿ¼
       sudo cp 61-NiMotionUSBCAN.rules /etc/udev/rules.d