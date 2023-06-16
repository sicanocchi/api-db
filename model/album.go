package model

type Album struct {
	ID         int     `json:"id"`
	Titolo     string  `json:"titolo"`
	Artista_ID int     `json:"artista_id"`
	Prezzo     float64 `json:"prezzo"`
	Artista    Artista `json:"artista"`
}

func (album Album) TableName() string {
	return "album"
}
