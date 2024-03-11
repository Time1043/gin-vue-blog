package main

import (
	"fmt"
	"time1043/gvb-server/core"
	"time1043/gvb-server/global"
)

func main() {
	core.InitConf() // 读取配置文件
	fmt.Println(global.Config)
}
