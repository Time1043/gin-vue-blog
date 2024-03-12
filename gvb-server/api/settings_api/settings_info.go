package settings_api

import (
	"github.com/gin-gonic/gin"
	"time1043/gvb-server/global"
	"time1043/gvb-server/models/res"
)

func (SettingApi) SettingsInfoView(ctx *gin.Context) {
	res.OkWithData(global.Config.SiteInfo, ctx)
}
