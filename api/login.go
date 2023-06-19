package api

import (
	"api-db/datasource"
	"api-db/model"
	"api-db/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {

	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := model.User{}

	u.Username = input.Username
	u.Password = input.Password

	token, err := datasource.LoginCheck(datasource.DB, u.Username, u.Password)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"token": token})

}

func CurrentUser(c *gin.Context) {

	user_id, err := utils.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := datasource.GetUserByID(datasource.DB, user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}
