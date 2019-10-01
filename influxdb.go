package main

import (
	"fmt"
	"log"
	"time"

	"github.com/influxdata/influxdb1-client/v2"
)

func PushData(te float64) {
	MyDB := "remo"
	username := ""
	password := ""
	MyMeasurement := "temperature"

	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://localhost:8086",
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		Precision: "s",
	})
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	tags := map[string]string{"remo": "3F"}
	fields := map[string]interface{}{
		"temperature": te,
	}

	pt, err := client.NewPoint(MyMeasurement, tags, fields, time.Now())
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	bp.AddPoint(pt)
	err = c.Write(bp)
	if err != nil {
		log.Fatalf("Unable to write value. %v", err)
	}

	fmt.Printf("success!\n")
}
