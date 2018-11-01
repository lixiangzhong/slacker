# 安装
```sh
cd $GOPATH/github.com/lixiangzhong
git clone git@d.dns.com:lixz/slacker.git
cd $GOPATH/github.com/lixiangzhong/slacker/slacker
go install
```

------
 

```sh
slacker #运行
```
```sh
NAME:
   slacker - 一键生成后台管理项目

USAGE:
   slacker [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
     new      create new project
     add      add table
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version

```

## 新建项目
```sh
slacker new <appname>
```

```
# flag

-addr=127.0.0.1:3306 -user=root -passwd=123456 -db=tablename
```

### 示例

```sh
slacker new app1
```
------

### 已有项目,添加代码
```sh
slacker add  <action>
```
> action :
>>        m : 生成models
>>        v : 生成vue 
>>        c : 生成controllers 
>>        js : 生成js 
>>        sql: 生成sql,会自动rebindata sql,自动插入 models.init.go:CreateTable
 