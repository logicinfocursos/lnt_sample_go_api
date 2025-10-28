package structs

type Movie struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `json:"name"`
	Overview  string `json:"overview"`
	Posterurl string `json:"posterurl"`
	Genres    string `json:"genres"`
}
