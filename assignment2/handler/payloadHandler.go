package handler

import (
	"assignment2/models"
	"regexp"

	"github.com/gin-gonic/gin"
)

func PayloadCheck(c *gin.Context) {
	var body models.Payload

	if err := c.Bind(&body); err != nil {
		c.JSON(401, gin.H{
			"Message": "data is not found please check the fields",
			"error":   err.Error(),
		})
		c.Abort()
	}
	regex := regexp.MustCompile(`\w+`)
	word := regex.FindAllString(body.Str, -1)
	if len(word) >= 8 {
		c.JSON(200, gin.H{
			"message": "sucess the word is atleast 8 words you can login now",
			"data":    body.Str,
		})
		c.Abort()
	} else {
		c.JSON(411, gin.H{
			"message": "content length is less than 8 please enter the content more than 8 words",
		})
		c.Abort()
	}

}
