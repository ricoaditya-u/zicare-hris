package controllers

// import (
// 	"errors"
// 	"fmt"
// 	"log"
// 	"math"
// 	"net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/ricoaditya-u/hris-master/db"
// 	"github.com/ricoaditya-u/hris-master/models"
// 	"gorm.io/gorm"
// )

// func CalculateKIS(salary int) int {
// 	return salary * 1 / 100
// }

// func CalculateKPJ(salary int) int {
// 	return salary * 3 / 100
// }

// func CalculatePPH(salary int, ptkp int) int {
// 	jabatan := math.Min(500000, float64(salary*5/100))
// 	kalkulasi := (salary - int(jabatan) - CalculateKIS(salary) - CalculateKPJ(salary) - ptkp) * 12

// 	var pph models.PPH
// 	db.DB.Select("Percentage").Where("value >= ?", kalkulasi).First(&pph)

// 	if kalkulasi > 0 {
// 		return int(kalkulasi) * int(pph.Percentage) / 100 / 12
// 	} else {
// 		return 0
// 	}
// }

// func CreateSalarySlipHeader(NikID int, Salary int) error {
// 	//INSERT HEADER
// 	salarySlip := models.SalarySlip{
// 		NikID:  int64(NikID),
// 		Period: time.Now().Format("2006-01"),
// 		Salary: int64(Salary),
// 		Status: 0,
// 	}
// 	result := db.DB.Create(&salarySlip)
// 	if result.Error != nil {
// 		return result.Error
// 	}
// 	return nil
// 	//END INSERT HEADER
// }
// func CreateSalarySlipDetail(NikID int, Salary int, PTKP int) error {

// 	var Header models.SalarySlip
// 	db.DB.Where("nik_id = ? AND period = ?", NikID, time.Now().Format("2006-01")).First(&Header)
// 	//INSERT DETAIL Deduction
// 	//KIS
// 	dataKIS := models.SalarySlipDetail{
// 		SalarySlipID: Header.ID,
// 		Type:         2,
// 		Name:         "BPJS Kesehatan",
// 		Value:        int64(CalculateKIS(int(Salary))),
// 	}
// 	result := db.DB.Create(&dataKIS)
// 	if result.Error != nil {
// 		return result.Error
// 	}

// 	//KPJ
// 	dataKPJ := models.SalarySlipDetail{
// 		SalarySlipID: Header.ID,
// 		Type:         2,
// 		Name:         "BPJS Ketenagakerjaan",
// 		Value:        int64(CalculateKPJ(int(Salary))),
// 	}

// 	result = db.DB.Create(&dataKPJ)
// 	if result.Error != nil {
// 		return result.Error
// 	}

// 	//PPH
// 	dataPPH := models.SalarySlipDetail{
// 		SalarySlipID: Header.ID,
// 		Type:         2,
// 		Name:         "PPH 21",
// 		Value:        int64(CalculatePPH(int(Salary), int(PTKP))),
// 	}

// 	result = db.DB.Create(&dataPPH)
// 	if result.Error != nil {
// 		return result.Error
// 	}

// 	return nil
// 	//INSERT DETAIL Additional
// }

// // Generate Slip Batch
// func GenerateSlip(c *gin.Context) {

// 	var result []models.Employee

// 	if err := db.DB.Joins("JOIN ptkps ON employees.ptkp_id = ptkps.id").Find(&result).Error; err != nil {
// 		log.Printf("Error querying employees: %s\n", err)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	for _, data := range result {
// 		//INSERT HEADER
// 		if err := CreateSalarySlipHeader(int(data.Nik), int(data.Salary)); err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create SalarySlip Header " + fmt.Sprint(data.Nik) + err.Error()})
// 			continue
// 		}
// 		//END INSERT HEADER

// 		//INSERT DETAIL
// 		if err := CreateSalarySlipDetail(int(data.Nik), int(data.Salary), int(data.Ptkp.Value)); err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create SalarySlip Detail " + fmt.Sprint(data.Nik) + err.Error()})
// 			continue
// 		}

// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "Salary Slip generated successfully"})

// }

// // Show List Period
// func SalarySlipShow(c *gin.Context) {
// 	period := c.Param("period")

// 	var SalarySlip []models.SalarySlip
// 	err := db.DB.Find(&SalarySlip, "period = ?", period).Error
// 	if err != nil {
// 		errors.Is(err, gorm.ErrRecordNotFound)
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"error": "record not found",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"data": SalarySlip,
// 	})
// }

// // Show Detail
// func SalarySlipDetailShow(c *gin.Context) {
// 	period := c.Param("period")
// 	nikID := c.Param("id")

// 	var SalarySlip []models.SalarySlip
// 	err := db.DB.Where("period = ? AND nik_id = ?", period, nikID).Preload("SalarySlipDetails").Find(&SalarySlip).Error

