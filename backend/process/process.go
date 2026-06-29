package process

import (
	"github.com/1234arushi/JD-Analyzer/database"
	"github.com/1234arushi/JD-Analyzer/database/models"
	"github.com/1234arushi/JD-Analyzer/services"
	"github.com/gin-gonic/gin"
)

func (req *JDRequest) Process(c *gin.Context) (err error) {
	cleanedJD, err := services.CleanJobDescription(req.JobDescription)
	if err != nil {
		return
	}
	job := models.Job{
		Company:  req.Company,
		Role:     req.Role,
		Location: req.Location,
		RawJD:    cleanedJD,
	}
	err = database.DB.Create(&job).Error
	if err != nil {
		return err
	}

	return

}
