package api

import (
	"api-db/datasource"
	"api-db/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	var albums []model.Album
	var err error
	titolo := c.Query("titolo")
	artista := c.Query("lista")
	if albums, err = datasource.AllAlbums(datasource.DB, titolo, artista); err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
	}

	c.IndentedJSON(http.StatusOK, albums)
}

func GetAlbumByID(c *gin.Context) {
	var album model.Album
	var err error
	var id int
	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		return
	}

	if album, err = datasource.AlbumByID(datasource.DB, id); err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
	}
	c.IndentedJSON(http.StatusFound, album)
}

func PostAlbums(c *gin.Context) {
	var newAlbum model.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		log.Println(c.BindJSON(&newAlbum))
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	if err := datasource.AddAlbum(datasource.DB, &newAlbum); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func DeleteAlbumByID(c *gin.Context) {
	var id int
	var err error
	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		return
	}
	if err = datasource.DeleteAlbumByID(datasource.DB, id); err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
	}
	c.IndentedJSON(http.StatusOK, nil)
}

func UpdateAlbum(c *gin.Context) {
	var newAlbum model.Album
	var id int
	var err error
	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		return
	}
	if err = c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	if id == newAlbum.ID {
		if err = datasource.UpdateAlbum(datasource.DB, newAlbum); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
	} else {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusOK, newAlbum)
}
