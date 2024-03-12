package settings_api

import (
	"github.com/gin-gonic/gin"
	"time1043/gvb-server/config"
	"time1043/gvb-server/core"
	"time1043/gvb-server/global"
	"time1043/gvb-server/models/res"
)

func (SettingApi) SettingsInfoUpdateView(ctx *gin.Context) {
	var cr config.SiteInfo
	err := ctx.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, ctx)
		return
	}
	global.Config.SiteInfo = cr

	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), ctx)
		return
	}
	res.OkWith(ctx)
}
