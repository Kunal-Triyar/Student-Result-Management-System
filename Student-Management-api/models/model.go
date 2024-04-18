package models

import (
    "github.com/jinzhu/gorm"
)

type Student struct {
    gorm.Model
    Regno       string `gorm:"type:varchar(100);not null" json:"regno"`
    Password    string `gorm:"type:varchar(100);not null" json:"password"`
    StudentName string `gorm:"type:varchar(100);not null" json:"studentName"`
    Marks       []Mark `gorm:"foreignKey:StudentID" json:"marks"`
}

type Mark struct {
    gorm.Model
    StudentID  uint   `gorm:"not null" json:"studentId"`
    Subject    string `gorm:"type:varchar(100);not null" json:"subject"`
    CAT1       uint   `gorm:"not null" json:"CAT1"`
    CAT2       uint   `gorm:"not null" json:"CAT2"`
    FAT        uint   `gorm:"not null" json:"FAT"`
}

