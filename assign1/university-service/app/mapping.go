package app

import (
	"university/controllers/ping"
	"university/controllers/univeristyinfocontroller"
)

func MapUrls() {

	router.GET("/ping", ping.Ping)
	router.GET("universityinfo/find", univeristyinfocontroller.FindInfo)
	router.GET("universityinfo/delete", univeristyinfocontroller.DeleteInfo)
	router.PUT("universityinfo/update", univeristyinfocontroller.UpdateInfo)
	router.POST("universityinfo/create", univeristyinfocontroller.CreateInfo)

}
