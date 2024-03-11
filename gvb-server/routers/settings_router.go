package routers

import (
	"time1043/gvb-server/api"
)

func (router RouterGroup) SettingsRouter() {
	settingApi := api.ApiGroupApp.SettingApi
	router.GET("settings", settingApi.SettingsInfoView) // 127.0.0.1:8080/api/settings
}
