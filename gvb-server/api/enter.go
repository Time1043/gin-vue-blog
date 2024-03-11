package api

import "time1043/gvb-server/api/settings_api"

type ApiGroup struct {
	SettingApi settings_api.SettingApi
}

var ApiGroupApp = new(ApiGroup)
