package datasource

import (
	"api-db/model"
	"fmt"

	"gorm.io/gorm"
)

// -----------------nuove query artisti---------------------------------
/*
func ArtistiByRagione(db *gorm.DB, nome string, cognome string) ([]model.Artista, error) {
	var artisti []model.Artista
	if err := db.Where("nome = ? AND cognome =?", nome, cognome).Find(&artisti).Error; err != nil {
		return nil, err
	}
	return artisti, nil
}

func ArtistiByNome(db *gorm.DB, nome string) ([]model.Artista, error) {
	var artisti []model.Artista
	if err := db.Where("nome = ?", nome).Find(&artisti).Error; err != nil {
		return nil, err
	}
	return artisti, nil
}

func ArtistiByCognome(db *gorm.DB, cognome string) ([]model.Artista, error) {
	var artisti []model.Artista
	if err := db.Where("cognome = ?", cognome).Find(&artisti).Error; err != nil {
		return nil, err
	}
	return artisti, nil
}*/

func AllArtist(db *gorm.DB, nome string, cognome string, lista string) ([]model.Artista, error) {
	var artisti []model.Artista
	if nome != "" {
		carattere := "%"
		ricerca := fmt.Sprintf("%s%s%s", carattere, nome, carattere)
		db = db.Where("nome LIKE ?", ricerca)
	}
	if cognome != "" {
		carattere := "%"
		ricerca := fmt.Sprintf("%s%s%s", carattere, cognome, carattere)
		db = db.Where("cognome LIKE ?", ricerca)
	}
	if lista != "" && lista == "t" {
		db = db.Preload("Album")
	}
	if err := db.Find(&artisti).Error; err != nil {
		return nil, err
	}
	return artisti, nil
}

func ArtistaByID(db *gorm.DB, id int) (model.Artista, error) {
	var art model.Artista
	if err := db.Where("id = ?", id).Find(&art).Error; err != nil {
		return model.Artista{}, err
	}
	return art, nil
}

func AddArtista(db *gorm.DB, art *model.Artista) error {
	return db.Create(art).Error
}

func DeleteArtistaByID(db *gorm.DB, id int) error {
	var artisti []model.Artista
	/*if err := db.Find(&albums).Error; err != nil {
		return err
	}*/
	return db.Delete(&artisti, "id = ?", id).Error

}

func UpdateArtista(db *gorm.DB, art model.Artista) error {
	return db.Save(&art).Error

}

/*
func ArtistiByCasaDiscografica(db *gorm.DB, nome string) ([]model.Artista, error) {
	var artisti []model.Artista
	var cDisco model.CasaDiscografica
	if err := db.Where("nome = ?", nome).Find(&cDisco).Error; err != nil {
		return nil, err
	}
	if err := db.Where("casa_discografica_id = ?", cDisco.ID).Find(&artisti).Error; err != nil {
		return nil, err
	}
	return artisti, nil
}
*/
