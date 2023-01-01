package models

import (
	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	Name        string  `json:"name" gorm:"type:varchar(100);not null;unique"`
	Description *string `json:"description" gorm:"type:varchar(255)"`
	Price       uint    `json:"price" gorm:"type:int;not null"`
	Tags        string  `json:"tags" gorm:"type:varchar(100)"`
	Like        uint    `json:"like" gorm:"type:int;default:0"`
}

func (m *Menu) Fav() error {
	m.Like++
	return DB.Save(m).Error
}

func (m *Menu) UnFav() error {
	m.Like--
	return DB.Save(m).Error
}

func (m *Menu) Update(input *Menu) error {
	m.Name = input.Name
	m.Description = input.Description
	m.Price = input.Price
	m.Tags = input.Tags
	return DB.Save(m).Error
}

func (m *Menu) Delete() error {
	return DB.Delete(m).Error
}

func (m *Menu) Create() error {
	return DB.Create(m).Error
}
