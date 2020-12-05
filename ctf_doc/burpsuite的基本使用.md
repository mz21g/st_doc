#### 配置

1. 在chrome浏览器的设置页面点击”打开您计算机的代理设置“

![image-20201204201208788](C:\Users\CLAY\Desktop\st_doc\ctf_doc\image\burpsuite\修改本地网络代理.png)

2. 选中”使用代理服务器“，将地址和端口分别修改为”127.0.0.1“和 8080 ，最后点击”保存“

![image-20201204201655292](C:\Users\CLAY\Desktop\st_doc\ctf_doc\image\burpsuite\修改本地网络代理2.png)

3. 在burpsuite中加入 127.0.0.1:8080 的监听

![image-20201204213658958](C:\Users\CLAY\Desktop\st_doc\ctf_doc\image\burpsuite\配置BurpSuite监听.png)

#### 修改返回数据

1. 在代理的设置选项中勾选拦截服务端数据

![image-20201204214021809](C:\Users\CLAY\Desktop\st_doc\ctf_doc\image\burpsuite\拦截服务端数据.png)

2. 访问目标网页，在代理的拦截选项中右键选择”Respense to this requests“

![image-20201204214401237](C:\Users\CLAY\Desktop\st_doc\ctf_doc\image\burpsuite\拦截返回消息.png)

3. 然后forward发送，可以看到服务器端返回的数据

![image-20201204223135881](C:\Users\CLAY\Desktop\st_doc\ctf_doc\image\burpsuite\服务器端返回的数据.png)

4. 这里还没有传递到浏览器，可以在这里修改数据，这里我删掉了script代码，并在body里边添加了我的名字

![image-20201204221512477](C:\Users\CLAY\Desktop\st_doc\ctf_doc\image\burpsuite\服务器端返回的数据.png)

5. 点击forward

![image-20201204221212108](C:\Users\CLAY\Desktop\st_doc\ctf_doc\image\burpsuite\修改数据后的效果.png)

6. 数据修改完毕

**有一些网址的后台登陆验证，如果检测登陆失败会返回302状态码，然后js跳转到首页，这个时候可以修改返回包的信息，状态码改成200，然后删除js代码即可进入后台。**