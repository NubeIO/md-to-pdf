package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (app *application) routes() http.Handler {
	r := gin.Default()
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, errorJson("Page not found"))
	})
	r.POST("/convert", func(c *gin.Context) {
		var json struct {
			Input          []byte `json:"input" binding:"required"`
			WriteToHomeDir bool   `json:"write_to_home_dir"`
		}

		if err := c.ShouldBindJSON(&json); err != nil {
			logger.Errorf("[CONVERT]: %v", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, errorJson("invalid request"))
			return
		}

		bin, err := app.convert(c.Request.Context(), json.Input, json.WriteToHomeDir)
		if err != nil {
			logger.Errorf("[CONVERT]: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorJson("error converting markdown"))
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"content": bin,
		})
	})

	r.POST("/convert/local", func(c *gin.Context) {
		var json struct {
			File           string `json:"file"`
			WriteToHomeDir bool   `json:"write_to_home_dir"`
		}
		if err := c.ShouldBindJSON(&json); err != nil {
			logger.Errorf("[CONVERT]: %v", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, errorJson("invalid request"))
			return
		}
		bin, err := app.convertFromRead(c.Request.Context(), json.File)
		if err != nil {
			logger.Errorf("[CONVERT]: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorJson("error converting markdown"))
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"content": bin,
		})
	})

	return r
}
