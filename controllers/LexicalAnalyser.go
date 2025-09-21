package controllers

import "github.com/gin-gonic/gin"

func LexicalAnalyser(c *gin.Context) {
	c.JSON(200,gin.H{
		"lexixal":"analyser",
	})

}
