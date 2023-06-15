package model

type Album struct {
	ID         int     `json:"id"`
	Titolo     string  `json:"titolo"`
	Artista_ID int     `json:"artista"`
	Prezzo     float64 `json:"prezzo"`
}

type Artista struct {
	ID                 int    `json:"id"`
	Nome               string `json:"nome"`
	Cognome            string `json:"cognome"`
	DataNascita        int    `json:"data_nascita"`
	CasaDiscograficaID int    `json:"casa_discografica_id"`
}

type CasaDiscografica struct {
	ID             int    `json:"id"`
	Nome           string `json:"nome"`
	DataFondazione int    `json:"data_fondazione"`
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
