package main

import (
	"fmt"
	"github.com/fconf"
)

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