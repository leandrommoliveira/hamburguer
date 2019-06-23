package models

type Company struct {
	ID           int
	Name         string
	Localization string
	Average      float32
}

type Avaliation struct {
	ID          int
	IDCompany   int
	Name        string
	Average     float32
	Description string
}
