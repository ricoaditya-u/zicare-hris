package models

import (
	"gorm.io/gorm"
)

type Mpp struct {
	gorm.Model
	Period       string     `json:"period"`
	DepartmentID int        `json:"departmentid"`
	Department   Department `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Numberreq    int        `json:"numberreq"`
	Budget       int        `json:"budget"`
	Status       int        `json:"status"`
}

// type DetailMpp struct {
// 	gorm.Model
// 	MppID          int         `json:"mppid"`
// 	Mpp            Mpp         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
// 	JobpositionId  int         `json:"jobposition"`
// 	Jobposition    Jobposition `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
// 	Statusemployee string      `json:"statusemployee"`
// }
