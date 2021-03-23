package models

type Language struct {
	Iso3 string `gorm:"column:iso3;size:3;unique" json:"iso3"`
	Iso2 string `gorm:"column:iso2;size:32;unique" json:"iso2"`
	Iso1 string `gorm:"column:iso1;size:2;unique" json:"iso1"`
	Name string `gorm:"column:name;size:64;not null" json:"name"`
}
