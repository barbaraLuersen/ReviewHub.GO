package model

type Serie struct {
	Id             int    `json:"id"`
	Nome           string `json:"nome"`
	Datalancamento string `json:"datalancamento"`
	Temporadas     int    `json:"temporadas"`
	Resenha        string `json:"resenha"`
}

var Series []Serie
