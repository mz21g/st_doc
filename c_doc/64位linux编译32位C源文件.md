## fatal error: bits/libc-header-start.h: 没有那个文件或目录

错误原因：

在64位的机器上编译32位的C源文件需要安装对应的gcc 32位的库

解决办法：

安装multilib 库，这个库可以在64位机器上产生32位的程序或者库文件

```bash
sudo apt-get install gcc-multilib
sudo apt-get install g++-multilib
```



## 编译方法

```bash
gcc -m32 -fno-stack-protector -no-pie -o fs1 1.c
```

