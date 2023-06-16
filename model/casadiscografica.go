package model

type CasaDiscografica struct {
	ID             int       `json:"id"`
	Nome           string    `json:"nome"`
	DataFondazione int       `json:"data_fondazione"`
	Artista        []Artista `json:"artisti"`
}

func (casaDiscografica CasaDiscografica) TableName() string {
	return "casadiscografica"
}
