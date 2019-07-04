

# {{.Name}}
字段名 | 值类型 | 说明
---------|----------|----------
{{range $i,$col:=.Columns}} {{$col.ColumnName}} | {{$col.Type}} | {{$col.Comment}}
{{end}}

## json示例
```json
{
{{range $i,$col:=.Columns}}    "{{$col.ColumnName}}":{{$col.ZeroValue}},
{{end}}
}
```

## 接口返回示例
```json
{
    "code":0,
    "message":"ok",
    "data":{
    {{range $i,$col:=.Columns}}    "{{$col.ColumnName}}":{{$col.ZeroValue}},
    {{end}}
    }
}
```

## 创建 
> post json 到 /api/{{.Name}}

## 全字段修改
> put json 到 /api/{{.Name}}/:id 

## 指定字段修改 
> patch json 到 /api/{{.Name}}/:id 
>> 例如  patch {"需要修改的字段":值} 到 /api/{{.Name}}/:id 


## 删除
> delete /api/{{.Name}}/:id 
 
## 获取一条
> get /api/{{.Name}}/:id 

## 获取多条
> get /api/{{.Name}}
> 查询条件以query参数方式组合,默认?offset=0&limit=0
>
> 返回
> ```json
> {
>     "code":0,
>     "message":"ok",
>     "data":{
>         "total":0, //当前查询条件下的总条数,用于分页
>         "data":[]
>     }
> }
> ```