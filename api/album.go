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
	if albums, err = datasource.AllAlbums(datasource.DB, titolo); err != nil {
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
	nomeCasaDiscografica := c.Query("nome_casa")
	if nomeCasaDiscografica != "" {
		if artisti, err = datasource.ArtistiByCasaDiscografica(datasource.DB, nomeCasaDiscografica); err != nil {
			c.IndentedJSON(http.StatusNotFound, err)
		}
	} else if nomeArtista != "" && cognomeArtista != "" {
		if artisti, err = datasource.ArtistiByRagione(datasource.DB, nomeArtista, cognomeArtista); err != nil {
			c.IndentedJSON(http.StatusNotFound, err)
		}
	} else if nomeArtista != "" && cognomeArtista == "" {
		if artisti, err = datasource.ArtistiByNome(datasource.DB, nomeArtista); err != nil {
			c.IndentedJSON(http.StatusNotFound, err)
		}
	} else if nomeArtista == "" && cognomeArtista != "" {
		if artisti, err = datasource.ArtistiByCognome(datasource.DB, cognomeArtista); err != nil {
			c.IndentedJSON(http.StatusNotFound, err)
		}
	} else {
		if artisti, err = datasource.AllArtist(datasource.DB); err != nil {
			c.IndentedJSON(http.StatusNotFound, err)
		}
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

// -----------api casa discografica-----------
func GetCasaDiscograficaByID(c *gin.Context) {
	var casaDiscografica model.CasaDiscografica
	var err error
	var id int
	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		return
	}

	if casaDiscografica, err = datasource.CasaByID(datasource.DB, id); err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
	}
	c.IndentedJSON(http.StatusFound, casaDiscografica)
}

func GetCasaByNome(c *gin.Context) {
	var casaDiscografica []model.CasaDiscografica
	var err error
	nomeCasa := c.Query("nome")
	if nomeCasa != "" {
		if casaDiscografica, err = datasource.CaseByNome(datasource.DB, nomeCasa); err != nil {
			c.IndentedJSON(http.StatusNotFound, err)
		}
	} else {
		if casaDiscografica, err = datasource.AllCaseDiscografiche(datasource.DB); err != nil {
			c.IndentedJSON(http.StatusNotFound, err)
		}
	}

	c.IndentedJSON(http.StatusOK, casaDiscografica)
}

func PostCasaDiscografica(c *gin.Context) {
	var newCasaDiscografica model.CasaDiscografica
	if err := c.BindJSON(&newCasaDiscografica); err != nil {
		log.Println(c.BindJSON(&newCasaDiscografica))
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	if err := datasource.AddCasaDiscografica(datasource.DB, &newCasaDiscografica); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusCreated, newCasaDiscografica)
}

func DeleteCasaDiscograficaByID(c *gin.Context) {
	var id int
	var err error
	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		return
	}
	if err = datasource.DeleteCasaDiscograficaByID(datasource.DB, id); err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
	}
	c.IndentedJSON(http.StatusOK, nil)
}

func UpdateCasaDiscografica(c *gin.Context) {
	var newCasaDiscografica model.CasaDiscografica
	var id int
	var err error
	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		return
	}
	if err = c.BindJSON(&newCasaDiscografica); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	if id == newCasaDiscografica.ID {
		if err = datasource.UpdateCasaDiscografica(datasource.DB, newCasaDiscografica); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
	} else {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusCreated, newCasaDiscografica)
}
