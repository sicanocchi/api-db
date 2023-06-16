package api

import (
	"api-db/datasource"
	"api-db/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// -----------api artisti--------------------
/*
func GetArtistiByCasaDiscografica(c *gin.Context) {
	var artisti []model.Artista
	var err error
	nomeCasaDiscografica := c.Query("nome")
	if nomeCasaDiscografica != "" {
		if artisti, err = datasource.ArtistiByCasaDiscografica(datasource.DB, nomeCasaDiscografica); err != nil {
			c.IndentedJSON(http.StatusNotFound, err)
		}
	} else {
		if artisti, err = datasource.AllArtist(datasource.DB); err != nil {
			c.IndentedJSON(http.StatusNotFound, err)
		}
	}

	c.IndentedJSON(http.StatusOK, artisti)
}
*/
func GetArtisti(c *gin.Context) {
	var artisti []model.Artista
	var err error
	nomeArtista := c.Query("nome")
	cognomeArtista := c.Query("cognome")
	listaAlbum := c.Query("lista")
	if artisti, err = datasource.AllArtist(datasource.DB, nomeArtista, cognomeArtista, listaAlbum); err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
	}
	c.IndentedJSON(http.StatusOK, artisti)
}

func GetArtistaByID(c *gin.Context) {
	var artista model.Artista
	var err error
	var id int
	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		return
	}

	if artista, err = datasource.ArtistaByID(datasource.DB, id); err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
	}
	c.IndentedJSON(http.StatusFound, artista)
}

func PostArtista(c *gin.Context) {
	var newArtista model.Artista
	if err := c.BindJSON(&newArtista); err != nil {
		log.Println(c.BindJSON(&newArtista))
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	if err := datasource.AddArtista(datasource.DB, &newArtista); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusCreated, newArtista)
}

func DeleteArtistaByID(c *gin.Context) {
	var id int
	var err error
	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		return
	}
	if err = datasource.DeleteArtistaByID(datasource.DB, id); err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
	}
	c.IndentedJSON(http.StatusOK, nil)
}

func UpdateArtista(c *gin.Context) {
	var newArtista model.Artista
	var id int
	var err error
	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		return
	}
	if err = c.BindJSON(&newArtista); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	if id != newArtista.ID {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	} else {
		if err = datasource.UpdateArtista(datasource.DB, newArtista); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
	}
	c.IndentedJSON(http.StatusCreated, newArtista)
}
