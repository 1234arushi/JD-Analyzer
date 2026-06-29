package process

import (
	"github.com/gin-gonic/gin"
)

// Go only allows methods to be defined in the same package as their receiver type.
func (req *JDRequest) Validate(c *gin.Context) (err error) {
	if err = c.ShouldBindJSON(&req); err != nil {
		return err
	}
	return

}
