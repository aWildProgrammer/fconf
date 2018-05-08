## fconf包读取配置文件数据
这是一个用于引入配置文件的包，加载类似于ini风格的配置文件。
与INI配置文件风格一样 根据顺序读取文件和每一行 如果在行首出现了;号，则认为是配置文件的注释。
除此之外，还提供多种可选类型。

## INI演示文件（DEMO）
```
[mysql]
db1.Name = testMysqlDB
db1.Host = 127.0.0.1
db1.Port = 3306
db1.User = root
db1.Pwd = test

; 测试INI风格的注释
; 这两行数据的前前面加入了分号，因此，这些数据被认为是注释，将不会影响配置文件正常情况

[tcp]
Port=3309
```

## 引用示例：
```
func main() {
	c, err := fconf.NewFileConf("./demo.ini")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(c.String("mysql.db1.Host"))
	fmt.Println(c.String("mysql.db1.Name"))
	fmt.Println(c.String("mysql.db1.User"))
	fmt.Println(c.String("mysql.db1.Pwd"))

	// 取得配置时指定类型
	port, err := c.Int("mysql.db1.Port")
	if err != nil {
		panic("~")
	}
	fmt.Println(port)
}
```
## 运行以上代码将输出（OUTPUT）

```
127.0.0.1
testMysqlDB
root
test
3306
```
