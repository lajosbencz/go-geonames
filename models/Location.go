package models

import "time"

type Location struct {
	GeonameID      uint       `gorm:"column:geoname_id;unique;not null" json:"geoname_id"`
	Name           string     `gorm:"column:name;size:256;index:,class:FULLTEXT;not null" json:"name"`
	NameAscii      string     `gorm:"column:name_ascii;size:256;index:,class:FULLTEXT;not null" json:"name_ascii"`
	NameAlternates List       `gorm:"column:name_alternates;index:,class:FULLTEXT;type:text" json:"name_alternates"`
	Latitude       float64    `gorm:"column:latitude;not null" json:"latitude"`
	Longitude      float64    `gorm:"column:longitude;not null" json:"longitude"`
	FeatureClass   string     `gorm:"column:feature_class;size:1;index;not null" json:"feature_class"`
	FeatureCode    string     `gorm:"column:feature_code;size:16;index;not null" json:"feature_code"`
	CountryCode    *string    `gorm:"column:country_code;size:2;index" json:"country_code"`
	CC2            *string    `gorm:"column:cc2;size:200" json:"cc2"`
	Admin1         *string    `gorm:"column:admin1;size:20" json:"admin1"`
	Admin2         *string    `gorm:"column:admin2;size:80" json:"admin2"`
	Admin3         *string    `gorm:"column:admin3;size:20" json:"admin3"`
	Admin4         *string    `gorm:"column:admin4;size:20" json:"admin4"`
	Population     *uint64    `gorm:"column:population" json:"population"`
	Elevation      *int64     `gorm:"column:elevation" json:"elevation"`
	Dem            *int64     `gorm:"column:dem" json:"dem"`
	Timezone       *string    `gorm:"column:timezone;size:40" json:"timezone"`
	DateChange     *time.Time `gorm:"column:date_change;type:date" json:"date_change"`
}
