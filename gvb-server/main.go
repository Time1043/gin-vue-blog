package main

import (
	"time1043/gvb-server/core"
	"time1043/gvb-server/global"
	"time1043/gvb-server/routers"
)

func main() {
	core.InitConf()                // 读取配置文件
	global.Log = core.InitLogger() // 初始化日志
	global.DB = core.InitGorm()    // 连接数据库
	router := routers.InitRouter() // 路由配置
	addr := global.Config.System.Addr()
	global.Log.Infof("gvb-server 运行在： %s", addr)
	router.Run(addr) // 验证成功
}
