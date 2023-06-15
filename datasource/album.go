package datasource

import (
	"api-db/model"

	"gorm.io/gorm"
)

var DB *gorm.DB

func AlbumsByArtist(db *gorm.DB, nome string, cognome string) ([]model.Album, error) {
	var albums []model.Album
	var artisti []model.Artista
	var idArtisti []int
	if err := db.Where("nome = ? AND cognome =?", nome, cognome).Find(&artisti).Error; err != nil {
		return nil, err
	}
	for _, v := range artisti {
		idArtisti = append(idArtisti, v.ID)
	}
	if err := db.Where("artista_id IN ?", idArtisti).Find(&albums).Error; err != nil {
		return nil, err
	}
	return albums, nil
}

func AlbumsByArtistNome(db *gorm.DB, nome string) ([]model.Album, error) {
	var albums []model.Album
	var artisti []model.Artista
	var idArtisti []int
	if err := db.Where("nome = ?", nome).Find(&artisti).Error; err != nil {
		return nil, err
	}
	for _, v := range artisti {
		idArtisti = append(idArtisti, v.ID)
	}
	if err := db.Where("artista_id IN ?", idArtisti).Find(&albums).Error; err != nil {
		return nil, err
	}
	return albums, nil
}

func AlbumsByArtistCognome(db *gorm.DB, cognome string) ([]model.Album, error) {
	var albums []model.Album
	var artisti []model.Artista
	var idArtisti []int
	if err := db.Where("cognome =?", cognome).Find(&artisti).Error; err != nil {
		return nil, err
	}
	for _, v := range artisti {
		idArtisti = append(idArtisti, v.ID)
	}
	if err := db.Where("artista_id IN ?", idArtisti).Find(&albums).Error; err != nil {
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
	/*if err := db.Find(&albums).Error; err != nil {
		return err
	}*/
	return db.Delete(&albums, "id = ?", id).Error

}

func UpdateAlbum(db *gorm.DB, alb model.Album) error {
	return db.Save(&alb).Error

}

// -----------------nuove query artisti---------------------------------
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
}

func AllArtist(db *gorm.DB) ([]model.Artista, error) {
	var artisti []model.Artista
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

// -----------------nuove query casa discografica-----------------------------
func CaseByNome(db *gorm.DB, nome string) ([]model.CasaDiscografica, error) {
	var caseDiscografiche []model.CasaDiscografica
	if err := db.Where("nome = ?", nome).Find(&caseDiscografiche).Error; err != nil {
		return nil, err
	}
	return caseDiscografiche, nil
}

func AllCaseDiscografiche(db *gorm.DB) ([]model.CasaDiscografica, error) {
	var caseDiscografiche []model.CasaDiscografica
	if err := db.Find(&caseDiscografiche).Error; err != nil {
		return nil, err
	}
	return caseDiscografiche, nil
}

func CasaByID(db *gorm.DB, id int) (model.CasaDiscografica, error) {
	var casaDiscografica model.CasaDiscografica
	if err := db.Where("id = ?", id).Find(&casaDiscografica).Error; err != nil {
		return model.CasaDiscografica{}, err
	}
	return casaDiscografica, nil
}

func AddCasaDiscografica(db *gorm.DB, casaDiscografica *model.CasaDiscografica) error {
	return db.Create(casaDiscografica).Error
}

func DeleteCasaDiscograficaByID(db *gorm.DB, id int) error {
	var caseDiscografiche []model.CasaDiscografica
	/*if err := db.Find(&albums).Error; err != nil {
		return err
	}*/
	return db.Delete(&caseDiscografiche, "id = ?", id).Error

}

func UpdateCasaDiscografica(db *gorm.DB, cDisco model.CasaDiscografica) error {
	return db.Save(&cDisco).Error

}
