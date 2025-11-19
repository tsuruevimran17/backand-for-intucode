package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	StudentId uint    `json:"student_id"`
	Student   Student `gorm:"foreignKey:StudentId"`
	Author    string  `json:"author"`
	Text      string  `json:"text"`
}

type NotePost struct {
	StudentId uint `json:"student_id"`
	Author string `json:"author"`
	Text   string `json:"text"`
}

type NotePatch struct {
	StudentId *uint `json:"student_id"`
	Author *string `json:"author"`
	Text   *string `json:"text"`
}
