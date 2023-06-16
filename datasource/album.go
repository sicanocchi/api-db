package datasource

import (
	"api-db/model"
	"fmt"

	"gorm.io/gorm"
)

var DB *gorm.DB

func AllAlbums(db *gorm.DB, titolo string, lista string) ([]model.Album, error) {
	var albums []model.Album
	if titolo != "" {
		carattere := "%"
		ricerca := fmt.Sprintf("%s%s%s", carattere, titolo, carattere)
		db = db.Where("titolo LIKE ?", ricerca)
	}
	if lista != "" && lista == "t" {
		db = db.Preload("Artista")
	}
	if err := db.Find(&albums).Error; err != nil {
		return nil, err
	}
	return albums, nil
}

func AlbumByID(db *gorm.DB, id int) (model.Album, error) {
	var alb model.Album
	if err := db.Where("id = ?", id).Find(&alb).Error; err != nil {
		return model.Album{}, err
	}
	return alb, nil
}

func AddAlbum(db *gorm.DB, alb *model.Album) error {
	return db.Create(alb).Error
}

func DeleteAlbumByID(db *gorm.DB, id int) error {
	var albums []model.Album
	/*if err := db.Find(&albums).Error; err != nil {
		return err
	}*/
	return db.Delete(&albums, "id = ?", id).Error

}

func UpdateAlbum(db *gorm.DB, alb model.Album) error {
	return db.Save(&alb).Error

}
