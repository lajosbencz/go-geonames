package scraper

const URLLocations = "http://download.geonames.org/export/dump/allCountries.zip"

const SQLLocations = `
LOAD DATA LOCAL INFILE '%[2]s'
INTO TABLE %[1]s
FIELDS TERMINATED BY '\t'
OPTIONALLY ENCLOSED BY ''
LINES TERMINATED BY '\n'
IGNORE 0 LINES
(
    geoname_id,
    name,
    name_ascii,
    @name_alternates,
    latitude,
    longitude,
    feature_class,
    feature_code,
    @country_code,
    @cc2,
    @admin1,
    @admin2,
    @admin3,
    @admin4,
    @population,
    @elevation,
    @dem,
    @timezone,
    @date_change)
SET
name_alternates = IF(@name_alternates = '', NULL, @name_alternates),
country_code = IF(@country_code = '', NULL, @country_code),
cc2 = IF(@cc2 = '', NULL, @cc2),
admin1 = IF(@admin1 = '', NULL, @admin1),
admin2 = IF(@admin2 = '', NULL, @admin2),
admin3 = IF(@admin3 = '', NULL, @admin3),
admin4 = IF(@admin4 = '', NULL, @admin4),
population = IF(@population = '', NULL, @population),
elevation = IF(@elevation = '', NULL, @elevation),
dem = IF(@dem = '', NULL, @dem),
timezone = IF(@timezone = '', NULL, @timezone),
date_change = IF(@date_change = '', NULL, @date_change)
;`
