# goTool 三个实用工具

## 环境配置
- 初始化：
```
mkdir -p /data/htdocs/go_project/goTool
cd /data/htdocs/go_project/goTool
go mod init github.com/goTool
```
- 安装Cobra：
```
go get -u github.com/spf13/cobra@v1.0.0
```

## 工具一：单词格式转换
- 该工具支持五种格式转换模式，可使用以下命令查看
```
go run main.go help word
```
- 例如：使用模式1进行全部单词转为大写
```
go run main.go word -s=cnyygj -m=1

# 输出
2020/07/31 20:26:50 输出结果： CNYYGJ
```

## 工具二：时间工具
- 获取当前时间
```
go run main.go time now
```
- 时间推算
```
go run main.go time calc -c="2020-07-31 20:42:26" -d=5m

# 输出
2020/07/31 20:44:43 输出结果： 2020-07-31 20:47:26, 1596228446

go run main.go time calc -c="2020-07-31 20:42:26" -d=-2h

# 输出
2020/07/31 20:45:38 输出结果： 2020-07-31 18:42:26, 1596220946
```

## 工具三：生成表结构体
- 示例
```
go run main.go sql struct --username=账号 --password=密码 --db=test --table=student

# 输出
type student struct {
	// id
	id int32  `json:id`     
    // cid
	cid int32  `json:cid`     
    // gid
	gid int32  `json:gid`     
    // name
	name string  `json:name`
}

func (model student) TableName() string {
	return "student"
}
```
