package main

// #go get 执行的时候优先使用官网代理：https://proxy.golang.org（这是 Go 1.13 之后的默认行为）
// #当从github下载不下来的时候，配置国内其他代理
// go env -w GO111MODULE=on
// go env -w GOPROXY=https://goproxy.cn,direct

// #恢复 Go 默认行为（推荐方式）,但是即便我本地配置了github的域名解析，能连接上github，但仍旧连接不上golang官方
// 所以还是使用以上代理
// go env -u GOPROXY（作用对象：仅限 go get、go mod 等 Go 命令。）
// go env -u GO111MODULE

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	s := make([]int, 0, 10)
	fmt.Println(s)

	// 1. 打开配置文件：config.json / app.yml 等
	file, _ := os.Open("config.json")
	defer file.Close()

	// 2. 把文件内容解析到 AppConfig 里
	//程序一启动，config.AppConfig 就已经装满了配置。

	json.NewDecoder(file).Decode(&AppConfig)
}
