package api

import (
	"api-db/datasource"
	"api-db/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

func GetCasa(c *gin.Context) {
	var casaDiscografica []model.CasaDiscografica
	var err error
	nomeCasa := c.Query("nome")
	lista := c.Query("lista")
	if casaDiscografica, err = datasource.AllCaseDiscografiche(datasource.DB, nomeCasa, lista); err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
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
