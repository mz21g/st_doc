#### 常用函数

```bash
# 返回一个bytes对象，其中包含根据格式字符串format打包的值v1, v2, ...参数个数必须与格式字符串所要求的的值完全匹配
struct.pack(format, v1, v2, ...)

# 根据格式字符串format从缓冲区buffer解包（假定由pack(format, ...)打包）。结果为一个元祖，即使其只有一个条目。缓冲区的字节大小必须匹配格式所要求的大小
struct.unpack(format, buffer)

# 返回与格式字符串format相对应的结构的大小（pack(format, ...)所产生的字节串对象的大小）
struct.calcsize(format)
```

#### 字节顺序

格式字符串（format）的第一个字符可根据下表指示打包数据的字节顺序、大小和对齐方式

| 字符 | 字节顺序     | 大小     | 对齐方式 |
| :--- | :----------- | :------- | :------- |
| @    | 按原字节     | 按原字节 | 按原字节 |
| =    | 按原字节     | 标准     | 无       |
| <    | 小端         | 标准     | 无       |
| >    | 大端         | 标准     | 无       |
| ！   | 网络（大端） | 标准     | 无       |

如果第一个字符不是其中之一，则假定为“@”

可使用 `sys.byteorder` 来检查自己系统的字节顺序

#### 格式字符串的含义

| 格式 | C类型              | Python类型      | 标准大小 |
| ---- | ------------------ | --------------- | -------- |
| x    | 填充字节           | 无              |          |
| c    | char               | 长度为1的字节串 | 1        |
| b    | signed char        | 整数            | 1        |
| B    | unsigned char      | 整数            | 1        |
| ?    | _Bool              | bool            | 1        |
| h    | short              | 整数            | 2        |
| H    | unsigned short     | 整数            | 2        |
| i    | int                | 整数            | 4        |
| I    | unsigned int       | 整数            | 4        |
| l    | long               | 整数            | 4        |
| L    | unsigned long      | 整数            | 4        |
| q    | long long          | 整数            | 8        |
| Q    | unsigned long long | 整数            | 8        |
| e    |                    | 浮点数          | 2        |
| f    | float              | 浮点数          | 4        |
| d    | double             | 浮点数          | 8        |
| s    | char[]             | 字节串          |          |
| P    | void *             | 整数            |          |

```python
# 打包和解包三个整数的示例
>>> from struct import *
>>> pack('hhl', 1, 2, 3)
b'\x00\x01\x00\x02\x00\x00\x00\x03'
>>> unpack('hhl', b'\x00\x01\x00\x02\x00\x00\x00\x03')
(1, 2, 3)
>>> calcsize('hhl')
8

# 解包的字段可通过将它们赋值给变量或将结果包装为一个具名元组来命名:
>>> record = b'raymond   \x32\x12\x08\x01\x08'
>>> name, serialnum, school, gradelevel = unpack('<10sHHb', record)

>>> from collections import namedtuple
>>> Student = namedtuple('Student', 'name serialnum school gradelevel')
>>> Student._make(unpack('<10sHHb', record))
Student(name=b'raymond   ', serialnum=4658, school=264, gradelevel=8)
```

