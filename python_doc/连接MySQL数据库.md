#### python 连接 MySQL数据库

*通过PyMySql的connect()方法声明一个MySql连接对象db，连接成功后，需要再调用cursor()方法获得MySql的操作游标，利用游标来执行SQL语句*

```python
import pymysql
db = pymysql.connect(host='192.168.14.235', user='root', password='cyberpeckerIIE', port=58840)
# 调用cursor()方法获得MySQL的操作游标
cursor = db.cursor()
cursor.execute("USE cyberpecker")
cursor.execute("SELECT project_id, task_id, ip From `task_result_info` where ip='%s'" % input_ip)
results = cursor.fetchall()
db.close()
```
