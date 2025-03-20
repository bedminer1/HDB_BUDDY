package graph

import (
	"time"

	"github.com/bedminer1/HDB_BUDDY/v3/internal/models"
)

func RecordsGraph(hdbRecords []models.HDBRecord, user models.User) []models.GraphData {
	graphData := []models.GraphData{}
	
	if len(hdbRecords) == 0 { return graphData}
	
	sortRecordsByDate(hdbRecords)
	startDate := hdbRecords[0].Time
	endDate := hdbRecords[len(hdbRecords)-1].Time

	currentDate := startDate
	lastPrice := 0.0
	lastScaledPrice := 0.0
	recordIndex := 0

	for !currentDate.After(endDate) {
		for recordIndex < len(hdbRecords) && hdbRecords[recordIndex].Time.Truncate(24*time.Hour).Equal(currentDate) {
			lastPrice = hdbRecords[recordIndex].PricePerArea
			lastScaledPrice = lastPrice * float64(user.FloorArea)
			recordIndex++
		}

		graphData = append(graphData, models.GraphData{
			PricePerArea:      lastPrice,
			ScaledResalePrice: lastScaledPrice,
			Date:              currentDate,
		})

		currentDate = currentDate.AddDate(0, 0, 1)
	}

	return graphData
}

func sortRecordsByDate(records []models.HDBRecord) {
	for i := range records {
		for j := i + 1; j < len(records); j++ {
			if records[i].Time.After(records[j].Time) {
				records[i], records[j] = records[j], records[i]
			}
		}
	}
}
