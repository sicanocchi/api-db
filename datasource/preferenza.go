package datasource

import (
	"api-db/model"

	"gorm.io/gorm"
)

func AddPreferenza(db *gorm.DB, pref *model.Preferenze) error {
	return db.Create(pref).Error
}

func AllPreferenze(db *gorm.DB, user_id int) ([]model.Preferenze, error) {
	var preferenze []model.Preferenze
	/*if err := db.Select("album.titolo", "artista.cognome", "artista.nome").Joins("JOIN album on preferenze.album_id = album.id").Joins("JOIN artista on album.artista_id = artista.id").Where("utente_id = ?", user_id).Find(&preferenze).Error; err != nil {
		return nil, err
	}*/
	if err := db.Preload("Album").Where("utente_id = ?", user_id).Find(&preferenze).Error; err != nil {
		return nil, err
	}
	return preferenze, nil
}

func DeletePreferenzaByAlbumID(db *gorm.DB, user_id int, album_id int) error {
	var preferenze []model.Preferenze
	return db.Delete(&preferenze, "album_id = ? AND utente_id=?", album_id, user_id).Error

}
