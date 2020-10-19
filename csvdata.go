package main

import (
	"encoding/csv"
	"log"
	"os"
)

type Shooting struct {
	Id                   string `json:"id"`
	Title                string `json:"title"`
	Location             string `json:"location"`
	Date                 string `json:"date"`
	Incident_Area        string `json:"incident_area"`
	Open_Close_Location  string `json:"open_close_location"`
	Target               string `json:"target"`
	Cause                string `json:"cause"`
	Summary              string `json:"summary"`
	Fatalities           string `json:"fatalities"`
	Injured              string `json:"injured"`
	Total_Victims        string `json:"total_victims"`
	Policeman_Killed     string `json:"policeman_killed"`
	Age                  string `json:"age"`
	Employeed_Y_N        string `json:"employeed_y_n"`
	Employed_At          string `json:"employed_at"`
	Mental_Health_Issues string `json:"mental_health_issues"`
	Race                 string `json:"race"`
	Gender               string `json:"gender"`
	Latitude             string `json:"latitude"`
	Longitude            string `json:"longitude"`
}

var shootings []Shooting

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func readData(filePath string) {
	file, err1 := os.Open(filePath)
	checkError("Unable to read input file "+filePath, err1)
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err2 := csvReader.ReadAll()
	checkError("Unable to parse file as CSV for "+filePath, err2)
	defer file.Close()

	shootings = []Shooting{}

	for _, record := range records {
		shooting := Shooting{
			Id:                   record[0],
			Title:                record[1],
			Location:             record[2],
			Date:                 record[3],
			Incident_Area:        record[4],
			Open_Close_Location:  record[5],
			Target:               record[6],
			Cause:                record[7],
			Summary:              record[8],
			Fatalities:           record[9],
			Injured:              record[10],
			Total_Victims:        record[11],
			Policeman_Killed:     record[12],
			Age:                  record[13],
			Employeed_Y_N:        record[14],
			Employed_At:          record[15],
			Mental_Health_Issues: record[16],
			Race:                 record[17],
			Gender:               record[18],
			Latitude:             record[18],
			Longitude:            record[19]}
		shootings = append(shootings, shooting)
	}
	file.Close()
}

func writeData(filePath string) {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC, 0644)
	checkError("Cannot create file", err)
	defer file.Close()

	file.Seek(0, 0)
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, shooting := range shootings {
		record := []string{
			shooting.Id,
			shooting.Title,
			shooting.Location,
			shooting.Date,
			shooting.Incident_Area,
			shooting.Open_Close_Location,
			shooting.Target,
			shooting.Cause,
			shooting.Summary,
			shooting.Fatalities,
			shooting.Injured,
			shooting.Total_Victims,
			shooting.Policeman_Killed,
			shooting.Age,
			shooting.Employeed_Y_N,
			shooting.Employed_At,
			shooting.Mental_Health_Issues,
			shooting.Race,
			shooting.Gender,
			shooting.Latitude,
			shooting.Longitude}
		err := writer.Write(record)
		checkError("Cannot write to file", err)
	}
	writer.Flush()
	file.Close()
}
