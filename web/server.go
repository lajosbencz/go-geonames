package web

import (
	"context"
	"net/http"

	"github.com/lajosbencz/go-geonames/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewServer(ctx context.Context, listen string, dsn string) (srv *http.Server, err error) {

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	db.Set("gorm:table_options", "ENGINE=MyISAM").AutoMigrate(&models.Country{}, &models.Language{}, &models.Location{}, &models.LocationAlternate{})

	handlers := &Handlers{
		DB: db,
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(handlers.HandleIndex))
	mux.Handle("/util/update", http.HandlerFunc(handlers.HandleScrape))
	mux.Handle("/list/countries", http.HandlerFunc(handlers.HandleListCountries))
	mux.Handle("/list/languages", http.HandlerFunc(handlers.HandleListLanguages))
	mux.Handle("/geocode", http.HandlerFunc(handlers.HandleGeocode))

	srv = &http.Server{
		Addr:    listen,
		Handler: mux,
	}

	return
}
