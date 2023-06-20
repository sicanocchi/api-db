package api

import (
	"api-db/datasource"
	"api-db/model"
	"api-db/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPreferenze(c *gin.Context) {
	var prefer []model.Preferenze
	var user model.User
	var err error
	var token_id int
	token_id, err = utils.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err = datasource.GetUserByID(datasource.DB, token_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if prefer, err = datasource.AllPreferenze(datasource.DB, user.ID); err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
	}

	c.IndentedJSON(http.StatusOK, prefer)
}

func PostAlbumsOnPrefer(c *gin.Context) {
	var newPreferenza model.Preferenze
	var user model.User
	var err error
	var token_id int
	token_id, err = utils.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err = datasource.GetUserByID(datasource.DB, token_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newPreferenza.UtenteId = user.ID

	if newPreferenza.AlbumId, err = strconv.Atoi(c.Param("id")); err != nil {
		return
	}

	if err = datasource.AddPreferenza(datasource.DB, &newPreferenza); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusCreated, newPreferenza)
}

func DeletePreferenzaByAlbumID(c *gin.Context) {
	var user model.User
	var err error
	var token_id int
	var id_album int
	token_id, err = utils.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err = datasource.GetUserByID(datasource.DB, token_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if id_album, err = strconv.Atoi(c.Param("id")); err != nil {
		return
	}
	if err = datasource.DeletePreferenzaByAlbumID(datasource.DB, user.ID, id_album); err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
	}
	c.IndentedJSON(http.StatusOK, nil)
}
