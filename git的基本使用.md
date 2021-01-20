#### GIt

```bash
git init # 初始化仓库
git status # 查看仓库的状态
git add # 向暂存区中添加文件，暂存区是提交之前的一个临时区域
git commit # 将当前暂存区中的文件实际保存到仓库的历史记录中，通过这些记录可以在工作树中复原文件
git log --graph # 以图表形式输出提交日志
git branch -a # 显示分支名并确认当前所在分支（包括远端分支）
git checkout -b feature-A # 创建并切换到feature-A分支
git checkout master # 切换到master分支
# 添加远程仓库
git remote add origin http://10.10.2.22/iot_sat/ics_exploit.git
git push -u origin master # 推送至远程仓库 -u 在推送的同时，将该分支设置为本地仓库当前分支的上游

note: 如果想推送至master以外分支，需要以同名形式push至远程仓库
git checkout -b feature-D
git push -u origin feature-D

# 获取远程仓库，默认处于master分支下
# 同时系统会将origin设置成该远程仓库的标识符
# 本地master分支与远端仓库master分支在内容上是一样的。
git clone http://10.10.2.22/iot_sat/ics_exploit.git

# 获取远程feature-D分支
git checkout -b feature-D
git pull origin feature-D

git branch -d <BranchName> # 删除本地分支
git push origin --delete <BranchName> # 删除远程分支

# 放弃修改，强制覆盖本地代码
git fetch --all
git reset --hard origin/cyz
git pull origin cyz

# 强制提交本地分支覆盖远程分支
git push origin <BranchName>  --force
```

