package web

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/lajosbencz/go-geonames/models"
	"github.com/lajosbencz/go-geonames/scraper"
	"gorm.io/gorm"
)

func writeErrorString(w http.ResponseWriter, err string) {
	log.Println("ERROR! " + err)
	w.WriteHeader(500)
	w.Write([]byte(err))
}

func writeError(w http.ResponseWriter, err error) {
	writeErrorString(w, err.Error())
}

func writeJson(w http.ResponseWriter, d interface{}) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(d)
	if err != nil {
		log.Println("ERROR! failed to encode JSON: " + err.Error())
	}
}

type Handlers struct {
	DB *gorm.DB
}

func (h *Handlers) HandleIndex(w http.ResponseWriter, r *http.Request) {
	html := `
<a href="/geocode">Geocode &lt;search[,exact,country,lat,lng,radius]]&gt;</a>
<hr/>
<a href="/list/languages">List languages</a>
<br/>
<a href="/list/countries">List countries</a>
	`
	w.Write([]byte(html))
}

func (h *Handlers) HandleListCountries(w http.ResponseWriter, r *http.Request) {
	var list []models.Country
	h.DB.Find(&list)
	d := make(map[string]interface{})
	d["list"] = list
	writeJson(w, d)
}

func (h *Handlers) HandleListLanguages(w http.ResponseWriter, r *http.Request) {
	var list []models.Language
	h.DB.Find(&list)
	d := make(map[string]interface{})
	d["list"] = list
	writeJson(w, d)
}

func (h *Handlers) HandleScrape(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("confirm") != "yes" {
		writeErrorString(w, "must confirm")
		return
	}
	scraper, err := scraper.NewScraper(h.DB, "", false)
	if err != nil {
		writeError(w, err)
		return
	}
	if err = scraper.ScrapeCountries(); err != nil {
		writeError(w, err)
		return
	}
	if err = scraper.ScrapeLanguages(); err != nil {
		writeError(w, err)
		return
	}
	if err = scraper.ScrapeLocations(); err != nil {
		writeError(w, err)
		return
	}
	if err = scraper.ScrapeLocationAlternates(); err != nil {
		writeError(w, err)
		return
	}
	d := make(map[string]interface{})
	d["status"] = "ok"
	writeJson(w, d)
}

func (h *Handlers) HandleGeocode(w http.ResponseWriter, r *http.Request) {
	params := make(map[string]interface{})
	query := r.URL.Query()
	search := query.Get("search")
	if len(search) < 2 {
		writeErrorString(w, "search too short")
		return
	}
	params["search"] = search
	params["exact"] = true
	if _, ok := query["exact"]; ok {
		tmp := strings.ToLower(query.Get("exact"))
		params["exact"] = tmp == "1" || tmp == "yes" || tmp == "y"
	}
	if _, ok := query["lat"]; ok {
		params["lat"], _ = strconv.ParseFloat(query.Get("lat"), 64)
	}
	if _, ok := query["lng"]; ok {
		params["lng"], _ = strconv.ParseFloat(query.Get("lng"), 64)
	}
	if _, ok := query["country"]; ok {
		params["country"] = query.Get("country")
	}
	params["radius"] = 10
	if _, ok := query["radius"]; ok {
		var err error
		if params["radius"], err = strconv.ParseInt(query.Get("radius"), 10, 64); err != nil {
			writeError(w, err)
			return
		}
	}

	log.Println("geocoding country", params)

	var results []GeocodeResult
	h.DB.Raw(`
	SELECT d.*
	FROM (
		SELECT DISTINCT
			l.*,
			IF(l.feature_class='P', 1, 0) is_city,
			p.distance_unit *
				DEGREES(ACOS(COS(RADIANS(p.latitude)) *
				COS(RADIANS(l.latitude)) *
				COS(RADIANS(p.longitude - l.longitude)) + SIN(RADIANS(p.latitude)) *
				SIN(RADIANS(l.latitude))))
			AS distance,
			p.radius,
			IF(l.name LIKE p.name, 1, IF(la.alternate_name LIKE p.name, 0.5, 0)) * IF(l.feature_class='P', 10, 1) score
		FROM
			locations l
		LEFT JOIN
			location_alternates la ON la.geoname_id=l.geoname_id
		INNER JOIN (
			SELECT
				@search name,
				@country country,
				@lat latitude,
				@lng longitude,
				@radius radius,
				@exact exact,
				111.045 distance_unit
			LIMIT 1
		) AS p
		WHERE (p.country IS NULL OR p.country = l.country_code)
		AND (
			(
				l.latitude BETWEEN p.latitude  - (p.radius / p.distance_unit) AND p.latitude  + (p.radius / p.distance_unit)
				AND
				l.longitude BETWEEN p.longitude - (p.radius / (p.distance_unit * COS(RADIANS(p.latitude)))) AND p.longitude + (p.radius / (p.distance_unit * COS(RADIANS(p.latitude))))
			)
			OR p.latitude IS NULL
		)
		AND (
			((NOT p.exact AND l.name LIKE CONCAT('%',p.name,'%')) OR p.name = l.name)
			OR
			((NOT p.exact AND la.alternate_name LIKE CONCAT('%',p.name,'%')) OR p.name = la.alternate_name)
		)
	) AS d
	WHERE d.distance <= d.radius OR d.distance IS NULL
	ORDER BY d.score DESC, d.distance ASC
	LIMIT 25
	`, params).Scan(&results)

	d := make(map[string]interface{})
	d["results"] = results
	writeJson(w, d)
}
