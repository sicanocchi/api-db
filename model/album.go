package model

type Album struct {
	ID         int     `json:"id"`
	Titolo     string  `json:"titolo"`
	Artista_ID int     `json:"artista_id"`
	Prezzo     float64 `json:"prezzo"`
	Artista    Artista `json:"artista"`
}

type Artista struct {
	ID                 int     `json:"id"`
	Nome               string  `json:"nome"`
	Cognome            string  `json:"cognome"`
	DataNascita        int     `json:"data_nascita"`
	CasaDiscograficaID int     `json:"casa_discografica_id"`
	Album              []Album `json:"album"`
}

type CasaDiscografica struct {
	ID             int       `json:"id"`
	Nome           string    `json:"nome"`
	DataFondazione int       `json:"data_fondazione"`
	Artista        []Artista `json:"artisti"`
}

func (album Album) TableName() string {
	return "album"
}

func (artista Artista) TableName() string {
	return "artista"
}

func (casaDiscografica CasaDiscografica) TableName() string {
	return "casadiscografica"
}
