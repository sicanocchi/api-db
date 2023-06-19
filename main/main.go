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
	protected.GET("/albums", api.GetAlbums)

	protected.GET("/albums/:id", api.GetAlbumByID)

	protected.POST("/albums", api.PostAlbums)

	protected.PUT("/albums/:id", api.UpdateAlbum)

	protected.DELETE("/albums/:id", api.DeleteAlbumByID)

	//----------------chiamate per artisti-------------------
	protected.GET("/artisti", api.GetArtisti)

	protected.GET("/artisti/:id", api.GetArtistaByID)

	protected.POST("/artisti", api.PostArtista)

	protected.PUT("/artisti/:id", api.UpdateArtista)

	protected.DELETE("/artisti/:id", api.DeleteArtistaByID)

	//----------------chiamate per case discografiche-------------------
	protected.GET("/caseDiscografiche", api.GetCasa)

	protected.GET("/caseDiscografiche/:id", api.GetCasaDiscograficaByID)

	protected.POST("/caseDiscografiche", api.PostCasaDiscografica)

	protected.PUT("/caseDiscografiche/:id", api.UpdateCasaDiscografica)

	protected.DELETE("/caseDiscografiche/:id", api.DeleteCasaDiscograficaByID)

	router.Run("localhost:8080")

}
