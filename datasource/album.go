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
	if err := db.Find(&albums).Error; err != nil {
		return err
	}
	return db.Delete(&albums, "id = ?", id).Error

}

func UpdateAlbum(db *gorm.DB, id int, alb model.Album) error {
	var album model.Album
	var err error
	if err = db.Where("id = ?", id).First(&album).Error; err != nil {
		return err
	}
	return db.Model(&album).Updates(model.Album{Title: alb.Title, Artist: alb.Artist, Price: alb.Price}).Error

}
