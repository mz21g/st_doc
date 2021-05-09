#### 创建一个独立的Python运行环境，命名为`venv`

```bash
virtualenv --no-site-packages venv
```

命令`virtualenv`就可以创建一个独立的Python运行环境，我们还加上了参数`--no-site-packages`，这样，已经安装到系统Python环境中的所有第三方包都不会复制过来，这样，我们就得到了一个不带任何第三方包的“干净”的Python运行环境。

#### 进入环境

```bash
source venv/bin/activate
```

注意到命令提示符变了，有个`(venv)`前缀，表示当前环境是一个名为`venv`的Python环境。

#### 退出当前环境

```bash
deactivate
```

