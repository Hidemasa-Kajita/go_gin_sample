package controller

import (
	"net/http"

	"github.com/Hidemasa-Kajita/go_api_sample/entity"
	"github.com/Hidemasa-Kajita/go_api_sample/repository"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	us := repository.FindAll()
	c.JSON(http.StatusOK, gin.H{
		"users": us,
	})
}

func Show(c *gin.Context) {
	id := c.Param("id")

	u := repository.FindOneById(id)

	c.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}

type UserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

func Store(c *gin.Context) {
	var r UserRequest
	c.ShouldBindJSON(&r)

	eu := entity.User{Name: r.Name, Age: r.Age, Sex: r.Sex}

	u := repository.Insert(eu)

	c.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}

func Update(c *gin.Context) {
	id := c.Param("id")

	var r UserRequest
	c.ShouldBindJSON(&r)

	eu := entity.User{Name: r.Name, Age: r.Age, Sex: r.Sex}

	u := repository.Update(id, eu)

	c.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}

func Delete(c *gin.Context) {
	id := c.Param("id")

	repository.Delete(id)

	c.JSON(http.StatusNoContent, gin.H{})
}
