package model

type Artista struct {
	ID                 int     `json:"id"`
	Nome               string  `json:"nome"`
	Cognome            string  `json:"cognome"`
	DataNascita        int     `json:"data_nascita"`
	CasaDiscograficaID int     `json:"casa_discografica_id"`
	Album              []Album `json:"album"`
}

func (artista Artista) TableName() string {
	return "artista"
}
