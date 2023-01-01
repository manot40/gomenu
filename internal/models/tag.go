package models

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name string `json:"name" gorm:"type:varchar(100);not null;unique"`
}

func (t *Tag) Delete() error {
	return DB.Delete(t).Error
}

func (t *Tag) Create() error {
	return DB.Create(t).Error
}
