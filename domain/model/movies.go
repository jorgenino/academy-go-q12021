package model

// Movie struct
type Movie struct {
	ID       int    `gorm:"primary_key" json:"ID"`
	Title    string `json:"Title"`
	Director string `json:"Director"`
}

// TableName function
func (Movie) TableName() string { return "movies" }
