package models

import (
	"gorm.io/gorm"
)

type SalarySlip struct {
	gorm.Model
	NikID             int64              `json:"nikid" gorm:"uniqueIndex:unique_employee_period"`
	Nik               Employee           `gorm:"references:Nik"`
	Period            string             `json:"period" gorm:"uniqueIndex:unique_employee_period"`
	Salary            int64              `json:"salary" gorm:"not null"`
	Status            int32              `json:"Status"`
	SalarySlipDetails []SalarySlipDetail `gorm:"foreignKey:SalarySlipID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type SalarySlipDetail struct {
	gorm.Model
	SalarySlipID uint   `json:"salaryslipid"`
	Type         int32  `json:"type" gorm:"not null"`
	Name         string `json:"name" gorm:"not null"`
	Value        int64  `json:"Value" gorm:"not null"`
}

type PTKP struct {
	gorm.Model
	Name  string `json:"name" gorm:"unique"`
	Value int64  `json:"Value"`
}

type PPH struct {
	gorm.Model
	Value      int64 `json:"value"`
	Percentage int32 `json:"percentage"`
}
