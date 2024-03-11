package settings_api

import (
	"github.com/gin-gonic/gin"
	"time1043/gvb-server/models/res"
)

func (SettingApi) SettingsInfoView(ctx *gin.Context) {
	res.Ok(map[string]string{}, "xxx", ctx)
	res.OkWithData(map[string]string{"id": "1043"}, ctx)
	res.OkWithMessage("很成功", ctx)

	res.Fail(map[string]string{}, "xxx", ctx)
	res.FailWithMessage("很失败", ctx)
	res.FailWithCode(res.SettingsError, ctx)
	res.FailWithCode(2, ctx)
}
