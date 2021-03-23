package web

import (
	"context"
	"log"
	"net/http"

	"github.com/lajosbencz/go-geonames/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewServer(ctx context.Context, listen string, dsn string) (err error) {

	// dsn := "geonames:geonames@tcp(127.0.0.1:3316)/geonames?charset=utf8mb4&parseTime=True&loc=Local"
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

	srv := &http.Server{
		Addr:    listen,
		Handler: mux,
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %+s\n", err)
		}
	}()

	log.Println("Server listening on " + listen)

	<-ctx.Done()

	log.Printf("server stopped")

	// ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer func() {
	// 	cancel()
	// }()

	// if err = srv.Shutdown(ctxShutDown); err != nil {
	// 	log.Fatalf("server Shutdown Failed:%+s", err)
	// }

	log.Printf("server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}

	return
}
