## 安装Django
```bash
pip3 install Django
```
## 创建项目
```bash
django-admin startproject MyDjango
```

其中“MyDjango”是项目名称，创建好项目后会包含5个py文件

![image-20210213103300501](/Users/chenyanzhi/Desktop/st_doc/python_doc/image/Django_startproject.png)

各个文件说明如下:

* manage.py：命令行工具，内置多种方式与项目进行交互，在命令行窗口输入python manage.py help可以查看该工具的指令信息
* \_\_init\_\_.py：初始化文件
* setting.py：项目的配置文件，项目的所有功能都需要在该文件中进行配置
* urls.py：项目的路由设置，设置网站的具体网址内容
* wsgi.py：Python服务器网关接口（Python Web Server Gateway Interface），是Python应用与web服务器之间的接口，用于Django项目在服务器上的部署和上线，一般不需要修改

## 创建项目应用（App）

```bash
python3 manage.py startapp index
```

其中“index”是应用的名称，项目会新建index文件夹，该文件夹内可看到多个py文件和migrations文件夹

![image-20210213125734100](/Users/chenyanzhi/Desktop/st_doc/python_doc/image/Django_startapp.png)

各个文件说明如下：

* migrations: 用于生成数据迁移文件，通过数据迁移文件可自动在数据库里生成相应的数据表
* \_\_init\_\_.py: index文件夹的初始化文件
* admin.py: 用于设置当前app的后台管理功能
* apps.py: 当前app的配置信息，一般情况下无需修改
* models.py: 定义数据库的映射类，每个类可以关联一张数据表，实现数据持久化，即MTV里面额模型（Model）
* tests.py: 自动化测试的模块，用于实现单元测试
* views.py: 视图文件，处理功能的业务逻辑，即MTV里面的视图（Views）

## 启动项目

```bash
python3 manage.py runserver 8001
```



