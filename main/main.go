package main

import (
	"api-db/api"
	"api-db/datasource"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	var err error

	dsn := "host= localhost user= postgres password= root  dbname= recordings  port= 5432  sslmode=disable"

	if datasource.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		panic(err)
	}

	router := gin.Default()

	router.POST("/register", api.Register)
	router.POST("/login", api.Login)

	protected := router.Group("/api/admin")

	protected.Use(api.JwtAuthMiddleware())
	protected.GET("/user", api.CurrentUser)

	//----------------chiamate per album-------------------
	router.GET("/albums", api.GetAlbums)

	router.GET("/albums/:id", api.GetAlbumByID)

	router.POST("/albums", api.PostAlbums)

	router.PUT("/albums/:id", api.UpdateAlbum)

	router.DELETE("/albums/:id", api.DeleteAlbumByID)

	//----------------chiamate per artisti-------------------
	router.GET("/artisti", api.GetArtisti)

	router.GET("/artisti/:id", api.GetArtistaByID)

	router.POST("/artisti", api.PostArtista)

	router.PUT("/artisti/:id", api.UpdateArtista)

	router.DELETE("/artisti/:id", api.DeleteArtistaByID)

	//----------------chiamate per case discografiche-------------------
	router.GET("/caseDiscografiche", api.GetCasa)

	router.GET("/caseDiscografiche/:id", api.GetCasaDiscograficaByID)

	router.POST("/caseDiscografiche", api.PostCasaDiscografica)

	router.PUT("/caseDiscografiche/:id", api.UpdateCasaDiscografica)

	router.DELETE("/caseDiscografiche/:id", api.DeleteCasaDiscograficaByID)

	router.Run("localhost:8080")

}
