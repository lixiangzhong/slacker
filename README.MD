# 安装

```sh
mkdir $GOPATH/src/github.com/lixiangzhong -p
cd $GOPATH/src/github.com/lixiangzhong
git clone git@d.dns.com:lixz/slacker.git
cd $GOPATH/src/github.com/lixiangzhong/slacker/slacker
go install
```

---

```sh
slacker #运行
```

```sh
NAME:
   slacker - 一键生成后台管理项目

USAGE:
   slacker [global options] command [command options] [arguments...]

VERSION:
   2.0.0

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

```sh
# flag

-addr=127.0.0.1:3306 -user=root -passwd=123456 -db=tablename
```

### 示例

```sh
slacker new app1 -addr=127.0.0.1:3306 -user=root -passwd=123456  -db=tablename
```

---

### 已有项目,添加代码

```sh
slacker add  <action>
```

> action :
>
> > m : 生成 models
> >
> > s : 生成 service
> >
> > d : 生成 dao
> >
> > v : 生成 vue
> >
> > c : 生成 controllers
> >
> > js : 生成 js
