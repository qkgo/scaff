package util

import (
	"log"
	"testing"
)

func TestGetDetailDay(t *testing.T) {
	totalOnlineDayDetail := GetDetailDayByMonth(3)
	log.Println("1 res.length:", len(totalOnlineDayDetail))
	OnlineDayDetail := RandomDay(totalOnlineDayDetail, 3)
	log.Println("2 res:", OnlineDayDetail)
	log.Println("2 res.length:", len(OnlineDayDetail))
	OnlineDayDetail = RandomDay(totalOnlineDayDetail, 31)
	log.Println("2 res:", OnlineDayDetail)
	log.Println("2 res.length:", len(OnlineDayDetail))
}