// 	if err != nil {
// 		errors.Is(err, gorm.ErrRecordNotFound)
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"error": "record not found",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"data": SalarySlip,
// 	})
// }

// // Create Header
// func SalarySlipCreate(c *gin.Context) {
// 	// Get data req
// 	var body models.SalarySlip

// 	if err := c.ShouldBindJSON(&body); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	var result models.Employee

// 	if err := db.DB.Where("nik = ?", body.ID).First(&result).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"error": "record not found",
// 		})
// 		return
// 	}

// 	//INSERT HEADER
// 	if err := CreateSalarySlipHeader(int(result.Nik), int(result.Salary)); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create SalarySlip Header " + fmt.Sprint(result.Nik) + err.Error()})
// 		return
// 	} else {
// 		c.JSON(http.StatusOK, gin.H{"message": "Salary Slip generated successfully"})
// 	}
// 	//END INSERT HEADER
// }

// func SalarySlipDetailCreate(c *gin.Context) {
// 	// Get data body
// 	var body models.SalarySlipDetail

// 	if err := c.ShouldBindJSON(&body); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Validasi bahwa name dan value tidak boleh kosong
// 	if body.Name == "" || body.Value <= 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "name, type, and value must not be empty"})
// 		return
// 	}

// 	if body.Type != 1 && body.Type != 2 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong Type Value"})
// 		return
// 	}

// 	var result models.SalarySlip

// 	if err := db.DB.Where("nik_id = ? AND period = ? AND status = 0", body.ID, time.Now().Format("2006-01")).First(&result).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"error": "record not found",
// 		})
// 		return
// 	}

// 	data := models.SalarySlipDetail{
// 		SalarySlipID: result.ID,
// 		Name:         body.Name,
// 		Type:         body.Type,
// 		Value:        body.Value,
// 	}

// 	err := db.DB.Create(&data)
// 	if err.Error != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error,
// 		})
// 		return
// 	}

// 	// Respond
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "create slip detail success",
// 	})
// }

// // UPDATE SALARY SLIP HEADER
// func SalarySlipUpdate(c *gin.Context) {
// 	// Get id
// 	id := c.Param("id")

// 	// Get data body
// 	var body models.SalarySlip

// 	if err := c.ShouldBindJSON(&body); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Find the data
// 	var salarySlip models.SalarySlip
// 	err := db.DB.Where("status = 0").First(&salarySlip, "id = ?", id).Error

// 	if err != nil {
// 		errors.Is(err, gorm.ErrRecordNotFound)
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"error": "record not found",
// 		})
// 		return
// 	}

// 	// Update
// 	db.DB.Model(&salarySlip).Updates(models.SalarySlip{
// 		Salary: body.Salary,
// 	})

// 	// Respond
// 	c.JSON(http.StatusOK, gin.H{
// 		"data": salarySlip,
// 	})
// }

// // UPDATE SALARY SLIP DETAIL
// func SalarySlipDetailUpdate(c *gin.Context) {
// 	// Get id
// 	id := c.Param("id")

// 	// Get data body
// 	var body models.SalarySlipDetail

// 	if err := c.ShouldBindJSON(&body); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	// Validasi bahwa name dan value tidak boleh kosong
// 	if body.Name == "" || body.Value <= 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "name, type, and value must not be empty"})
// 		return
// 	}

// 	if body.Type != 1 && body.Type != 2 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong Type Value"})
// 		return
// 	}
// 	// Find the data
// 	var salarySlipDetail models.SalarySlipDetail
// 	err := db.DB.First(&salarySlipDetail, "id = ?", id).Error

// 	if err != nil {
// 		errors.Is(err, gorm.ErrRecordNotFound)
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"error": "record not found",
// 		})
// 		return
// 	}

// 	// Update
// 	db.DB.Model(&salarySlipDetail).Updates(models.SalarySlipDetail{
// 		Name:  body.Name,
// 		Type:  body.Type,
// 		Value: body.Value,
// 	})

// 	// Respond
// 	c.JSON(http.StatusOK, gin.H{
// 		"data": salarySlipDetail,
// 	})
// }

// // Delete Header
// func SalarySlipDelete(c *gin.Context) {
// 	id := c.Param("id")

// 	// Delete Header
// 	var salarySlip models.SalarySlip
// 	if err := db.DB.Unscoped().Delete(&salarySlip, id).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"error": "record not found",
// 		})
// 		return
// 	}

// 	// Respond
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "delete success",
// 	})
// }

// // Delete Detail
// func SalarySlipDetailDelete(c *gin.Context) {
// 	id := c.Param("id")

// 	// Delete Header
// 	var SalarySlipDetail models.SalarySlipDetail
// 	if err := db.DB.Unscoped().Delete(&SalarySlipDetail, id).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"error": "record not found",
// 		})
// 		return
// 	}

// 	// Respond
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "delete success",
// 	})
// }
