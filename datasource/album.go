package datasource

import (
	"api-db/model"

	"gorm.io/gorm"
)

var DB *gorm.DB

func AlbumsByArtist(db *gorm.DB, name string) ([]model.Album, error) {
	var albums []model.Album
	if err := db.Where("artist = ?", name).Find(&albums).Error; err != nil {
		return nil, err
	}
	return albums, nil
}

func AllAlbums(db *gorm.DB) ([]model.Album, error) {
	var albums []model.Album
	if err := db.Find(&albums).Error; err != nil {
		return nil, err
	}
	return albums, nil
}

func AlbumByID(db *gorm.DB, id string) (model.Album, error) {
	var alb model.Album
	if err := db.Where("id = ?", id).Find(&alb).Error; err != nil {
		return model.Album{}, err
	}
	return alb, nil
}

func AddAlbum(db *gorm.DB, alb *model.Album) error {
	return db.Create(alb).Error
}

func DeleteAlbumByID(db *gorm.DB, id string) error {
	var albums []model.Album
	if err := db.Find(&albums).Error; err != nil {
		return err
	}
	return db.Delete(&albums, "id = ?", id).Error

}

//------------------------------------------------------------------------implementare update passandogli i vari campi

func UpdateAlbum(db *gorm.DB, id string) (model.Album, error) {
	var album model.Album
	if err := db.Where("id = ?", id).Find(&album).Error; err != nil {
		return model.Album{}, err
	}
	return model.Album{}, nil
}
