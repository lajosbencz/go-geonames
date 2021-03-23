package scraper

const URLAlternates = "http://download.geonames.org/export/dump/alternateNamesV2.zip"

const SQLLocationAlternates = `
LOAD DATA LOCAL INFILE '%[2]s'
INTO TABLE %[1]s
FIELDS TERMINATED BY '\t'
OPTIONALLY ENCLOSED BY ''
LINES TERMINATED BY '\n'
IGNORE 0 LINES
(
    id,
    geoname_id,
    @iso_language,
    alternate_name,
    @is_preferred_name,
    @is_short_name,
    @is_colloquial,
    @is_historic,
    @used_from,
    @used_to)
SET
iso_language = IF(@iso_language = '', NULL, @iso_language),
is_preferred_name = IF(@is_preferred_name IN ('', '0'), b'0', b'1'),
is_short_name = IF(@is_short_name IN ('', '0'), b'0', b'1'),
is_colloquial = IF(@is_colloquial IN ('', '0'), b'0', b'1'),
is_historic = IF(@is_historic IN ('', '0'), b'0', b'1'),
used_from = IF(@used_from = '', NULL, @used_from),
used_to = IF(@used_to = '', NULL, @used_to)
;`
