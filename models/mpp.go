package models

import (
	"gorm.io/gorm"
)

type Mpp struct {
	gorm.Model
	EmployeeID    int64          `json:"employeeid"`
	Employee      Employee       `gorm:"references:EmployeeId"`
	Year          string         `json:"year"`
	Month         string         `json:"month"`
	DepartmentID  int            `json:"departmentid"`
	Department    Department     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Numberreq     int            `json:"numberreq"`
	Budget        int            `json:"budget"`
	Status        int            `json:"status"`
	Reqheadcounts []Reqheadcount `gorm:"foreignKey:MppID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Reqheadcount struct {
	gorm.Model
	MppID            int            `json:"mppid"`
	EmployeeID       int64          `json:"employeeid"`
	Employee         Employee       `gorm:"references:EmployeeId"`
	LevelID          int            `json:"levelid"`
	Level            Level          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
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
