package models

import (
	"gorm.io/gorm"
)

type Reqheadcount struct {
	gorm.Model
	MppID            int            `json:"mppid"`
	Mpp              []Mpp          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	RequestedBy      int64          `json:"requestedBy"`
	Employee         Employee       `gorm:"foreignKey:RequestedBy"`
	Structure        string         `json:"structure"`
	GradeID          int            `json:"gradeid"`
	Grade            Grade          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Statusemployee   string         `json:"statusemployee"`
	Reasonhiring     string         `json:"reasonhiring"`
	Degree           string         `json:"degree"`
	Minexp           string         `json:"minexp"`
	JobDescriptionID int            `json:"jobdescriptionid"`
	JobDescription   JobDescription `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Specification    string         `json:"specification"`
	Gender           string         `json:"gender"`
	Age              string         `json:"age"`
	Maritalstatus    string         `json:"maritalstatus"`
	Recruitmenttype  string         `json:"recruitmenttype"`
	Status           string         `json:"status"`
}
