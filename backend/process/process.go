package process

import (
	"github.com/1234arushi/JD-Analyzer/database"
	"github.com/1234arushi/JD-Analyzer/database/models"
	"github.com/1234arushi/JD-Analyzer/services"
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

func (req *JDRequest) Process(c *gin.Context) (err error) {
	var existingJob models.Job
	cleanedJD, err := services.CleanJobDescription(req.JobDescription) //cleanedJSD -> still a string
	if err != nil {
		return
	}
	job := models.Job{
		Company:  req.Company,
		Role:     req.Role,
		Location: req.Location,
		RawJD:    datatypes.JSON([]byte(cleanedJD)), //string to bytes as JSON internally is stored as bytes,
		//.JSON takes that byte and convert it to JSON
	}

	//checking if I have stored that company
	err = database.DB.Where("company = ? AND role = ?", req.Company, req.Role).First(&existingJob).Error
	if err == nil {
		return nil //record already exists,don't insert again
	}

	err = database.DB.Create(&job).Error
	if err != nil {
		return err
	}

	return

}
