package datasource

import (
	"api-db/model"
	"fmt"

	"gorm.io/gorm"
)

// -----------------nuove query casa discografica-----------------------------
func AllCaseDiscografiche(db *gorm.DB, nome string, lista string) ([]model.CasaDiscografica, error) {
	var caseDiscografiche []model.CasaDiscografica
	if nome != "" {
		carattere := "%"
		ricerca := fmt.Sprintf("%s%s%s", carattere, nome, carattere)
		db = db.Where("nome LIKE ?", ricerca)
	}
	if lista != "" && lista == "t" {
		db = db.Preload("Artista")
	}
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
