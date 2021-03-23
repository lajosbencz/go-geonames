package models

type LocationAlternate struct {
	ID              uint    `gorm:"column:id;unique;not null" json:"id"`
	GeonameID       uint    `gorm:"column:geoname_id;index;not null" json:"geoname_id"`
	IsoLanguage     *string `gorm:"column:iso_language;size:7" json:"iso_language"`
	AlternateName   string  `gorm:"column:alternate_name;size:255;not null" json:"alternate_name"`
	IsPreferredName bool    `gorm:"column:is_preferred_name;index;not null" json:"is_preferred_name"`
	IsShortName     bool    `gorm:"column:is_short_name;index;not null" json:"is_short_name"`
	IsColloquial    bool    `gorm:"column:is_colloquial;index;not null" json:"is_colloquial"`
	IsHistoric      bool    `gorm:"column:is_historic;index;not null" json:"is_historic"`
	UsedFrom        *string `gorm:"column:used_from;size:64" json:"used_from"`
	UsedTo          *string `gorm:"column:used_to;size:64" json:"used_to"`
}
