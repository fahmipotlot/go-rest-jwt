package model

import (
	"diary_go_api/database"

	"gorm.io/gorm"
)

type Entry struct {
	gorm.Model
	Content string `gorm:"type:text" json:"content"`
	UserID  uint
}

func (entry *Entry) Save() (*Entry, error) {
	err := database.Database.Create(&entry).Error
	if err != nil {
		return &Entry{}, err
	}
	return entry, nil
}

func FindEntryById(id string, userID uint) (Entry, error) {
	var entry Entry
	err := database.Database.Where("id=?", id).Where("user_id=?", userID).Find(&entry).Error
	if err != nil {
		return Entry{}, err
	}
	return entry, nil
}

func DeleteEntryById(id string, userID uint) (Entry, error) {
	var entry Entry
	err := database.Database.Where("id=?", id).Where("user_id=?", userID).Delete(&entry).Error
	if err != nil {
		return Entry{}, err
	}
	return entry, nil
}

func (entry *Entry) Update() (*Entry, error) {
	err := database.Database.Save(&entry).Error
	if err != nil {
		return &Entry{}, err
	}
	return entry, nil
}
