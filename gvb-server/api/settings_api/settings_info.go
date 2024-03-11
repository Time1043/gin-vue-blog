package settings_api

import "github.com/gin-gonic/gin"

func (SettingApi) SettingsInfoView(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"msg": "xxx"})
}
