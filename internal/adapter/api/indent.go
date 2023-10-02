package api

import (
	"bnt/bnt-indent-service/internal/core/domain/dto"
	errorhelper "bnt/bnt-indent-service/internal/core/helper/error-helper"
	validation "bnt/bnt-indent-service/internal/core/helper/validation-helper"

	"github.com/gin-gonic/gin"
)

func (hdl *HTTPHandler) GetIndentByRef(c *gin.Context) {
	indent, err := hdl.indentService.GetIndentByRef(c.Param("reference"))

	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}

	c.JSON(200, indent)
}

func (hdl *HTTPHandler) GetIndentList(c *gin.Context) {
	indents, err := hdl.indentService.GetIndentList(c.Param("page"))

	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}

	c.JSON(200, indents)
}
func (hdl *HTTPHandler) CreateIndent(c *gin.Context) {
	body := dto.CreateIndent{}
	_ = c.BindJSON(&body)

	if err := validation.Validate(&body); err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}

	indent, err := hdl.indentService.CreateIndent(body)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}
	c.JSON(201, gin.H{"reference": indent})
}

func (hdl *HTTPHandler) UpdateIndent(c *gin.Context) {
	body := dto.UpdateIndent{}
	_ = c.BindJSON(&body)
	if err := validation.Validate(&body); err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}
	indent, err := hdl.indentService.UpdateIndent(c.Param("reference"), body)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return

	}
	c.JSON(200, gin.H{"reference": indent})
}
func (hdl *HTTPHandler) DisableIndent(c *gin.Context) {

	indent, err := hdl.indentService.DisableIndent(c.Param("reference"))
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.UpdateError, err.Error()))
		return
	}
	c.JSON(200, gin.H{"reference": indent})
}
func (hdl *HTTPHandler) EnableIndent(c *gin.Context) {
	indent, err := hdl.indentService.EnableIndent(c.Param("reference"))
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.UpdateError, err.Error()))
		return
	}
	c.JSON(200, gin.H{"reference": indent})
}
func (hdl *HTTPHandler) DeleteIndent(c *gin.Context) {
	indent, err := hdl.indentService.DeleteIndent(c.Param("reference"))
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.UpdateError, err.Error()))
		return
	}
	c.JSON(200, gin.H{"reference": indent})
}
func (hdl *HTTPHandler) GetYearList(c *gin.Context) {
	years, err := hdl.indentService.GetYearList(c.Param("count"))
	if err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}
	c.JSON(200, years)
}
func (hdl *HTTPHandler) GetDenominationList(c *gin.Context) {
	denominations, err := hdl.indentService.GetDenominationList()
	if err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}
	c.JSON(200, denominations)
}
func (hdl *HTTPHandler) GetLocationList(c *gin.Context) {
	locations, err := hdl.indentService.GetLocationList()
	if err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}
	c.JSON(200, locations)
}

func (hdl *HTTPHandler) GetIndentByDenomination(c *gin.Context) {
	indents, err := hdl.indentService.GetIndentByDenomination(c.Param("denomination"), c.Param("page"))

	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}

	c.JSON(200, indents)
}
func (hdl *HTTPHandler) GetIndentByYear(c *gin.Context) {
	indents, err := hdl.indentService.GetIndentByYear(c.Param("year"), c.Param("page"))

	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}

	c.JSON(200, indents)
}
