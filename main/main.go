package main

import (
	"api-db/api"
	"api-db/datasource"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host= localhost user= postgres password= root  dbname= recordings  port= 5432  sslmode=disable"
	var err error
	if datasource.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		panic(err)
	}

	router := gin.Default()

	router.GET("/albums", api.GetAlbumsByArtist)

	router.GET("/albums/:id", api.GetAlbumByID)

	router.POST("/albums", api.PostAlbums)

	router.PUT("/albums/:id", api.UpdateAlbum)

	router.DELETE("/albums/:id", api.DeleteAlbumByID)

	router.Run("localhost:8080")

}
