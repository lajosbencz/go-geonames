package web

import "github.com/lajosbencz/go-geonames/models"

type GeocodeResult struct {
	models.Location
	IsCity   bool    `gorm:"column:is_city" json:"is_city"`
	Distance float64 `gorm:"column:distance" json:"distance"`
	Radius   uint32  `gorm:"column:radius" json:"radius"`
	Score    float32 `gorm:"column:score" json:"score"`
}
