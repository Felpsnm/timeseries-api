package main

import (
	"github.com/OpenDataTelemetry/timeseries-api/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default() // Create a new gin router instance
	api_sl := r.Group("/api/timeseries/v0.1/smartcampusmaua/SmartLights")
	{
		api_sl.GET("", controller.GetSmartLights)
		api_sl.GET("deviceName/:nodename", controller.GetSmartLightbyNodeName)
		api_sl.GET("deviceId/:devEUI", controller.GetSmartLightbyDevEUI)
	}

	api_sc_wtl := r.Group("/api/timeseries/v0.1/smartcampusmaua/GetWaterTankLevel")
	{
		api_sc_wtl.GET("", controller.GetWaterTankLevel)
	}

	api_sc_h := r.Group("/api/timeseries/v0.1/smartcampusmaua/GetHydrometer")
	{
		api_sc_h.GET("", controller.GetHydrometer)
	}

	api_sc_aw := r.Group("/api/timeseries/v0.1/smartcampusmaua/GetArtesianWell")
	{
		api_sc_aw.GET("", controller.GetArtesianWell)
	}

	r.Run(":8888")
}
