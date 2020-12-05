#### aliases命令(列举所有的简写命令)

```bash
gef➤  aliases
[+] Aliases defined:
pf                              →  print-format
uf                              →  disassemble
dps                             →  dereference
dq                              →  hexdump qword
dd                              →  hexdump dword
[...]
```

#### 常用命令

| 命令         | 全称          | 解释                           |
| ------------ | ------------- | ------------------------------ |
| l            | list          | 查看源码                       |
| b            | break         | 设置断点                       |
| clear        | clear         | 删除断点                       |
| disable b_id | disable b_id  | 使中断失效，b_id 为中断编号    |
| enable b_id  | enable b_id   | 使中断生效，b_id 为中断编号    |
| n            | next          | 单条语句执行                   |
| c            | continue      | 继续运行程序，下一个断点处停止 |
| p            | print         | 打印                           |
| q            | quit          | 退出GDB                        |
| i            | info          | 查看信息                       |
| run          | run           | 开始调试程序                   |
| r            | register info | 查看寄存器信息                 |



#### 打断点的三种方式

1. 根据函数名，查找符号(symbol)设置断点

   b func_name

2. 根据代码行设置断点

   b 3 <------- 在第三行设置断点

3. 根据运行时的地址设置断点

   b *0xffffd1e4 <------ *必须加在地址前，0xffffd1e4为函数指针的地址