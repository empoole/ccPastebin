package model

import (
	"ccPasteBinServer/database"

	"crypto/md5"
	"encoding/hex"

	"gorm.io/gorm"
)

type Note struct {
	gorm.Model

	ID int `json:"id,primary_key"`
	Message string `gorm:"not null" json:"message"`
	Hash string `gorm:"size:255;not null" json:"hash"`
}

func (note *Note) Save() (*Note, error) {
	err := database.DBConnection.Create(&note).Error
	if(err != nil) {
		return &Note{}, err
	}

	return note, nil
}

func (note *Note) BeforeSave(*gorm.DB) error {
	note.Hash = GetMD5Hash(note.Message)
	return nil
}

func FindNoteByHash(hash string) (Note, error) {
	var note Note
	err := database.DBConnection.Where("hash=?", hash).Find(&note).Error
	if err != nil {
		return Note{}, err
	}
	
	return note, nil
}

func FindNoteById(id uint) (Note, error) {
	var note Note
	err := database.DBConnection.Preload("Entries").Where("ID=?", id).Find(&note).Error
	if err != nil {
			return Note{}, err
	}
	return note, nil
}

func GetMD5Hash(text string) string {
 hash := md5.Sum([]byte(text))
 return hex.EncodeToString(hash[:])
}