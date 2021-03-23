package scraper

import (
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestScraper(t *testing.T) {
	dsn := "geonames:geonames@tcp(127.0.0.1:3316)/geonames?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	scr, err := NewScraper(db, "test/fixtures/", false)
	if err != nil {
		t.Fatal(err)
	}
	err = scr.ScrapeCountries()
	if err != nil {
		t.Fatal(err)
	}
}
