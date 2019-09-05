package controllers

import (
	"github.com/gin-gonic/gin"
)

func ControlOne(c *gin.Context) {
	c.JSON(200, gin.H{
		"Case" : "One",
		"Status" : "SUCCESS",
	})
}

func ControlTwo(c *gin.Context) {
	name := c.PostForm("name")
	c.JSON(200, gin.H{
		"Case" : "Two",
		"nextId" : name,
		"Status" : "SUCCESS",
	})
}

func ControlThree(c *gin.Context) {
	name := c.PostForm("name")
	nextId := c.PostForm("nextId")
	if name == nextId {
		c.JSON(200, gin.H{
			"Case" : "Three",
			"Status" : "SUCCESS",
		})
	} else {
		c.JSON(400, gin.H{
			"Message" : "Invalid nextId",
		})
	}
}
