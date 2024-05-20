package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/OpenDataTelemetry/timeseries-api/database"
	"github.com/gin-gonic/gin"
)

func GetWaterTankLevel(c *gin.Context) {
	var objs = []gin.H{}
	influxDB, err := database.ConnectToDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close()

	query := `
		SELECT *
		FROM "WaterTankLavel"
		WHERE
		time >= now() - interval '1 hour'
		ORDER BY time DESC;
	`

	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() { // Iterate over query response
		value := iterator.Value() // Value of the current row
		obj := gin.H{
			"fields": gin.H{
				"data_boardVoltage":            value["data_boardVoltage"],
				"fCnt":                         value["fCnt"],
				"rxInfo_altitude_0":            value["rxInfo_altitude_0"],
				"rxInfo_altitude_1":            value["rxInfo_altitude_1"],
				"rxInfo_latitude_0":            value["rxInfo_latitude_0"],
				"rxInfo_latitude_1":            value["rxInfo_latitude_1"],
				"rxInfo_longitude_0":           value["rxInfo_longitude_0"],
				"rxInfo_longitude_1":           value["rxInfo_longitude_1"],
				"rxInfo_loRaSNR_0":             value["rxInfo_loRaSNR_0"],
				"rxInfo_loRaSNR_1":             value["rxInfo_loRaSNR_1"],
				"rxInfo_rssi_0":                value["rxInfo_rssi_0"],
				"rxInfo_rssi_1":                value["rxInfo_rssi_1"],
				"txInfo_dataRate_spreadFactor": value["txInfo_dataRate_spreadFactor"],
				"txInfo_frequency":             value["txInfo_frequency"],
				"data_distance":                value["data_distance"],
			},
			"name": "WaterTankLavel",
			"tags": gin.H{

				"applicationID":              value["applicationID"],
				"devEUI":                     value["devEUI"],
				"fPort":                      value["fPort"],
				"nodeName":                   value["nodeName"],
				"rxInfo_mac_0":               value["rxInfo_mac_0"],
				"rxInfo_name_0":              value["rxInfo_name_0"],
				"txInfo_adr":                 "true",
				"txInfo_codeRate":            "4/5",
				"txInfo_dataRate_bandwidth":  "125",
				"txInfo_dataRate_modulation": "LORA",
			},
			"timestamp": value["time"],
		}
		// Convert the row to a gin.H map (JSON)
		objs = append(objs, obj) // Append the row to the objs slice
	}
	fmt.Println(len(objs))
	c.IndentedJSON(http.StatusOK, objs)
}

func GetHydrometer(c *gin.Context) {
	var objs = []gin.H{}
	influxDB, err := database.ConnectToDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close()

	query := `
		SELECT *
		FROM "Hidrometer"
		WHERE
		time >= now() - interval '1 hour'
		ORDER BY time DESC;
	`

	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() { // Iterate over query response
		value := iterator.Value() // Value of the current row
		obj := gin.H{
			"fields": gin.H{
				"data_boardVoltage":            value["data_boardVoltage"],
				"fCnt":                         value["fCnt"],
				"rxInfo_altitude_0":            value["rxInfo_altitude_0"],
				"rxInfo_altitude_1":            value["rxInfo_altitude_1"],
				"rxInfo_latitude_0":            value["rxInfo_latitude_0"],
				"rxInfo_latitude_1":            value["rxInfo_latitude_1"],
				"rxInfo_longitude_0":           value["rxInfo_longitude_0"],
				"rxInfo_longitude_1":           value["rxInfo_longitude_1"],
				"rxInfo_loRaSNR_0":             value["rxInfo_loRaSNR_0"],
				"rxInfo_loRaSNR_1":             value["rxInfo_loRaSNR_1"],
				"rxInfo_rssi_0":                value["rxInfo_rssi_0"],
				"rxInfo_rssi_1":                value["rxInfo_rssi_1"],
				"txInfo_dataRate_spreadFactor": value["txInfo_dataRate_spreadFactor"],
				"txInfo_frequency":             value["txInfo_frequency"],
				"data_counter":                 value["data_counter"],
			},
			"name": "Hydrometer",
			"tags": gin.H{

				"applicationID":              value["applicationID"],
				"devEUI":                     value["devEUI"],
				"fPort":                      value["fPort"],
				"nodeName":                   value["nodeName"],
				"rxInfo_mac_0":               value["rxInfo_mac_0"],
				"rxInfo_name_0":              value["rxInfo_name_0"],
				"txInfo_adr":                 "true",
				"txInfo_codeRate":            "4/5",
				"txInfo_dataRate_bandwidth":  "125",
				"txInfo_dataRate_modulation": "LORA",
			},
			"timestamp": value["time"],
		}
		// Convert the row to a gin.H map (JSON)
		objs = append(objs, obj) // Append the row to the objs slice
	}
	fmt.Println(len(objs))
	c.IndentedJSON(http.StatusOK, objs)
}

func GetArtesianWell(c *gin.Context) {
	var objs = []gin.H{}
	influxDB, err := database.ConnectToDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close()

	query := `
		SELECT *
		FROM "ArtesianWell"
		WHERE
		time >= now() - interval '1 hour'
		ORDER BY time DESC;
	`

	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() { // Iterate over query response
		value := iterator.Value() // Value of the current row
		obj := gin.H{
			"fields": gin.H{
				"data_boardVoltage":            value["data_boardVoltage"],
				"fCnt":                         value["fCnt"],
				"rxInfo_altitude_0":            value["rxInfo_altitude_0"],
				"rxInfo_altitude_1":            value["rxInfo_altitude_1"],
				"rxInfo_latitude_0":            value["rxInfo_latitude_0"],
				"rxInfo_latitude_1":            value["rxInfo_latitude_1"],
				"rxInfo_longitude_0":           value["rxInfo_longitude_0"],
				"rxInfo_longitude_1":           value["rxInfo_longitude_1"],
				"rxInfo_loRaSNR_0":             value["rxInfo_loRaSNR_0"],
				"rxInfo_loRaSNR_1":             value["rxInfo_loRaSNR_1"],
				"rxInfo_rssi_0":                value["rxInfo_rssi_0"],
				"rxInfo_rssi_1":                value["rxInfo_rssi_1"],
				"txInfo_dataRate_spreadFactor": value["txInfo_dataRate_spreadFactor"],
				"txInfo_frequency":             value["txInfo_frequency"],
				"data_pressure_0":              value["data_pressure_0"],
				"data_pressure_1":              value["data_pressure_1"],
			},
			"name": "ArtesianWell",
			"tags": gin.H{

				"applicationID":              value["applicationID"],
				"devEUI":                     value["devEUI"],
				"fPort":                      value["fPort"],
				"nodeName":                   value["nodeName"],
				"rxInfo_mac_0":               value["rxInfo_mac_0"],
				"rxInfo_name_0":              value["rxInfo_name_0"],
				"txInfo_adr":                 "true",
				"txInfo_codeRate":            "4/5",
				"txInfo_dataRate_bandwidth":  "125",
				"txInfo_dataRate_modulation": "LORA",
			},
			"timestamp": value["time"],
		}
		// Convert the row to a gin.H map (JSON)
		objs = append(objs, obj) // Append the row to the objs slice
	}
	fmt.Println(len(objs))
	c.IndentedJSON(http.StatusOK, objs)
}
