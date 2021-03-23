package scraper

const URLLanguages = "http://download.geonames.org/export/dump/iso-languagecodes.txt"

const SQLLanguages = `
LOAD DATA LOCAL INFILE '%[2]s'
INTO TABLE %[1]s
FIELDS TERMINATED BY '\t'
OPTIONALLY ENCLOSED BY ''
LINES TERMINATED BY '\n'
IGNORE 1 LINES
(
    @iso3,
    @iso2,
    @iso1,
    name
)
SET
iso3 = IF(@iso3 = '', NULL, @iso3),
iso2 = IF(@iso2 = '', NULL, @iso2),
iso1 = IF(@iso1 = '', NULL, @iso1)
;`
