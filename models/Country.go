package models

type Country struct {
	Iso                string   `gorm:"column:iso;size:2;unique;not null" json:"iso"`
	Iso3               *string  `gorm:"column:iso3;size:3;unique" json:"iso3"`
	IsoNumeric         *string  `gorm:"column:iso_numeric;size:6;unique" json:"iso_numeric"`
	Fips               *string  `gorm:"column:fips;size:2;index" json:"fips"`
	Name               string   `gorm:"column:name;size:128;index:,class:FULLTEXT;not null" json:"name"`
	Capital            *string  `gorm:"column:capital;size:128;index:,class:FULLTEXT" json:"capital"`
	Area               *float32 `gorm:"column:area" json:"area"`
	Population         *uint    `gorm:"column:population" json:"population"`
	Continent          *string  `gorm:"column:continent;size:32;index" json:"continent"`
	Tld                *string  `gorm:"column:tld;size:16" json:"tld"`
	CurrencyCode       *string  `gorm:"column:currency_code;size:6" json:"currency_code"`
	CurrencyName       *string  `gorm:"column:currency_name;size:32" json:"currency_name"`
	Phone              *string  `gorm:"column:phone;size:32" json:"phone"`
	PostCodeFormat     *string  `gorm:"column:post_code_format;size:64" json:"post_code_format"`
	PostCodeRegex      *string  `gorm:"column:post_code_regex;size:256" json:"post_code_regex"`
	Languages          *string  `gorm:"column:languages;size:128" json:"languages"`
	GeonameID          uint     `gorm:"column:geoname_id;unique;not null" json:"geoname_id"`
	Neighbours         *string  `gorm:"column:neighbours;size:128" json:"neighbours"`
	EquivalentFipsCode *string  `gorm:"column:equivalent_fips_code;size:32" json:"equivalent_fips_code"`
}
