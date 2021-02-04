package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pclubiitk/puppy-love/models"
	"gopkg.in/mgo.v2/bson"
)

// GetStats returns useful statistics
func GetStats(c *gin.Context) {
	var users []models.User
	var hearts []models.Heart
	if err := Db.GetCollection("user").Find(bson.M{"dirty": false}).All(&users); err != nil {
		c.String(http.StatusInternalServerError, "Could not get database info")
		return
	}
	if err := Db.GetCollection("heart").Find(nil).All(&hearts); err != nil {
		c.String(http.StatusInternalServerError, "Could not get database info")
		return
	}

	var y20males, y19males, y18males, y17males, othermales int
	var y20females, y19females, y18females, y17females, otherfemales int

	for _, user := range users {
		if user.Gender == "1" {
			if strings.HasPrefix(user.Id, "20") {
				y20males++
			} else if strings.HasPrefix(user.Id, "19") {
				y19males++
			} else if strings.HasPrefix(user.Id, "18") {
				y18males++
			} else if strings.HasPrefix(user.Id, "17") {
				y17males++
			} else {
				othermales++
			}
		} else {
			if strings.HasPrefix(user.Id, "20") {
				y20females++
			} else if strings.HasPrefix(user.Id, "19") {
				y19females++
			} else if strings.HasPrefix(user.Id, "18") {
				y18females++
			} else if strings.HasPrefix(user.Id, "17") {
				y17females++
			} else {
				otherfemales++
			}
		}
	}

	var y20maleHearts, y19maleHearts, y18maleHearts, y17maleHearts, othermaleHearts int
	var y20femaleHearts, y19femaleHearts, y18femaleHearts, y17femaleHearts, otherfemaleHearts int

	for _, heart := range hearts {
		if heart.Gender == "1" {
			if strings.HasPrefix(heart.Id, "20") {
				y20maleHearts++
			} else if strings.HasPrefix(heart.Id, "19") {
				y19maleHearts++
			} else if strings.HasPrefix(heart.Id, "18") {
				y18maleHearts++
			} else if strings.HasPrefix(heart.Id, "17") {
				y17maleHearts++
			} else {
				othermaleHearts++
			}
		} else {
			if strings.HasPrefix(heart.Id, "20") {
				y20femaleHearts++
			} else if strings.HasPrefix(heart.Id, "19") {
				y19femaleHearts++
			} else if strings.HasPrefix(heart.Id, "18") {
				y18femaleHearts++
			} else if strings.HasPrefix(heart.Id, "17") {
				y17femaleHearts++
			} else {
				otherfemaleHearts++
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"users":             len(users),
		"y20males":          y20males,
		"y19males":          y19males,
		"y18males":          y18males,
		"y17males":          y17males,
		"othermales":        othermales,
		"y20females":        y20females,
		"y19females":        y19females,
		"y18females":        y18females,
		"y17females":        y17females,
		"otherfemales":      otherfemales,
		"y20maleHearts":     y20maleHearts,
		"y19maleHearts":     y19maleHearts,
		"y18maleHearts":     y18maleHearts,
		"y17maleHearts":     y17maleHearts,
		"othermaleHearts":   othermaleHearts,
		"y20femaleHearts":   y20femaleHearts,
		"y19femaleHearts":   y19femaleHearts,
		"y18femaleHearts":   y18femaleHearts,
		"y17femaleHearts":   y17femaleHearts,
		"otherfemaleHearts": otherfemaleHearts,
	})
}
