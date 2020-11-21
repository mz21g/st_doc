#### Python中str()与repr()函数的区别

`print` 语句结合 `str()` 函数实际上是调用了对象的 `__str__` 方法来输出结果。

 `print` 结合 `repr()` 实际上是调用对象的 `__repr__` 方法输出结果。

```python
>>> from datetime import datetime
>>> now = datetime.now()
>>> print(str(now))
2017-04-22 15:41:33.012917
>>> print(repr(now))
datetime.datetime(2017, 4, 22, 15, 41, 33, 12917)
```

* str() 的输出追求可读性，输出格式要便于理解，适合用于输出内容到用户终端。
* repr() 的输出追求明确性，除了对象内容，还需要展示出对象的数据类型信息，适合开发和调试阶段使用。



