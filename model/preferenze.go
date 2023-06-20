package model

type Preferenze struct {
	UtenteId int `json:"id_utente"`
	AlbumId  int `json:"id_album"`
	Album    Album
}

func (preferenze Preferenze) TableName() string {
	return "preferenze"
}
