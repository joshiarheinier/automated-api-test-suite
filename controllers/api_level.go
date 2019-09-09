package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func ControlOne(c *gin.Context) {
	c.JSON(200, gin.H{
		"case" : "One",
		"status" : "SUCCESS",
	})
}

func ControlTwo(c *gin.Context) {
	tmp := make(map[string]string)
	body, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(body, &tmp)
	c.JSON(200, gin.H{
		"case" : "Two",
		"nextId" : tmp["name"],
		"status" : 1,
	})
}

func ControlThree(c *gin.Context) {
	tmp := make(map[string]string)
	body, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(body, &tmp)
	if tmp["name"] == tmp["nextId"] {
		c.JSON(200, gin.H{
			"case" : "Three",
			"status" : "SUCCESS",
		})
	} else {
		c.JSON(400, gin.H{
			"message" : "Invalid nextId",
		})
	}
}
