package main

import (
	"time1043/gvb-server/core"
	"time1043/gvb-server/global"
	"time1043/gvb-server/routers"
)

func main() {
	core.InitConf() // 读取配置文件
	//fmt.Println(global.Config) // 验证成功

	global.Log = core.InitLogger() // 初始化日志
	//global.Log.Warnln("ha ha ha ")
	//global.Log.Error("ha ha ha ")
	//global.Log.Info("ha ha ha ")
	//logrus.Warnln("ha ha ha ")
	//logrus.Error("ha ha ha ")
	//logrus.Info("ha ha ha ")

	global.DB = core.InitGorm() // 连接数据库
	//println(global.DB)          // 验证过

	// 路由配置
	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("gvb-server 运行在： %s", addr)
	router.Run(addr) // 验证成功
}
