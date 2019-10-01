package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

type RemoDevices []struct {
	Name              string    `json:"name"`
	ID                string    `json:"id"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	FirmwareVersion   string    `json:"firmware_version"`
	TemperatureOffset int       `json:"temperature_offset"`
	HumidityOffset    int       `json:"humidity_offset"`
	Users             []struct {
		ID        string `json:"id"`
		Nickname  string `json:"nickname"`
		Superuser bool   `json:"superuser"`
	} `json:"users"`
	NewestEvents struct {
		Te struct {
			Val       float64   `json:"val"`
			CreatedAt time.Time `json:"created_at"`
		} `json:"te"`
	} `json:"newest_events"`
}

var remoDevices RemoDevices

func encodeJson4Remo(url string) (RemoDevices, error) {
	token := "Bearer token"
	m := map[string]string{
		"url":         url,
		"headerKey":   "Authorization",
		"headerValue": token}
	res, err := getResponse(m)
	if err != nil {
		return remoDevices, err
	}

	byteArray, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return remoDevices, err
	}
	defer res.Body.Close()

	if err := json.Unmarshal(byteArray, &remoDevices); err != nil {
		log.Fatalf("Error!: %v", err)
	}
	return remoDevices, err
}

func Remo(url string) (te float64) {
	res, err := encodeJson4Remo(url)
	if err != nil {
		log.Fatalf("Error!: %v", err)
	}
	for _, r := range res {
		te = r.NewestEvents.Te.Val

		jst := time.FixedZone("Asia/Tokyo", 9*60*60)

		fmt.Println("Name: " + r.Name)
		fmt.Println("ID: " + r.ID)
		fmt.Printf("UpdatedAt: %s\n", r.UpdatedAt.In(jst).Format("2006/01/02 15:04:05"))
		fmt.Printf("TemperatureOffset: %d\n", r.TemperatureOffset)
		fmt.Println("User ID: " + r.Users[0].ID)
		fmt.Println("User Nickname: " + r.Users[0].Nickname)
		fmt.Printf("User Superuser: %t\n", r.Users[0].Superuser)
		fmt.Printf("Temperature val: %f\n", r.NewestEvents.Te.Val)
		fmt.Printf("Temperature CreatedAt: %s\n", r.NewestEvents.Te.CreatedAt.In(jst).Format("2006/01/02 15:04:05"))
	}

	return te
}
