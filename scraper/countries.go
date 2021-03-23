package scraper

const URLCountries = "http://download.geonames.org/export/dump/countryInfo.txt"

const SQLCountries = `
LOAD DATA LOCAL INFILE '%[2]s'
INTO TABLE %[1]s
FIELDS TERMINATED BY '\t'
OPTIONALLY ENCLOSED BY ''
LINES TERMINATED BY '\n'
IGNORE 0 LINES
(
    iso,
    @iso3,
    @iso_numeric,
    @fips,
    name,
    @capital,
    @area,
    @population,
    @continent,
    @tld,
    @currency_code,
    @currency_name,
    @phone,
    @post_code_format,
    @post_code_regex,
    @languages,
    geoname_id,
    @neighbours,
    @equivalent_fips_code
)
SET
iso3 = IF(@iso3 = '', NULL, @iso3),
iso_numeric = IF(@iso_numeric = '', NULL, @iso_numeric),
fips = IF(@fips = '', NULL, @fips),
capital = IF(@capital = '', NULL, @capital),
area = IF(@area = '', NULL, @area),
population = IF(@population = '', NULL, @population),
continent = IF(@continent = '', NULL, @continent),
tld = IF(@tld = '', NULL, @tld),
currency_code = IF(@currency_code = '', NULL, @currency_code),
currency_name = IF(@currency_name = '', NULL, @currency_name),
phone = IF(@phone = '', NULL, @phone),
post_code_format = IF(@post_code_format = '', NULL, @post_code_format),
post_code_regex = IF(@post_code_regex = '', NULL, @post_code_regex),
languages = IF(@languages = '', NULL, @languages),
neighbours = IF(@neighbours = '', NULL, @neighbours),
equivalent_fips_code = IF(@equivalent_fips_code = '', NULL, @equivalent_fips_code)
`
