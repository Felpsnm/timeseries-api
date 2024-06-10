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
		WITH ranked_data AS (
			SELECT *,
				ROW_NUMBER() OVER (PARTITION BY "nodeName" ORDER BY time DESC) AS row_num
			FROM "WaterTankLavel"
			WHERE time >= now() - interval '1 hour'
		)
		SELECT *
		FROM ranked_data
		WHERE row_num = 1
		ORDER BY "nodeName" ASC;
	`

	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() { // Iterate over query response
		value := iterator.Value() // Value of the current row
		obj := gin.H{
			"fields": gin.H{
				"data_distance": value["data_distance"],
			},
			"tags": gin.H{
				"nodeName": value["nodeName"],
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
		WITH ranked_data AS (
			SELECT *,
			ROW_NUMBER() OVER (PARTITION BY "nodeName" ORDER BY time DESC) AS row_num
			FROM "Hidrometer"
			WHERE time >= now() - interval '1 hour'
		)
		SELECT *
		FROM ranked_data
		WHERE row_num = 1
		ORDER BY "nodeName" ASC;
	`

	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() { // Iterate over query response
		value := iterator.Value() // Value of the current row
		obj := gin.H{
			"fields": gin.H{
				"data_counter": value["data_counter"],
			},
			"tags": gin.H{
				"nodeName": value["nodeName"],
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
		WITH ranked_data AS (
			SELECT *,
			ROW_NUMBER() OVER (PARTITION BY "nodeName" ORDER BY time DESC) AS row_num
			FROM "ArtesianWell"
			WHERE time >= now() - interval '1 hour'
		)
		SELECT *
		FROM ranked_data
		WHERE row_num = 1
		ORDER BY "nodeName" ASC;
	`

	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() { // Iterate over query response
		value := iterator.Value() // Value of the current row
		obj := gin.H{
			"fields": gin.H{
				"data_pressure_0": value["data_pressure_0"],
				"data_pressure_1": value["data_pressure_1"],
			},
			"tags": gin.H{
				"nodeName": value["nodeName"],
			},
			"timestamp": value["time"],
		}
		// Convert the row to a gin.H map (JSON)
		objs = append(objs, obj) // Append the row to the objs slice
	}
	fmt.Println(len(objs))
	c.IndentedJSON(http.StatusOK, objs)
}
