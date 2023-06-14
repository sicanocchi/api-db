package api

import (
	"api-db/datasource"
	"api-db/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAlbumsByArtist(c *gin.Context) {
	var albums []model.Album
	var err error
	artist := c.Query("artist")
	if artist != "" {
		if albums, err = datasource.AlbumsByArtist(datasource.DB, artist); err != nil {
			c.IndentedJSON(http.StatusNotFound, err)
		}
	} else {
		if albums, err = datasource.AllAlbums(datasource.DB); err != nil {
			c.IndentedJSON(http.StatusNotFound, err)
		}
	}

	c.IndentedJSON(http.StatusOK, albums)
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var album model.Album
	var err error
	if album, err = datasource.AlbumByID(datasource.DB, id); err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
	}
	c.IndentedJSON(http.StatusFound, album)
}

func PostAlbums(c *gin.Context) {
	var newAlbum model.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	if err := datasource.AddAlbum(datasource.DB, &newAlbum); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func DeleteAlbumByID(c *gin.Context) {
	id := c.Param("id")
	if err := datasource.DeleteAlbumByID(datasource.DB, id); err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
	}
	c.IndentedJSON(http.StatusOK, nil)
}

// ------------creare questa
func UpdateAlbum(c *gin.Context) {
	id := c.Param("id")
	var newAlbum model.Album
	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// gli dovrò passare i campi della nbuova variabile che vorrò cambiare
	if _, err := datasource.UpdateAlbum(datasource.DB, id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
