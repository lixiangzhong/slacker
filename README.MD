# 安装
```sh
cd $GOPATH/github.com/lixiangzhong
git clone git@d.dns.com:lixz/slacker.git
cd $GOPATH/github.com/lixiangzhong/slacker/slacker
go install
```

------

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
 